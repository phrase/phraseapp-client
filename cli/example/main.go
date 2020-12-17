package main

import (
	"log"
	"strings"

	"github.com/phrase/phraseapp-client/cli"
)

type BaseEntity struct {
	Af string `cli:"type=opt short=a desc='something set in the base entity'"`
	Bf int    `cli:"type=opt short=b desc='something set in the base entity'"`
}

type ShowEntity struct {
	BaseEntity
	Verbose bool     `cli:"type=opt short=v long=verbose default=true desc='make command more verbose'"`
	Option  string   `cli:"type=opt short=o long=option required=true desc='some option'"`
	Ip      string   `cli:"type=opt short=i default='127.0.0.1' desc='another option'"`
	Name    string   `cli:"type=arg required=true desc='the first argument'"`
	Others  []string `cli:"type=arg desc='some variadic argument'"`
	w       string
	x       string
}

func (se *ShowEntity) Run() error {
	log.Printf("oha, running show entity, with:")
	log.Printf("\tverbose: %t", se.Verbose)
	log.Printf("\toption:  %s", se.Option)
	log.Printf("\tip:      %s", se.Ip)
	log.Printf("\tname:    %s", se.Name)
	log.Printf("\tothers   %s", strings.Join(se.Others, ", "))
	log.Printf("\tA        %s", se.Af)
	log.Printf("\tB        %d", se.Bf)
	log.Printf("\tw        %s", se.w)
	return nil
}

func main() {
	router := cli.NewRouter()

	router.Register("entities/show", &ShowEntity{w: "foog"}, "an example action")
	router.Register("entities/list", &ShowEntity{w: "foog"}, "an example action")
	router.Register("entit/list", &ShowEntity{w: "foog", Option: "barz"}, "an example action")
	router.RegisterFunc("version", func() error { log.Printf("x.y"); return nil }, "show the apps version")

	if e := router.RunWithArgs(); e != nil {
		log.Fatal(e)
	}
}
