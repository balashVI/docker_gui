package main

import (
	"github.com/fsouza/go-dockerclient"
)

func NewTasksRunner(dockerEndpoint string) *TasksRunner {
	dockerClient, err := docker.NewClient(dockerEndpoint)
	checkErrors(err)
	tasksRunner := &TasksRunner{
		dockerClient: dockerClient,
	}
	return tasksRunner
}

type TasksRunner struct {
	dockerClient *docker.Client
	timeout      uint
}

func (self *TasksRunner) StartContainer(id string) {
	go func() {
		err := self.dockerClient.StartContainer(id, &docker.HostConfig{})
		checkErrors(err)
	}()
}

func (self *TasksRunner) StopContainer(id string) {
	go func() {
		err := self.dockerClient.StopContainer(id, self.timeout)
		checkErrors(err)
	}()
}

func (self *TasksRunner) DeleteContainer(id string) {
	go func() {
		self.dockerClient.RemoveContainer(docker.RemoveContainerOptions{ID: id})
	}()
}
