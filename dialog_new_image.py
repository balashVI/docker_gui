import threading
from gi.repository import Gtk

from docker_client import cli


class NewImageDialog(Gtk.Window):
    def __init__(self, parent):
        Gtk.Window.__init__(self)
        # налаштування вікна
        self.set_transient_for(parent)
        self.set_modal(True)
        # self.set_type_hint(Gtk.WindowType.POPUP)
        self.set_default_size(800, 600)
        self.set_border_width(10)
        self.set_title('Add new image from DockerHub')

        builder = Gtk.Builder()
        builder.add_from_file('./ui/dialog_new_image.glade')
        builder.connect_signals(self)
        self.stack = builder.get_object('stack')
        self.add(self.stack)

        self.search_entry = builder.get_object('search')
        """ :type : Gtk.SearchEntry """
        self.store = builder.get_object('store')
        """ :type : Gtk.ListStore """
        self.add_btn = builder.get_object('add_btn')
        """ :type : Gtk.Button """

    def on_close(self, widget):
        self.close()

    def on_search(self, widget):
        query = self.search_entry.get_text()
        res = cli.search(query)
        self.store.clear()
        for i, row in enumerate(res):
            self.store.append([i, row['name'], row['star_count'],
                               row['is_official'], row['is_automated'],
                               row['description']])

    def on_selected_changed(self, selection):
        """
        :type selection: Gtk.TreeSelection
        """
        _, selected = selection.get_selected_rows()
        if (len(selected) > 0) and (self.add_btn.get_sensitive() is False):
            self.add_btn.set_sensitive(True)
        elif (len(selected) == 0) and (self.add_btn.get_sensitive() is True):
            self.add_btn.set_sensitive(False)

    def on_add(self, selection):
        """
        :type selection: Gtk.TreeSelection
        """
        print(selection.get_selected_rows()[1][0])
