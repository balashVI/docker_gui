package main

import (
	"github.com/fsouza/go-dockerclient"
	"gopkg.in/qml.v1"
)

func NewDockerContainers(client *docker.Client) *DockerContainers {
	containers := &DockerContainers{
		client:          client,
		containersCache: map[string]*Container{},
		List:            NewContainersList(client),
	}
	return containers
}

type DockerContainers struct {
	client          *docker.Client
	List            *ContainersList
	containersCache map[string]*Container
}

// Get returns detailed info about container
func (self *DockerContainers) Get(containerId string) *Container {
	res, ok := self.containersCache[containerId]
	if !ok {
		res = NewContainer(containerId, self.client)
		self.containersCache[containerId] = res
	}
	return res
}

func (self *DockerContainers) OnContainerDied(id string) {
	if container, ok := self.containersCache[id]; ok {
		container.Status = "stopped"
		qml.Changed(container, &container.Status)
	}
	self.List.UpdateContainerStatus(id, false)
}

func (self *DockerContainers) OnContainerStarted(id string) {
	if container, ok := self.containersCache[id]; ok {
		container.Status = "running"
		qml.Changed(container, &container.Status)
	}
	self.List.UpdateContainerStatus(id, true)
}

func (self *DockerContainers) OnContainerDestroyed(id string) {
	delete(self.containersCache, id)
	self.List.Remove(id)
}
