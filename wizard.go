package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-go/phraseapp"
)

const docsURL = "http://docs.phraseapp.com/developers/cli/configuration"
const parrot = `
                               ppe############eep
                         p############################pp
                    p######################################p
                p#############################################pp
            p###################ERRR  hrrr   PREE################p
    p pp################ERHhr                        HE#############p
   E############EERHrr                                  rPE###########
      rPPPPrr                          pp############pp###R PE##########
                                    p########EEE#########R     S#########p
      #                           ########R   ##p#######E       H#########p
    p####                       p########h    E##h ####R    E##e P##########
    #######p                   ##########p        a###R      E##  A##########
   ########h                  ############p      p###E        P##  E#########p
  S#######E                  a#######################          A#p  ##########
  ########                   ######################R            E#  ##########p
 #########                   #################ERH                #  ###########
 S#######h                   A##############R                      a###########h
 S#######h                    P#############                       ############h
 P#######h                     S############                      #############h
 S#######p                     rE############p                   S#############h
 E########                       PS############p               p###############
 P########p                        PE############p            #################
  S########p                           PRRSS#######       p###################h
   S########                                 p#####h  pp#####################R
   P#########p                            a##################################
    P#########p                            #################################
      ##########p                          S##############################E
       E##########p                         S############################R
        H###########p                        S#########################Eh
          E############p                      H#######################R
           AE##############p                    P###################R
              PS################ppp               PS#############E
                 R############################################ER
                    HS#####################################Rh
                        PRS##########################SRRhr
                              PHRASEAPP###PPAESARHP
`

type WizardCommand struct {
	Host  string `cli:"opt --host"`
	Debug bool   `cli:"opt --verbose -v"`
}

func (cmd *WizardCommand) Run() error {
	Debug = cmd.Debug
	data := WizardData{Host: cmd.Host}
	err := DisplayWizard(&data, "", "")
	if err != nil {
		printError(err)
	}
	return nil
}

type WizardData struct {
	Host        string `yaml:"host,omitempty"`
	AccessToken string `yaml:"access_token"`
	ProjectID   string `yaml:"project_id"`
	Format      string `yaml:"file_format"`
	MainFormat  string `yaml:"-"`
	Step        string `yaml:"-"`
	Push        struct {
		Sources WizardSources
	}
	Pull struct {
		Targets WizardTargets
	}

	FormatExtension string `yaml:"-"`
}

type WizardWrapper struct {
	Data *WizardData `yaml:"phraseapp"`
}

type WizardSources []*WizardPushConfig
type WizardTargets []*WizardPullConfig

type WizardPushConfig struct {
	Dir         string            `yaml:"dir,omitempty"`
	File        string            `yaml:"file,omitempty"`
	ProjectID   string            `yaml:"project_id,omitempty"`
	AccessToken string            `yaml:"access_token,omitempty"`
	Params      *WizardPushParams `yaml:"params,omitempty"`
}

type WizardPullConfig struct {
	Dir         string            `yaml:"dir,omitempty"`
	File        string            `yaml:"file,omitempty"`
	ProjectID   string            `yaml:"project_id,omitempty"`
	AccessToken string            `yaml:"access_token,omitempty"`
	Params      *WizardPullParams `yaml:"params,omitempty"`
}

type WizardPullParams struct {
	FileFormat string `yaml:"file_format,omitempty"`
	LocaleId   string `yaml:"locale_id,omitempty"`
}
type WizardPushParams struct {
	FileFormat string `yaml:"file_format,omitempty"`
	LocaleId   string `yaml:"locale_id,omitempty"`
}

func clean() error {
	switch runtime.GOOS {
	case "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		errmsg := fmt.Sprintf("%s is unsupported for the wizard", runtime.GOOS)
		ReportError("Unsupported OS", errmsg)
		return fmt.Errorf(errmsg)
	}
	return nil
}

