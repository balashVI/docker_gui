package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

type EventsMonitor struct {
	dockerClient *docker.Client

	containers *Containers
	images     *Images
}

func (self *EventsMonitor) Run(dockerClient *docker.Client) {
	self.dockerClient = dockerClient

	ch := make(chan *docker.APIEvents)
	self.dockerClient.AddEventListener(ch)
	for {
		event := <-ch
		fmt.Println("id: ", event.ID, ", from: ", event.From, ", status: ", event.Status)

		// events from the containers
		if self.containers != nil {
			switch event.Status {
			case "start":
				self.containers.OnContainerStarted(event.ID)
			case "die":
				self.containers.OnContainerDied(event.ID)
			case "destroy":
				self.containers.OnContainerDestroyed(event.ID)
			}
		}

		// events from the images
		if self.images != nil {

		}
	}
}
