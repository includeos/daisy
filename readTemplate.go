package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type daisyTemplate struct {
	GwLeftNet      string
	GwLeftNetmask  string
	LeftAddress    string
	GwRightNet     string
	GwRightNetmask string
	RightAddress   string
	LeftPort       string
	NextHopAddress string
	NextHopPort    string
}

func parse(config *daisyTemplate) {
	f, err := ioutil.ReadFile("template.nacl")
	if err != nil {
		log.Print(err)
		return
	}

	// Parse requires a string
	t, err := template.New("test").Parse(string(f))
	if err != nil {
		log.Print(err)
		return
	}

	err = t.Execute(os.Stdout, config)
	if err != nil {
		log.Print(err)
	}
}
