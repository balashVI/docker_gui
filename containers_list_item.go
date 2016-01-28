package main

type ContainersListItem struct {
	Id        string // container ID
	Name      string
	Image     string
	IsRunning bool
}
