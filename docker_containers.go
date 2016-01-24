package main

import (
	"github.com/fsouza/go-dockerclient"
	"gopkg.in/qml.v1"
)

func NewContainers(dockerClient *docker.Client) *DockerContainers {
	containers := &DockerContainers{
		dockerClient: dockerClient,
	}
	containers.containersCache = map[string]*Container{}
	containers.UpdateList()
	return containers
}

type DockerContainers struct {
	dockerClient    *docker.Client
	List            ContainersList
	containersCache map[string]*Container
}

func (self *DockerContainers) UpdateList() {
	response, err := self.dockerClient.ListContainers(docker.ListContainersOptions{All: true})
	checkErrors(err)

	containers := make([]ContainersListItem, len(response))

	var isRunning bool
	for i, con := range response {
		if con.Status[0] == 'U' {
			isRunning = true
		} else {
			isRunning = false
		}

		convName := func(names []string) string {
			if len(names) == 0 {
				return ""
			}
			res := ""
			for _, name := range names {
				res = res + ", " + name
			}
			return res[1:]
		}

		containers[i] = ContainersListItem{
			Id:        con.ID,
			Name:      convName(con.Names),
			Image:     con.Image,
			IsRunning: isRunning,
		}
	}

	self.List.Clear()
	self.List.Add(containers)
	qml.Changed(&self.List, &self.List.Len)
}

// Inspect returns detailed info about container
func (self *DockerContainers) Inspect(containerId string) *Container {
	res, ok := self.containersCache[containerId]
	if !ok {
		res = NewContainer(containerId, self.dockerClient)
		self.containersCache[containerId] = res
	}
	return res
}

func (self *DockerContainers) OnContainerDied(id string) {
	if container, ok := self.containersCache[id]; ok {
		container.Status = "stopped"
		qml.Changed(container, &container.Status)
	}

	for i, _ := range self.List.list {
		if self.List.list[i].Id == id {
			self.List.list[i].IsRunning = false
			qml.Changed(&self.List.list[i], &self.List.list[i].IsRunning)
			break
		}
	}
}

func (self *DockerContainers) OnContainerStarted(id string) {
	if container, ok := self.containersCache[id]; ok {
		container.Status = "running"
		qml.Changed(container, &container.Status)
	}

	for i, _ := range self.List.list {
		if self.List.list[i].Id == id {
			self.List.list[i].IsRunning = true
			qml.Changed(&self.List.list[i], &self.List.list[i].IsRunning)
			break
		}
	}
}

func (self *DockerContainers) OnContainerDestroyed(id string) {
	delete(self.containersCache, id)
	self.UpdateList()
}
