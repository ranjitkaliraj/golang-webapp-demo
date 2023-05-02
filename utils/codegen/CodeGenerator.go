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
	entity := config.entity
	fileName := config.fileName
	tmpl, err := template.New("commandHandler").Parse(CommandHandlerTemplate.template)
	writeToFile(tmpl, err, fileName, entity)
}

func generateQueryCode(config QueryConfig) {
	// code generating queries, query handler etc.
	entity := config.entity
	fileName := config.fileName
	tmpl, err := template.New("queryHandler").Parse(QueryHandlerTemplate.template)
	writeToFile(tmpl, err, fileName, entity)
}

func generateEventsCode(config EventsConfig) {
	// code generating events, event handler etc.
}

// Code generating repository
func generateRepositoryCode(config RepositoryConfig) {
	entity := config.entity
	fileName := config.fileName
	tmpl, err := template.New("repository").Parse(RepositoryTemplate.repoTemplate)
	writeToFile(tmpl, err, fileName, entity)
}

func writeToFile(tmpl *template.Template, err error, fileName string, entity string) {
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