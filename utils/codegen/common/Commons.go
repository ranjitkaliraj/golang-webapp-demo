package main

type RepoEntity struct {
	Name        string
	Package     string
	DBName      string
	Fields      []RepoField
	Imports     []string
}

type RepoField struct {
	Name      string
	Type      string
	Tag       string
	IsPrimary bool
}