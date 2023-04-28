package main

type ContactConfig struct{}

func New() {
	return &ContactConfig{}
}

// Configuration for Repository code generation of Contact.
func (a ContactConfig) GetRepositoryConfig() Config.RepositoryConfig {
	entity := RepoEntity{
		Name:    "Contact",
		Package: "repository",
		DBName:  "contacts",
		Fields: []RepoField{
			{Name: "ID", Type: "uuid.UUID", Tag: "`bson:\"_id,omitempty\"`", IsPrimary: true},
			{Name: "FirstName", Type: "string", Tag: "`bson:\"first_name\"`", IsPrimary: false},
			{Name: "LastName", Type: "string", Tag: "`bson:\"last_name\"`", IsPrimary: false},
		},
		Imports: []string{
			`"context"`,
			`"github.com/google/uuid"`,
			`"go.mongodb.org/mongo-driver/bson"`,
			`"go.mongodb.org/mongo-driver/bson/primitive"`,
			`"go.mongodb.org/mongo-driver/mongo"`,
			`"go.mongodb.org/mongo-driver/mongo/options"`,
		},
	}
	return RepositoryConfig{entity}
}

// Configuration for Command code generation of Contact.
func (a ContactConfig) GetCommandConfig() Config.CommandConfig {
	// Define commands code generation configuration here.
	return CommandConfig{}
}

// Configuration for Query code generation of Contact.
func (a ContactConfig) GetQueryConfig() Config.QueryConfig {
	// Define query code generation configuration here.
	return QueryConfig{}
}

// Configuration for Events code generation of Contact.
func (a ContactConfig) GetEventConfig() Config.CommandConfig {
	// Define events code generation configuration here.
	return QueryConfig{}
}
