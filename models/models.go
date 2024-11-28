package models

type Id string

type User struct {
	FirstName string
	LastName  string
	FullName  string
	Biography string
}

type Application struct {
	Data map[Id]User
}
