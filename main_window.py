from gi.repository import Gtk

from header_bar import HeaderBar
from images import Images
from containers import Containers
from info import Info


class MainWindow(Gtk.Window):
    def __init__(self):
        Gtk.Window.__init__(self)
        self.set_size_request(600, 400)
        self.stack = Gtk.Stack()

        self.add(self.stack)

        # заголовок
        self.header = HeaderBar()
        self.header.stack_switcher.set_stack(self.stack)
        self.header.btn_go_back.connect('clicked', self.on_go_back)
        self.set_titlebar(self.header)

        # образи
        images = Images()
        self.stack.add_titled(images, 'images', 'Images')
        images.connect('notify::visible-child', self.on_visible_child_changed)

        # контейнери
        containers = Containers()
        self.stack.add_titled(containers, 'containers', 'Containers')
        containers.connect('notify::visible-child', self.on_visible_child_changed)

        # сторінка інформації
        self.stack.add_titled(Info(), 'info', 'Info')

        self.stack.connect('notify::visible-child', self.on_visible_child_changed)

    def on_go_back(self, widget):
        self.stack.get_visible_child().go_back()

    def on_visible_child_changed(self, widget, param):
        current_child = self.stack.get_visible_child()
        if current_child and current_child.get_visible_child_name() == 'first':
            self.header.btn_go_back.set_sensitive(False)
        else:
            self.header.btn_go_back.set_sensitive(True)
        return True
