package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type daisyTemplate struct {
	GwLeftNet      string
	GwLeftNetmask  string
	LeftAddress    string
	GwRightNet     string
	GwRightNetmask string
	RightAddress   string
	LeftPort       int
	NextHopAddress string
	NextHopPort    int
}

func parse(config daisyTemplate, fileName string) error {
	f, err := ioutil.ReadFile("template.nacl")
	if err != nil {
		return fmt.Errorf("error reading template: %v", err)
	}

	// Parse requires a string
	t, err := template.New("test").Parse(string(f))
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	if err = os.MkdirAll("nacls", 0755); err != nil {
		return fmt.Errorf("error creating nacls folder: %v", err)
	}

	fullPath := fmt.Sprintf("nacls/%s.nacl", fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}

	err = t.Execute(file, config)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
