package print

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

func Parrot() {
	WithColor(ct.Cyan, parrot)
}

func Success(msg string, args ...interface{}) {
	WithColor(ct.Green, msg, args...)
}

func Failure(msg string, args ...interface{}) {
	WithColor(ct.Red, msg, args...)
}

func WithColor(color ct.Color, msg string, args ...interface{}) {
	fprintWithColor(os.Stdout, color, msg, args...)
}

func Error(err error) {
	fprintWithColor(os.Stderr, ct.Red, "ERROR: %s", err)
}

func fprintWithColor(w io.Writer, color ct.Color, msg string, args ...interface{}) {
	ct.Foreground(color, true)
	fmt.Fprintf(w, msg, args...)
	fmt.Fprintln(w)
	ct.ResetColor()
}
