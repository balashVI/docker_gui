from gi.repository import Gtk

from header_bar import HeaderBar
from images_list import ImagesList
from containers_list import ContainersList
from info import Info


class MainWindow(Gtk.Window):
    def __init__(self):
        Gtk.Window.__init__(self)
        self.set_size_request(600, 400)
        stack = Gtk.Stack()
        self.add(stack)

        # заголовок
        self.header = HeaderBar()
        self.header.stack_switcher.set_stack(stack)
        self.set_titlebar(self.header)

        # список образів
        stack.add_titled(ImagesList(), 'images', 'Images')

        # список контейнерів
        stack.add_titled(ContainersList(), 'containers', 'Containers')

        # сторінка інформації
        stack.add_titled(Info(), 'info', 'Info')
