package main

import (
	"encoding/json"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"strings"
)

func NewDockerImages(dockerClient *docker.Client) *DockerImages {
	images := &DockerImages{
		dockerClient: dockerClient,
	}
	images.UpdateList()
	return images
}

type DockerImages struct {
	dockerClient *docker.Client
	List         ImagesList
}

func (self *DockerImages) UpdateList() {
	response, err := self.dockerClient.ListImages(docker.ListImagesOptions{All: false})
	checkErrors(err)

	images := make([]ImagesListItem, len(response))

	for i, img := range response {
		repoTag := strings.Split(img.RepoTags[0], ":")
		images[i] = ImagesListItem{
			Id:   img.ID,
			Name: repoTag[0],
			Tag:  repoTag[1],
		}
	}

	self.List.Clear()
	self.List.Add(images)
}

func (self *DockerImages) Get(id string) {
	image, _ := self.dockerClient.InspectImage("percona")
	//checkErrors(err)
	res, _ := json.MarshalIndent(image, "", "   ")
	fmt.Println(string(res))
}
