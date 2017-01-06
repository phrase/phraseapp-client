package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/phrase/phraseapp-go/phraseapp"
)

// The steps for a successful project initialization.
const (
	StepAskForToken   = "ask for token"
	StepSelectProject = "select project"
	StepSelectFormat  = "select format"
	StepConfigSources = "config sources"
	StepConfigTargets = "config targets"
	StepWriteConfig   = "write configuration file"
	StepFinished      = "finished"
)

var nextStep = map[string]string{
	StepAskForToken:   StepSelectProject,
	StepSelectProject: StepSelectFormat,
	StepSelectFormat:  StepConfigSources,
	StepConfigSources: StepConfigTargets,
	StepConfigTargets: StepWriteConfig,
	StepWriteConfig:   StepFinished,
}

type stepFunc func(*InitCommand) error

var stepFuncs = map[string]stepFunc{
	StepAskForToken:   (*InitCommand).askForToken,
	StepSelectProject: (*InitCommand).selectProject,
	StepSelectFormat:  (*InitCommand).selectFormat,
	StepConfigSources: (*InitCommand).configureSources,
	StepConfigTargets: (*InitCommand).configureTargets,
	StepWriteConfig:   (*InitCommand).writeConfig,
}

// structs that can be marshalled to YAML to create a valid configuration file

type ConfigYAML struct {
	Host        string                            `yaml:"host,omitempty"`
	AccessToken string                            `yaml:"access_token,omitempty"`
	ProjectID   string                            `yaml:"project_id"`
	FileFormat  string                            `yaml:"file_format,omitempty"`
	PerPage     int                               `yaml:"per_page,omitempty"`
	Defaults    map[string]map[string]interface{} `yaml:"defaults,omitempty"`
	Push        PushYAML                          `yaml:"push,omitempty"`
	Pull        PullYAML                          `yaml:"pull,omitempty"`
}

type PushYAML struct {
	Sources []SourcesYAML `yaml:"sources,omitempty"`
}

type PullYAML struct {
	Targets []TargetsYAML `yaml:"targets,omitempty"`
}

type SourcesYAML struct {
	File   string                 `yaml:"file,omitempty"`
	Params map[string]interface{} `yaml:"params,omitempty"`
}

type TargetsYAML SourcesYAML

// the actual command

type InitCommand struct {
	*phraseapp.Config

	client     *phraseapp.Client
	YAML       ConfigYAML
	FileFormat *phraseapp.Format
}

func (cmd *InitCommand) Run() error {
	if cmd.Config.Credentials == nil {
		cmd.Config.Credentials = &phraseapp.Credentials{}
	} else {
		// keep host if specified in config file or as command line parameter
		if cmd.Config.Credentials.Host != "" {
			cmd.YAML.Host = cmd.Config.Credentials.Host
		}
	}

	step := StepAskForToken

	for step != StepFinished {
		err := stepFuncs[step](cmd)
		if err != nil {
			return err
		}

		step = nextStep[step]
	}

	return nil
}

func (cmd *InitCommand) askForToken() error {
	printParrot()
	fmt.Println("PhraseApp.com API Client Setup")
	fmt.Println()

	token := ""
	err := prompt("Please enter your API Access Token (you can generate one in your profile at phraseapp.com): ", &token)
	if err != nil {
		return err
	}

	token = strings.ToLower(token)
	success, err := regexp.MatchString("^[0-9a-f]{64}$", token)
	if err != nil {
		return err
	}
	if !success {
		return fmt.Errorf("A valid Access Token is 64 characters long and contains only a-f, 0-9")
	}

	cmd.YAML.AccessToken = token

	cmd.Credentials.Token = token
	client, err := newClient(cmd.Credentials)
	if err != nil {
		return err
	}

	cmd.client = client
	return nil
}

func (cmd *InitCommand) selectProject() error {
	taskResult := make(chan []*phraseapp.Project, 1)
	taskErr := make(chan error, 1)

	client, err := newClient(cmd.Credentials)
	if err != nil {
		return err
	}

	withSpinner("Loading projects... ", func(taskFinished chan<- struct{}) {
		projects, err := client.ProjectsList(1, 25)
		taskResult <- projects
		taskErr <- err
		taskFinished <- struct{}{}
	})

	projects := <-taskResult
	if err := <-taskErr; err != nil {
		if strings.Contains(err.Error(), "401") {
			return fmt.Errorf("%s is not a valid Access Token. It may be revoked. Please create a new Token.", cmd.Credentials.Token)
		}
		return err
	}

	if len(projects) == 0 {
		fmt.Println("Since you don't have any projects yet, a new one will be created.")
		return cmd.newProject()
	}

	if len(projects) == 1 {
		answer := ""
		err := promptWithDefault(fmt.Sprintf("You have only one project. Answer 'y' to use '%s' or 'n' to create a new project: ", projects[0].Name), &answer, "n")
		if err != nil {
			return err
		}

		if answer == "y" {
			cmd.YAML.ProjectID = projects[0].ID
			return nil
		}

		return cmd.newProject()
	}

	for i, project := range projects {
		fmt.Printf("%2d: %s (Id: %s)\n", i+1, project.Name, project.ID)
	}
	fmt.Printf("%2d: Create new project\n", len(projects)+1)

	selection := 0
	for {
		err = prompt("Select project: ", &selection)
		if err != nil {
			return err
		}

		if selection < 1 || selection > len(projects)+1 {
			fmt.Println("Please select a project from the list by specifying its position in the list, e.g. 2 for the second project.")
		} else {
			break
		}
	}

	if selection == len(projects)+1 {
		return cmd.newProject()
	}

	cmd.YAML.ProjectID = projects[selection-1].ID
	cmd.DefaultFileFormat = projects[selection-1].MainFormat

	return nil
}

