package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"gopkg.in/qml.v1"
	"os"
)

func main() {
	err := qml.Run(run)
	checkErrors(err)
}

func checkErrors(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	dockerClient, err := docker.NewClient("unix:///var/run/docker.sock")
	checkErrors(err)

	engine := qml.NewEngine()

	// configuring context
	images := NewDockerImages(dockerClient)
	engine.Context().SetVar("DockerImages", images)

	containers := NewDockerContainers(dockerClient)
	engine.Context().SetVar("DockerContainers", containers)

	tasks := NewTasksRunner("unix:///var/run/docker.sock")
	engine.Context().SetVar("DockerTasks", tasks)

	eventsMonitor := NewEventsMonitor(dockerClient, containers, images)
	go eventsMonitor.Run()
	engine.Context().SetVar("DockerEvents", eventsMonitor)

	// creating window
	controls, err := engine.LoadFile("./qml/main_window.qml")
	if err != nil {
		return err
	}
	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}
