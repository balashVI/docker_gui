from gi.repository import Gtk

from store import store
from dialog_new_image import NewImageDialog


class ImagesList(Gtk.ScrolledWindow):
    def __init__(self):
        Gtk.ScrolledWindow.__init__(self)

        # create builder and load ui file
        builder = Gtk.Builder()
        builder.add_from_file('./ui/images_list.glade')
        builder.connect_signals(self)

        # view
        self.images_view = builder.get_object('view')
        self.images_view.set_model(store.images_store)
        self.add(self.images_view)

        # context menu
        self.context_menu = builder.get_object('menu')
        self.delete_image_menu = builder.get_object('delete_image')

    def on_update_list_activate(self, event):
        store.update_images_store()

    def on_list_view_btn_press(self, widget, event):
        if event.button == 3:
            # check selected rows
            show_delete_image_menu = self.images_view.get_selection().count_selected_rows()
            self.delete_image_menu.set_sensitive(show_delete_image_menu)

            self.context_menu.popup(None, None, None, None, event.button, event.time)

    def on_delete_image_activate(self, widget):
        # отримання виділеного рядка
        images_store, selected = self.images_view.get_selection().get_selected_rows()
        row = images_store[selected[0]]

        # відображення вікна видалення образу
        dialog = Gtk.MessageDialog(self.get_toplevel(), 0, Gtk.MessageType.QUESTION,
                                   Gtk.ButtonsType.YES_NO, 'Deleting image')
        dialog.format_secondary_text('Do you really want to delete selected image?')
        response = dialog.run()
        if response == Gtk.ResponseType.YES:
            store.add_delete_image_task(row[0])

        dialog.destroy()

    def on_new_image_activate(self, widget):
        dialog = NewImageDialog(self.get_toplevel())
        dialog.show_all()
