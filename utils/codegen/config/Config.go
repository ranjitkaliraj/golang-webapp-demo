package main

type Config interface {
	GetRepositoryConfig() RepositoryConfig
	GetCommandConfig() CommandConfig
	GetQueryConfig() QueryConfig
	GetEventConfig() EventConfig
}

struct RepositoryConfig {
	// configuration for repository code generation.
	entityInfo: RepoEntity
}

struct CommandConfig {
	// configuration for commands code generation.
	entity: string, 
	cmdToEntityMapping: string
}

struct QueryConfig {
	// configuration for queries code generation.
	entity: string, 
}

struct EventConfig {
	// configuration for events code generation.
}