from gi.repository import Gtk

from docker_client import cli


class Info(Gtk.ScrolledWindow):
    def __init__(self):
        Gtk.ScrolledWindow.__init__(self)
        self.box = Gtk.Box(orientation=Gtk.Orientation.VERTICAL)
        self.add(self.box)

        self.update()

    def update(self):
        # видалення старого вмісту
        children = self.box.get_children()
        for ch in children:
            self.box.remove(ch)

        # додавання інформації
        info = cli.info()
        for key, value in info.items():
            label = Gtk.Label(key + ': ' + str(value), xalign=0)
            self.box.pack_start(label, True, True, 0)
