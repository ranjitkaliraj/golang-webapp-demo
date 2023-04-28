package main

import (
	"os"
	"text/template"
)

func generate() {
	contactConfig := ContactConfig.New
	process(contactConfig)

	// Add for other entities e.g. account, leads, opportunities etc. 
}

func process(config Configurable) {
	repoConfig := config.GetRepositoryConfig()
	cmdConfig := config.GetCommandConfig()
	queryConfig := config.GetQueryConfig()
	eventsConfig := config.GetQueryConfig()

	generateRepositoryCode(repoConfig)
	generateCommandsCode(cmdConfig)
	generateQueryCode(queryConfig)
	generateEventsCode(eventsConfig)
}

func generateCommandsCode(config CommandsConfig) {
	// code generating commands, command handler etc.
}

func generateQueryCode(config QueryConfig) {
	// code generating queries, query handler etc.
}

func generateEventsCode(config EventsConfig) {
	// code generating events, event handler etc.
}

// Code generating repository
func generateRepositoryCode(config RepositoryConfig) {
	entity := config.entity
	fileName := config.fileName
	tmpl, err := template.New("repository").Parse(RepositoryTemplate.repoTemplate)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, entity)
	if err != nil {
		panic(err)
	}
}
