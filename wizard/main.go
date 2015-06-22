package main

import "github.com/phrase/phraseapp-client/wizard/wizard"

func main() {
	data := wizard.WizardData{}
	wizard.DisplayWizard(&data, "", "")
}
