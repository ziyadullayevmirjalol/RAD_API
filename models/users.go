package models

type User struct {
	Id int
	Firstname string
	Lastname string
	EmailUsername string
	Password string
	Notes []Notes
	Task []Task
}