from gi.repository import Gtk
from images_list import ImagesList


class Images(Gtk.Stack):
    def __init__(self):
        Gtk.Stack.__init__(self)

        # images list
        self.images_list = ImagesList()
        self.add_named(self.images_list, 'first')
