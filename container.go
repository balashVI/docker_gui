package main

import (
	"bytes"
	"github.com/fsouza/go-dockerclient"
	"gopkg.in/qml.v1"
	"strings"
	"time"
)

// NewContainerInfo creates new instance of ContainerInfo.
func NewContainer(containerId string, dockerClient *docker.Client) *Container {
	container, err := dockerClient.InspectContainer(containerId)
	checkErrors(err)
	// base info
	var status string
	if container.State.Running {
		status = "running"
	} else {
		status = "stopped"
	}

	res := &Container{
		dockerClient: dockerClient,
		Id:           container.ID,
		Name:         container.Name,
		Image:        container.Config.Image,
		Created:      container.Created.Format("2006-01-02 15:04:05"),
		Status:       status,
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

	return res
}

type Container struct {
	dockerClient *docker.Client

	Id      string
	Name    string
	Image   string
	Created string
	Status  string
	Env     EnvironmentVariablesList
	Mounts  MountsList
	Ports   PortBindingsList

	log         string
	lastLogSync int64
}

func (self *Container) GetLogs(all bool) string {
	var res bytes.Buffer
	since := self.lastLogSync
	self.lastLogSync = time.Now().Unix()
	self.dockerClient.Logs(docker.LogsOptions{
		Container:    self.Id,
		OutputStream: &res,
		ErrorStream:  &res,
		Stdout:       true,
		Stderr:       true,
		Since:        since,
	})
	self.log = self.log + res.String()
	if all {
		return self.log
	}
	return res.String()
}

type EnvironmentVariablesList struct {
	list []EnvironmentVariable
	Len  int
}

func (self *EnvironmentVariablesList) Add(env []EnvironmentVariable) {
	self.list = append(self.list, env...)
	self.Len = len(self.list)
	qml.Changed(self, &self.Len)
}

func (self *EnvironmentVariablesList) Get(position int) *EnvironmentVariable {
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

type PortBinding struct {
	ContainerPort string
	HostPort      string
}

type PortBindingsList struct {
	list []PortBinding
	Len  int
}

func (self *PortBindingsList) Add(ports []PortBinding) {
	self.list = append(self.list, ports...)
	self.Len = len(self.list)
	qml.Changed(self, &self.Len)
}

func (self *PortBindingsList) Get(position int) *PortBinding {
	return &self.list[position]
}
