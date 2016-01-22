package main

import (
	"gopkg.in/qml.v1"
)

type ImagesListItem struct {
	Id   string
	Name string
	Tag  string
}

type ImagesList struct {
	images []ImagesListItem
	Len    int
}

func (self *ImagesList) Get(index int) *ImagesListItem {
	return &self.images[index]
}

func (self *ImagesList) Clear() {
	self.images = self.images[:0]
	self.Len = 0
	qml.Changed(self, &self.Len)
}

func (self *ImagesList) Add(images []ImagesListItem) {
	self.images = append(self.images, images...)
	self.Len = len(self.images)
	qml.Changed(self, &self.Len)
}
