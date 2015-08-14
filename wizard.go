package main

import (
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

	"gopkg.in/yaml.v1"

	"github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-go/phraseapp"
)

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
	Host string `cli:"opt --host"`
}

func (cmd *WizardCommand) Run() error {
	data := WizardData{Host: cmd.Host}
	DisplayWizard(&data, "", "")
	return nil
}

type WizardData struct {
	Host        string `yaml:"host"`
	AccessToken string `yaml:"access_token"`
	ProjectId   string `yaml:"project_id"`
	Format      string `yaml:"file_format"`
	MainFormat  string `yaml:"-"`
	Step        string `yaml:"-"`
	Push        struct {
		Sources WizardSources
	}
	Pull struct {
		Targets WizardTargets
	}
}

type WizardWrapper struct {
	Data *WizardData `yaml:"phraseapp"`
}

type WizardSources []*WizardPushConfig
type WizardTargets []*WizardPullConfig

type WizardPushConfig struct {
	Dir         string            `yaml:"dir,omitempty"`
	File        string            `yaml:"file,omitempty"`
	ProjectId   string            `yaml:"project_id,omitempty"`
	AccessToken string            `yaml:"access_token,omitempty"`
	Params      *WizardPushParams `yaml:"params,omitempty"`
}

type WizardPullConfig struct {
	Dir         string            `yaml:"dir,omitempty"`
	File        string            `yaml:"file,omitempty"`
	ProjectId   string            `yaml:"project_id,omitempty"`
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

func clean() {
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
		fmt.Printf("%s unsupported", runtime.GOOS)
		panic("Do not know")
	}
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

func printError(errorMsg string) {
	printWithColor(errorMsg, ct.Red, true)
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

func DisplayWizard(data *WizardData, step string, errorMsg string) {
	clean()

	if errorMsg != "" {
		printError(errorMsg)
	}
	switch {

	case step == "" || data.AccessToken == "":
		data.Step = "token"
		tokenStep(data)
		return
	case step == "newProject":
		data.Step = "newProject"
		newProjectStep(data)
		return
	case step == "selectProject":
		data.Step = "selectProject"
		selectProjectStep(data)
		return
	case step == "selectFormat":
		data.Step = "selectFormat"
		selectFormat(data)
		return
	case step == "pushConfig":
		data.Step = "pushConfig"
		pushConfig(data)
		return
	case step == "pullConfig":
		data.Step = "pullConfig"
		pullConfig(data)
		return
	case step == "finish":
		writeConfig(data, ".phraseapp.yml")
		return
	}

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

func pushConfig(data *WizardData) {
	defaultPath, err := defaultFilePath(data.Format)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Enter the path from where you want to upload your locale files to Phrase [Press enter to use default: %s]: ", defaultPath)
	var pushPath string
	fmt.Scanln(&pushPath)
	if pushPath == "" {
		pushPath = defaultPath
	}

	data.Push.Sources = make(WizardSources, 1)
	data.Push.Sources[0] = &WizardPushConfig{File: pushPath, Params: &WizardPushParams{FileFormat: data.Format}}

	DisplayWizard(data, next(data), "")
}

func pullConfig(data *WizardData) {
	defaultPath, err := defaultFilePath(data.Format)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Enter the path you want to put the downloaded locale files from Phrase [Press enter to use default: %s]: ", defaultPath)
	var pullPath string
	fmt.Scanln(&pullPath)
	if pullPath == "" {
		pullPath = defaultPath
	}

	data.Pull.Targets = make([]*WizardPullConfig, 1)
	data.Pull.Targets[0] = &WizardPullConfig{File: pullPath, Params: &WizardPullParams{FileFormat: data.Format}}
	DisplayWizard(data, next(data), "")
}

var client *phraseapp.Client

func selectFormat(data *WizardData) {
	auth := phraseapp.Credentials{Token: data.AccessToken}
	var err error
	client, err = phraseapp.NewClient(auth, nil)
	formats, err := client.FormatsList(1, 25)
	if err != nil {
		panic(err.Error())
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
	fmt.Scanln(&id)
	if id == "" && data.MainFormat != "" {
		data.Format = data.MainFormat
		DisplayWizard(data, next(data), "")
		return
	}
	number, err := strconv.Atoi(id)
	if err != nil || number < 1 || number > len(formats)+1 {
		DisplayWizard(data, "selectFormat", fmt.Sprintf("Argument Error: Please select a format from the list by specifying its position in the list."))
		return
	}
	data.Format = formats[number-1].ApiName
	DisplayWizard(data, next(data), "")
}

func writeConfig(data *WizardData, filename string) {
	wrapper := WizardWrapper{Data: data}
	bytes, err := yaml.Marshal(wrapper)
	if err != nil {
		panic(err.Error())
	}
	err = ioutil.WriteFile(filename, bytes, 0655)
	if err != nil {
		panic(err.Error())
	}
	str := fmt.Sprintf("Success! We have created the config file for you %s:", filename)
	printSuccess(str)
	fmt.Println("")
	fmt.Println(string(bytes))

	printSuccess("You can make changes to this file, see this documentation for more advanced options: http://docs.phraseapp.com/api/v2/config")
	printSuccess("Now start using phraseapp push & pull for your workflow:")
	fmt.Println("")
	fmt.Println("$ phraseapp push")
	fmt.Println("$ phraseapp pull")
	fmt.Println("")
	var initialPush string
	fmt.Print("Enter \"y\" to upload your locales now for the first time (Default: \"y\"): ")
	fmt.Scanln(&initialPush)
	if initialPush == "y" || initialPush == "" {
		err = firstPush()
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Setup completed!")
}

func firstPush() error {
	cmd := &PushCommand{}
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

func tokenStep(data *WizardData) {
	printParrot()
	fmt.Println("PhraseApp.com API Client Setup")
	fmt.Println("")
	fmt.Print("Please enter you API Access Token (Generate one in your profile at phraseapp.com): ")
	fmt.Scanln(&data.AccessToken)
	data.AccessToken = strings.ToLower(data.AccessToken)
	success, err := regexp.MatchString("^[0-9a-f]{64}$", data.AccessToken)
	if err != nil {
		panic(err.Error())
	}
	if success == true {
		DisplayWizard(data, next(data), "")
	} else {
		data.AccessToken = ""
		DisplayWizard(data, "", "Argument Error: AccessToken must be 64 letters long and can only contain a-f, 0-9")
	}
}

func newProjectStep(data *WizardData) {
	fmt.Print("Enter name of new project: ")
	projectParam := &phraseapp.ProjectParams{}
	fmt.Scanln(&projectParam.Name)

	res, err := client.ProjectCreate(projectParam)
	if err != nil {
		success, match_err := regexp.MatchString("401", err.Error())
		if match_err != nil {
			fmt.Println(err.Error())
			panic(match_err.Error())
		}
		if success {
			data.AccessToken = ""
			DisplayWizard(data, "", fmt.Sprintf("Argument Error: Your AccessToken '%s' has no write scope. Please create a new Access Token with read and write scope.", data.AccessToken))
		} else {
			success, match_err := regexp.MatchString("Validation failed", err.Error())
			if match_err != nil {
				fmt.Println(err.Error())
				panic(match_err.Error())
			}
			if success {
				DisplayWizard(data, "newProject", err.Error())
				return
			} else {
				panic(err.Error())
			}
		}
	} else {
		data.ProjectId = res.Id
		DisplayWizard(data, next(data), "")
		return
	}
}

type ChannelEnd struct {
	closed bool
}

func selectProjectStep(data *WizardData) {
	auth := phraseapp.Credentials{Token: data.AccessToken, Host: data.Host}
	fmt.Println("Please select your project:")
	var err error
	client, err = phraseapp.NewClient(auth, nil)
	if err != nil {
		panic(err)
	}
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
		success, match_err := regexp.MatchString("401", err.Error())
		if match_err != nil {
			fmt.Println(err.Error())
			panic(match_err.Error())
		}
		if success {
			errorMsg := fmt.Sprintf("Argument Error: AccessToken '%s' is invalid. It may be revoked. Please create a new Access Token.", data.AccessToken)
			data.AccessToken = ""
			DisplayWizard(data, "", errorMsg)
		} else {
			panic(err.Error())
		}
	}

	if len(projects) == 1 {
		data.ProjectId = projects[0].Id
		data.MainFormat = projects[0].MainFormat
		fmt.Printf("You've got one project, \"%s\". Answer \"y\" to select this or \"n\" to create a new project: ", projects[0].Name)
		var answer string
		fmt.Scanln(&answer)
		if answer == "y" {
			DisplayWizard(data, next(data), "")
			return
		} else {
			data.ProjectId = ""
			data.MainFormat = ""
			DisplayWizard(data, "newProject", "")
			return
		}
	}
	for counter, project := range projects {
		fmt.Printf("%2d. %s (Id: %s)\n", counter+1, project.Name, project.Id)
	}
	fmt.Printf("%2d. Create new project\n", len(projects)+1)
	fmt.Print("Select project: ")
	var id string
	fmt.Scanln(&id)
	number, err := strconv.Atoi(id)
	if err != nil || number < 1 || number > len(projects)+1 {
		DisplayWizard(data, "selectProject", fmt.Sprintf("Argument Error: Please select a project from the list by specifying its position in the list, e.g. 2 for the second project."))
		return
	}

	if number == len(projects)+1 {
		DisplayWizard(data, "newProject", "")
		return
	}

	selectedProject := projects[number-1]
	data.ProjectId = selectedProject.Id
	data.MainFormat = selectedProject.MainFormat
	DisplayWizard(data, next(data), "")
}
