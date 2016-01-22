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
	qml.RegisterTypes("DockerGUI", 1, 0, []qml.TypeSpec{
		{
			Init: func(v *Containers, obj qml.Object) {
				v.Init(dockerClient)
			},
			Singleton: true,
		},
		{
			Init: func(v *Images, obj qml.Object) {
				v.Init(dockerClient)
			},
			Singleton: true,
		},
	})

	engine := qml.NewEngine()

	controls, err := engine.LoadFile("./qml/main_window.qml")
	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}
