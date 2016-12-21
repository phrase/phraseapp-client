package paclient

import (
	"fmt"
	"io"
	"os"

	ct "github.com/daviddengcn/go-colortext"
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

func PrintParrot() {
	printWithColor(ct.Cyan, parrot)
}

func PrintSuccess(msg string, args ...interface{}) {
	printWithColor(ct.Green, msg, args...)
}

func PrintFailure(msg string, args ...interface{}) {
	printWithColor(ct.Red, msg, args...)
}

func printWithColor(color ct.Color, msg string, args ...interface{}) {
	fprintWithColor(os.Stdout, color, msg, args...)
}

func PrintError(err error) {
	fprintWithColor(os.Stderr, ct.Red, "ERROR: %s", err)
}

func fprintWithColor(w io.Writer, color ct.Color, msg string, args ...interface{}) {
	ct.Foreground(color, true)
	fmt.Fprintf(w, msg, args...)
	fmt.Fprintln(w)
	ct.ResetColor()
}

func sharedMessage(method string, localeFile *LocaleFile) {
	local := localeFile.RelPath()

	if method == "pull" {
		remote := localeFile.Message()
		fmt.Print("Downloaded ")
		ct.Foreground(ct.Green, true)
		fmt.Print(remote)
		ct.ResetColor()
		fmt.Print(" to ")
		ct.Foreground(ct.Green, true)
		fmt.Print(local, "\n")
		ct.ResetColor()
	} else {
		fmt.Print("Uploaded ")
		ct.Foreground(ct.Green, true)
		fmt.Print(local)
		ct.ResetColor()
		fmt.Println(" successfully.")
	}
}
