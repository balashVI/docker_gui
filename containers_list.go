package main

import (
	"github.com/fsouza/go-dockerclient"
	"gopkg.in/qml.v1"
	"strings"
)

func NewContainersList(client *docker.Client) *ContainersList {
	res := &ContainersList{
		client: client,
	}
	res.Reload()
	return res
}

// ContainersList represents list of short information about containers
type ContainersList struct {
	client *docker.Client
	list   []ContainersListItem
	Len    int
}

// Get returns short info about container by position in the list
func (self *ContainersList) Get(position int) *ContainersListItem {
	return &self.list[position]
}

// Reload list
func (self *ContainersList) Reload() {
	response, err := self.client.ListContainers(docker.ListContainersOptions{All: true})
	checkErrors(err)

	containers := make([]ContainersListItem, len(response))

	var isRunning bool
	for i, con := range response {
		if con.Status[0] == 'U' {
			isRunning = true
		} else {
			isRunning = false
		}

		containers[i] = ContainersListItem{
			Id:        con.ID,
			Name:      strings.Join(con.Names, ", "),
			Image:     con.Image,
			IsRunning: isRunning,
		}
	}

	self.list = containers
	self.Len = len(self.list)

	qml.Changed(self, &self.Len)
}

func (self *ContainersList) Remove(id string) {
	for i, _ := range self.list {
		if self.list[i].Id == id {
			self.list = append(self.list[:i], self.list[i+1:]...)
			self.Len = self.Len - 1
			qml.Changed(self, &self.Len)
		}
	}
}

func (self *ContainersList) UpdateContainerStatus(id string, isRunning bool) {
	for i, _ := range self.list {
		if self.list[i].Id == id {
			self.list[i].IsRunning = isRunning
			qml.Changed(&self.list[i], &self.list[i].IsRunning)
			break
		}
	}
}
