package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"gopkg.in/qml.v1"
)

func NewEventsMonitor(dockerClient *docker.Client, containers *DockerContainers, images *Images) *EventsMonitor {
	eventsMonitor := &EventsMonitor{
		dockerClient: dockerClient,
		containers:   containers,
		images:       images,
	}
	return eventsMonitor
}

type EventsMonitor struct {
	dockerClient *docker.Client

	containers *DockerContainers
	images     *Images

	LastEvent string
}

func (self *EventsMonitor) event(ev string) {
	self.LastEvent = ev
	qml.Changed(self, &self.LastEvent)
}

func (self *EventsMonitor) Run() {
	ch := make(chan *docker.APIEvents)
	self.dockerClient.AddEventListener(ch)
	for event := range ch {
		fmt.Println("id: ", event.ID, ", from: ", event.From, ", status: ", event.Status)

		// events from the containers
		if self.containers != nil {
			switch event.Status {
			case "start":
				self.event("Started container with ID " + event.ID)
				self.containers.OnContainerStarted(event.ID)
			case "die":
				self.event("Stoped container with ID " + event.ID)
				self.containers.OnContainerDied(event.ID)
			case "destroy":
				self.event("Destroyed container with ID " + event.ID)
				self.containers.OnContainerDestroyed(event.ID)
			}
		}

		// events from the images
		if self.images != nil {

		}
	}
}
