from gi.repository import Gtk, GObject

from store import store
from containers_list import ContainersList
from container_info import ContainerInfo


class Containers(Gtk.Stack):
    def __init__(self):
        Gtk.Stack.__init__(self)

        # containers list
        self.containers_list = ContainersList()
        self.add_named(self.containers_list, 'first')
        self.containers_list.connect(
                'container_activated', self.on_container_activated)

        # container info
        self.container_info = ContainerInfo()
        self.add_named(self.container_info, 'second')

    def on_container_activated(self, widget, container_id):
        self.container_info.update_info(container_id)
        self.set_visible_child_name('second')

    def go_back(self):
        self.set_visible_child_name('first')
