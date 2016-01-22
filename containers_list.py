from gi.repository import Gtk, GObject
from store import store


class ContainersList(Gtk.ScrolledWindow):
    __gsignals__ = {
        'container_activated': (GObject.SIGNAL_RUN_FIRST, None, (str,))
    }

    def __init__(self):
        Gtk.ScrolledWindow.__init__(self)

        builder = Gtk.Builder()
        builder.add_from_file('./ui/containers_list.glade')
        builder.connect_signals(self)

        self.view = builder.get_object('view')
        self.view.set_model(store.containers_store)
        self.add(self.view)

        self.context_menu = builder.get_object('view_menu')

    def on_update_list_activate(self, widget):
        store.update_containers_store()

    def on_view_button_press_event(self, widget, event):
        if event.button == 3:
            # check selected rows
            # show_delete_image_menu = self.images_view.get_selection().count_selected_rows()
            # self.delete_image_menu.set_sensitive(show_delete_image_menu)

            self.context_menu.popup(
                    None, None, None, None, event.button, event.time)

    def on_view_row_activated(self, view, rows, column):
        container_id = store.containers_store[rows[0]][0]
        self.emit('container_activated', container_id)