func spinner(waitMsg string, position int, channelEnd *ChannelEnd, wg *sync.WaitGroup) {
	if channelEnd.closed {
		wg.Done()
		return
	}

	wg.Add(1)
	chars := []string{".", "-", ".", "-", "-", ".", "-", ".", "-", "-", ".", "-", ".", "-", "-", "."}
	if position > len(chars)-1 {
		position = 0
	}
	postfix := ""
	prefix := ""
	for counter, str := range chars {
		if counter < position {
			postfix = fmt.Sprint(postfix, str)
		} else {
			prefix = fmt.Sprint(prefix, str)
		}
	}

	clean()
	switch runtime.GOOS {
	case "windows":
		prefix = "..."
		postfix = ""
	}
	printWait(fmt.Sprintf("%s %s%s", waitMsg, prefix, postfix))
	time.Sleep(100 * time.Millisecond)

	spinner(waitMsg, position+1, channelEnd, wg)

	wg.Done()
}

func printParrot() {

	parrotLines := strings.Split(parrot, "\n")
	ct.Foreground(ct.Cyan, true)
	for _, line := range parrotLines {
		fmt.Println(line)
	}
	ct.ResetColor()
}

func printErrorStr(errorMsg string) {
	printWithColor(errorMsg, ct.Red, true)
}
func printError(err error) {
	printWithColor(err.Error(), ct.Red, true)
}

func printWait(msg string) {
	printWithColor(msg, ct.Yellow, true)
}

func printWithColor(msg string, color ct.Color, bright bool) {
	ct.Foreground(color, bright)
	fmt.Println(msg)
	ct.ResetColor()
}

func printSuccess(msg string) {
	printWithColor(msg, ct.Green, true)
}

func DisplayWizard(data *WizardData, step string, errorMsg string) error {
	err := clean()
	if err != nil {
		return err
	}

	if errorMsg != "" {
		printErrorStr(errorMsg)
	}
	switch {

	case step == "" || data.AccessToken == "":
		data.Step = "token"
		return tokenStep(data)
	case step == "newProject":
		data.Step = "newProject"
		return newProjectStep(data)
	case step == "selectProject":
		data.Step = "selectProject"
		return selectProjectStep(data)
	case step == "selectFormat":
		data.Step = "selectFormat"
		return selectFormat(data)
	case step == "pushConfig":
		data.Step = "pushConfig"
		return pushConfig(data)
	case step == "pullConfig":
		data.Step = "pullConfig"
		return pullConfig(data)
	case step == "finish":
		return writeConfig(data, ".phraseapp.yml")
	}

	errmsg := fmt.Sprintf("Step %s not known in init wizard", step)
	ReportError("Invalid Wizard Step", errmsg)
	return fmt.Errorf(errmsg)

}

func defaultFilePath(fileFormat string) (string, error) {
	formats, err := client.FormatsList(1, 30)
	if err != nil {
		return "", err
	}
	for _, format := range formats {
		if format.ApiName == fileFormat {
			return format.DefaultFile, nil
		}
	}
	return "", nil
}