func (cmd *InitCommand) newProject() error {
	params := &phraseapp.ProjectParams{
		Name: new(string),
	}
	err := prompt("Enter the name of the new project: ", params.Name)
	if err != nil {
		return err
	}

	details, err := cmd.client.ProjectCreate(params)
	if err != nil {
		if strings.Contains(err.Error(), "401") {
			return fmt.Errorf("Your access token %s is not valid for the 'write' scope. Please create a new Access Token with read and write scope.", cmd.Credentials.Token)
		} else if strings.Contains(err.Error(), "Validation failed") {
			// TODO: really retry?
			return fmt.Errorf("Validation failed. Please try a different token.")
		}
		println(err)
		return err
	}

	cmd.YAML.ProjectID = details.ID

	return nil
}

func (cmd *InitCommand) selectFormat() error {
	formats, err := cmd.client.FormatsList(1, 25)
	if err != nil {
		return err
	}

	// ensure that the default file format from the config file is a valid format
	for _, format := range formats {
		if format.ApiName == cmd.DefaultFileFormat {
			cmd.FileFormat = format
			break
		}
	}

	for i, format := range formats {
		fmt.Printf("%2d: %s - %s, file extension: %s\n", i+1, format.ApiName, format.Name, format.Extension)
	}

	promptText := "Select the format to use for language files you download from PhraseApp"
	if cmd.FileFormat != nil {
		promptText += fmt.Sprintf(" (or leave blank to use the default, %s)", cmd.FileFormat.Name)
	}
	promptText += ": "

	selection := 0
	err = prompt(promptText, &selection)

	if err != nil {
		if (err == io.EOF && cmd.FileFormat == nil) && (selection < 1 || selection > len(formats)+1) {
			return fmt.Errorf("Please select a format from the list by specifying its position in the list.")
		}
	}

	cmd.FileFormat = formats[selection-1]
	return nil
}

func (cmd *InitCommand) configureSources() error {
	fmt.Println("Enter the path to the language file you want to upload to PhraseApp.")
	fmt.Printf("For documentation, see %s#push\n", docsConfigUrl)

	pushPath := ""
	for {
		err := promptWithDefault("Source file path: ", &pushPath, cmd.FileFormat.DefaultFile)
		if err != nil {
			return err
		}

		println(pushPath)

		err = ValidPath(pushPath, cmd.FileFormat.ApiName, cmd.FileFormat.Extension)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	sourceYAML := SourcesYAML{
		File: pushPath,
		Params: map[string]interface{}{
			"file_format": cmd.FileFormat.ApiName,
		},
	}

	cmd.YAML.Push.Sources = append(cmd.YAML.Push.Sources, sourceYAML)

	return nil
}

func (cmd *InitCommand) configureTargets() error {
	fmt.Println("Enter the path to which to download language files from PhraseApp.")
	fmt.Printf("For documentation, see %s#pull\n", docsConfigUrl)

	pullPath := ""
	for {
		err := promptWithDefault("Target file path: ", &pullPath, cmd.FileFormat.DefaultFile)
		if err != nil {
			return err
		}

		err = ValidPath(pullPath, cmd.FileFormat.ApiName, cmd.FileFormat.Extension)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	targetYAML := TargetsYAML{
		File: pullPath,
		Params: map[string]interface{}{
			"file_format": cmd.FileFormat.ApiName,
		},
	}

	cmd.YAML.Pull.Targets = append(cmd.YAML.Pull.Targets, targetYAML)

	return nil
}

func (cmd *InitCommand) writeConfig() error {
	wrapper := struct {
		Config ConfigYAML `yaml:"phraseapp"`
	}{
		Config: cmd.YAML,
	}

	yamlBytes, err := yaml.Marshal(wrapper)
	if err != nil {
		return err
	}

	filename := ".phraseapp.yml"
	err = ioutil.WriteFile(filename, yamlBytes, 0655)
	if err != nil {
		return err
	}

	printSuccess("We created the following configuration file for you: " + filename)

	fmt.Println()
	fmt.Println(string(yamlBytes))

	printSuccess("For advanced configuration options, take a look at the documentation: " + docsConfigUrl)
	printSuccess("You can now use the push & pull commands in your workflow:")
	fmt.Println()
	fmt.Println("$ phraseapp push")
	fmt.Println("$ phraseapp pull")
	fmt.Println()

	pushNow := ""
	err = promptWithDefault("Do you want to upload your locales now for the first time? (y/n) ", &pushNow, "y")
	if pushNow == "y" {
		err = firstPush()
		if err != nil {
			return err
		}
	}

	printSuccess("Project initialization completed!")

	return nil
}

func firstPush() error {
	cfg, err := phraseapp.ReadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}
	cmd := &PushCommand{Config: cfg}
	return cmd.Run()
}
