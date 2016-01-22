package main

import "github.com/fsouza/go-dockerclient"

type TasksRunner struct {
	dockerClient *docker.Client
	timeout      uint
}

func (self *TasksRunner) Init(endpoint string) {
	var err error
	self.dockerClient, err = docker.NewClient(endpoint)
	checkErrors(err)

	// time for wait before killing container
	self.timeout = 60
}

func (self *TasksRunner) StartContainer(id string) {
	go func() {
		self.dockerClient.StartContainer(id, &docker.HostConfig{})
	}()
}

func (self *TasksRunner) StopContainer(id string) {
	go func() {
		self.dockerClient.StopContainer(id, self.timeout)
	}()
}

func (self *TasksRunner) DeleteContainer(id string) {
	go func() {
		self.dockerClient.RemoveContainer(docker.RemoveContainerOptions{ID: id})
	}()
}