func pushConfig(data *WizardData) error {
	defaultPath, err := defaultFilePath(data.Format)
	if err != nil {
		return err
	}

	fmt.Println("Enter the path to your local language files, you want to upload to PhraseApp.")
	fmt.Println("For documentation see http://docs.phraseapp.com/developers/cli/configuration/#sources")

	var pushPath string
	for {
		fmt.Printf("\nSource path [default: %s]: ", defaultPath)
		pushPath = prompt()
		if pushPath == "" {
			pushPath = defaultPath
		}

		err := ValidPath(pushPath, data.Format, data.FormatExtension)
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	data.Push.Sources = make(WizardSources, 1)
	data.Push.Sources[0] = &WizardPushConfig{File: pushPath, Params: &WizardPushParams{FileFormat: data.Format}}

	return DisplayWizard(data, next(data), "")
}

func pullConfig(data *WizardData) error {
	defaultPath, err := defaultFilePath(data.Format)
	if err != nil {
		return err
	}

	fmt.Println("Enter the path to where you want to store downloaded language files from PhraseApp.")
	fmt.Println("For documentation see http://docs.phraseapp.com/developers/cli/configuration/#targets")

	var pullPath string
	for {
		fmt.Printf("\nTarget path [default: %s]: ", defaultPath)
		pullPath = prompt()
		if pullPath == "" {
			pullPath = defaultPath
		}

		err := ValidPath(pullPath, data.Format, data.FormatExtension)
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	data.Pull.Targets = make([]*WizardPullConfig, 1)
	data.Pull.Targets[0] = &WizardPullConfig{File: pullPath, Params: &WizardPullParams{FileFormat: data.Format}}
	return DisplayWizard(data, next(data), "")
}

var client *phraseapp.Client

func selectFormat(data *WizardData) error {
	auth := &phraseapp.Credentials{Token: data.AccessToken}
	client, err := phraseapp.NewClient(auth)
	if err != nil {
		return err
	}
	formats, err := client.FormatsList(1, 25)
	if err != nil {
		return err
	}

	for counter, format := range formats {
		fmt.Printf("%2d. %s - %s, File-Extension: %s\n", counter+1, format.ApiName, format.Name, format.Extension)
	}

	var id string
	mainFormatDefault := ""
	if data.MainFormat != "" {
		mainFormatDefault = fmt.Sprintf(" [Press enter for default: %s]", data.MainFormat)
	}
	fmt.Printf("Select the format you want to use for language files you download from PhraseApp%s: ", mainFormatDefault)
	id = prompt()

	if id == "" && data.MainFormat != "" {
		data.Format = data.MainFormat
		return DisplayWizard(data, next(data), "")
	}
	number, err := strconv.Atoi(id)
	if err != nil {
		number = 0
		for index, format := range formats {
			if format.ApiName == id {
				number = index + 1
			}
		}

		if number < 1 {
			return DisplayWizard(data, "selectFormat", fmt.Sprintf("Argument Error: Please select a format from the list by specifying its position in the list."))
		}

	} else if number < 1 || number > len(formats)+1 {
		return DisplayWizard(data, "selectFormat", fmt.Sprintf("Argument Error: Please select a format from the list by specifying its position in the list."))
	}
	data.Format = formats[number-1].ApiName
	data.FormatExtension = formats[number-1].Extension
	return DisplayWizard(data, next(data), "")
}

func writeConfig(data *WizardData, filename string) error {
	wrapper := WizardWrapper{Data: data}
	bytes, err := yaml.Marshal(wrapper)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, bytes, 0655)
	if err != nil {
		return err
	}
	str := fmt.Sprintf("Success! We have created the config file for you %s:", filename)
	printSuccess(str)
	fmt.Println("")
	fmt.Println(string(bytes))

	printSuccess("You can make changes to this file, see this documentation for more advanced options: " + docsURL)
	printSuccess("Now start using phraseapp push & pull for your workflow:")
	fmt.Println("")
	fmt.Println("$ phraseapp push")
	fmt.Println("$ phraseapp pull")
	fmt.Println("")
	var initialPush string
	fmt.Print("Enter \"y\" to upload your locales now for the first time (Default: \"y\"): ")
	initialPush = prompt()
	if initialPush == "y" || initialPush == "" {
		err = firstPush()
		if err != nil {
			return err
		}
	}
	fmt.Println("Setup completed!")
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

func next(data *WizardData) string {
	switch data.Step {
	case "", "token":
		return "selectProject"
	case "selectProject":
		return "selectFormat"
	case "newProject":
		return "selectFormat"
	case "selectFormat":
		return "pushConfig"
	case "pushConfig":
		return "pullConfig"
	case "pullConfig":
		return "finish"
	}
	return ""
}

func tokenStep(data *WizardData) error {
	printParrot()
	fmt.Println("PhraseApp.com API Client Setup")
	fmt.Println("")
	fmt.Print("Please enter you API Access Token (Generate one in your profile at phraseapp.com): ")
	data.AccessToken = prompt()
	data.AccessToken = strings.ToLower(data.AccessToken)
	success, err := regexp.MatchString("^[0-9a-f]{64}$", data.AccessToken)
	if err != nil {
		return err
	}
	if !success {
		data.AccessToken = ""
		return DisplayWizard(data, "", "Argument Error: AccessToken must be 64 letters long and can only contain a-f, 0-9")
	}
	return DisplayWizard(data, next(data), "")
}

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	return string(bytes)
}

func newProjectStep(data *WizardData) error {
	fmt.Print("Enter name of new project: ")
	projectParam := &phraseapp.ProjectParams{}
	str := prompt()
	projectParam.Name = &str

	res, err := client.ProjectCreate(projectParam)
	if err != nil {
		success, match_err := regexp.MatchString("401", err.Error())
		if match_err != nil {
			return match_err
		}
		if success {
			data.AccessToken = ""
			return DisplayWizard(data, "", fmt.Sprintf("Argument Error: Your AccessToken '%s' has no write scope. Please create a new Access Token with read and write scope.", data.AccessToken))
		} else {
			_, match_err := regexp.MatchString("Validation failed", err.Error())
			if match_err != nil {
				return match_err
			}
			return DisplayWizard(data, "newProject", err.Error())
		}
	}
	data.ProjectID = res.ID
	return DisplayWizard(data, next(data), "")
}

type ChannelEnd struct {
	closed bool
}

func selectProjectStep(data *WizardData) error {
	auth := &phraseapp.Credentials{Token: data.AccessToken, Host: data.Host}
	fmt.Println("Please select your project:")
	var err error
	client, err = phraseapp.NewClient(auth)

	var wg sync.WaitGroup
	out := make(chan []phraseapp.Project, 1)
	wg.Add(1)
	channelEnd := ChannelEnd{}
	getProjects := func(channelEnd *ChannelEnd) {
		var projects []*phraseapp.Project
		// time.Sleep(500 * time.Millisecond)
		projects, err = client.ProjectsList(1, 25)
		var array []phraseapp.Project
		for _, res := range projects {
			array = append(array, *res)
		}
		out <- array
		channelEnd.closed = true
		return
	}
	go getProjects(&channelEnd)
	go func(channelEnd *ChannelEnd, wg *sync.WaitGroup) {
		spinner("Loading projects... ", 0, channelEnd, wg)
	}(&channelEnd, &wg)
	var projects []phraseapp.Project

	projects = <-out
	clean()
	wg.Wait()
	close(out)

	if err != nil {
		unauth_match, match_err := regexp.MatchString("401", err.Error())
		if match_err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			panic(match_err)
		}
		if unauth_match {
			errmsg := fmt.Sprintf("Argument Error: AccessToken '%s' is invalid. It may be revoked. Please create a new Access Token.", data.AccessToken)
			data.AccessToken = ""
			ReportError("Invalid AccessToken", errmsg)
			return fmt.Errorf(errmsg)
		} else {
			return err
		}
	}

	if len(projects) == 1 {
		data.ProjectID = projects[0].ID
		data.MainFormat = projects[0].MainFormat
		fmt.Printf("You've got one project, \"%s\". Answer \"y\" to select this or \"n\" to create a new project: ", projects[0].Name)
		answer := prompt()
		if answer == "y" {
			return DisplayWizard(data, next(data), "")
		} else {
			data.ProjectID = ""
			data.MainFormat = ""
			return DisplayWizard(data, "newProject", "")
		}
	}
	for counter, project := range projects {
		fmt.Printf("%2d. %s (Id: %s)\n", counter+1, project.Name, project.ID)
	}
	fmt.Printf("%2d. Create new project\n", len(projects)+1)
	fmt.Print("Select project: ")
	id := prompt()
	number, err := strconv.Atoi(id)
	if err != nil || number < 1 || number > len(projects)+1 {
		return DisplayWizard(data, "selectProject", fmt.Sprintf("Argument Error: Please select a project from the list by specifying its position in the list, e.g. 2 for the second project."))
	}

	if number == len(projects)+1 {
		return DisplayWizard(data, "newProject", "")
	}

	selectedProject := projects[number-1]
	data.ProjectID = selectedProject.ID
	data.MainFormat = selectedProject.MainFormat
	return DisplayWizard(data, next(data), "")
}
