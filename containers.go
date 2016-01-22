package main

import (
	"github.com/fsouza/go-dockerclient"
	"gopkg.in/qml.v1"
	"strings"
)

type Containers struct {
	dockerClient *docker.Client
	List         ContainersList
	catch        map[string]*ContainerInfo
}

func (self *Containers) Init(dockerClient *docker.Client) {
	self.dockerClient = dockerClient
	self.catch = map[string]*ContainerInfo{}

	self.UpdateList()
}

func (self *Containers) UpdateList() {
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
		containers[i] = ContainersListItem{
			Id:        con.ID,
			Name:      con.Names[0],
			Image:     con.Image,
			IsRunning: isRunning,
		}
	}

	self.List.Clear()
	self.List.Add(containers)
	qml.Changed(&self.List, &self.List.Len)
}

func (self *Containers) Inspect(container_id string) *ContainerInfo {
	if container_id == "" {
		return nil
	}
	res, ok := self.catch[container_id]
	if !ok {
		container, err := self.dockerClient.InspectContainer(container_id)
		checkErrors(err)
		// base info
		res = &ContainerInfo{
			Id:      container_id,
			Name:    container.Name,
			Image:   container.Config.Image,
			Created: container.Created.Format("2006-01-02 15:04:05"),
			Running: container.State.Running,
		}

		// env
		env := make([]EnvironmentVariable, len(container.Config.Env))
		for i, j := range container.Config.Env {
			keyVal := strings.Split(j, "=")
			env[i].Key = keyVal[0]
			env[i].Value = keyVal[1]
		}
		res.Env.Add(env)

		// mounts
		mounts := make([]Mounts, len(container.Mounts))
		for i, j := range container.Mounts {
			mounts[i].Destination = j.Destination
			mounts[i].Source = j.Source
		}
		res.Mounts.Add(mounts)

		// ports
		ports := []PortBinding{}
		for key, value := range container.HostConfig.PortBindings {
			for i, _ := range value {
				ports = append(ports, PortBinding{
					ContainerPort: string(key),
					HostPort:      string(value[i].HostPort),
				})
			}
		}
		res.Ports.Add(ports)

		self.catch[container_id] = res
	}

	return res
}

func (self *Containers) OnContainerDied(id string) {
	if container, ok := self.catch[id]; ok {
		container.Running = false
		qml.Changed(container, &container.Running)
	}

	for i, _ := range self.List.list {
		if self.List.list[i].Id == id {
			self.List.list[i].IsRunning = false
			qml.Changed(&self.List.list[i], &self.List.list[i].IsRunning)
			break
		}
	}
}

func (self *Containers) OnContainerStarted(id string) {
	if container, ok := self.catch[id]; ok {
		container.Running = true
		qml.Changed(container, &container.Running)
	}

	for i, _ := range self.List.list {
		if self.List.list[i].Id == id {
			self.List.list[i].IsRunning = true
			qml.Changed(&self.List.list[i], &self.List.list[i].IsRunning)
			break
		}
	}
}

func (self *Containers) OnContainerDestroyed(id string) {
	delete(self.catch, id)
	self.UpdateList()
}
