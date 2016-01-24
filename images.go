package main

import (
	"github.com/fsouza/go-dockerclient"
	"strings"
)

func NewImages(dockerClient *docker.Client) *Images {
	images := &Images{
		dockerClient: dockerClient,
	}
	images.UpdateList()
	return images
}

type Images struct {
	dockerClient *docker.Client
	List         ImagesList
}

func (self *Images) UpdateList() {
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
