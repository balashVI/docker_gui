package main

type ContainersListItem struct {
	Id        string
	Name      string
	Image     string
	IsRunning bool
}

type ContainersList struct {
	list []ContainersListItem
	Len  int
}

func (self *ContainersList) Get(position int) *ContainersListItem {
	return &self.list[position]
}

func (self *ContainersList) Clear() {
	self.list = self.list[:0]
	self.Len = 0
	//qml.Changed(self, &self.Len)
}

func (self *ContainersList) Add(containers []ContainersListItem) {
	self.list = append(self.list, containers...)
	self.Len = len(self.list)
}
