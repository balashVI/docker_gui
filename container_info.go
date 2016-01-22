package main

import (
	"fmt"
	"gopkg.in/qml.v1"
)

type ContainerInfo struct {
	Id      string
	Name    string
	Image   string
	Created string
	Running bool
	Env     EnvironmentVariableList
	Mounts  MountsList
}

type EnvironmentVariableList struct {
	list []EnvironmentVariable
	Len  int
}

func (self *EnvironmentVariableList) Add(env []EnvironmentVariable) {
	self.list = append(self.list, env...)
	self.Len = len(self.list)
	qml.Changed(self, &self.Len)
}

func (self *EnvironmentVariableList) Get(position int) *EnvironmentVariable {
	if position >= self.Len {
		fmt.Println("errrrr", position, self.Len)
	}
	return &self.list[position]
}

type EnvironmentVariable struct {
	Key   string
	Value string
}

type Mounts struct {
	Source      string
	Destination string
}

type MountsList struct {
	list []Mounts
	Len  int
}

func (self *MountsList) Add(env []Mounts) {
	self.list = append(self.list, env...)
	self.Len = len(self.list)
	qml.Changed(self, &self.Len)
}

func (self *MountsList) Get(position int) *Mounts {
	return &self.list[position]
}
