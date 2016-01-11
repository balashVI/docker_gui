from datetime import datetime
from gi.repository import Gtk

from docker_client import cli
from dialog_delete_image import DeleteImageDialog


class ImagesList(Gtk.ScrolledWindow):
    def __init__(self):
        Gtk.ScrolledWindow.__init__(self)

        # create builder and load ui file
        builder = Gtk.Builder()
        builder.add_from_file('./ui/images_list.glade')
        builder.connect_signals({
            'on_list_view_btn_press': self.show_context_menu,
            'on_update_list': self.update_store,
            'on_delete_image': self.show_delete_image_dialog
        })

        # view
        self.images_view = builder.get_object('view')
        self.add(self.images_view)

        # store
        self.images_store = builder.get_object('store')
        self.update_store()

        # context menu
        self.context_menu = builder.get_object('menu')
        self.delete_image_menu = builder.get_object('delete_image')

    def update_store(self, event=None):
        self.images_store.clear()
        images = cli.images()
        for image in images:
            repo_tags = image['RepoTags'][0].split(':')
            repo, tags = repo_tags[0], repo_tags[1]
            created = str(datetime.fromtimestamp(image['Created']))
            self.images_store.append([image['Id'], repo, tags, created])
        return True

    def show_context_menu(self, widget, event):
        if event.button == 3:
            # check selected rows
            show_delete_image_menu = self.images_view.get_selection().count_selected_rows()
            self.delete_image_menu.set_sensitive(show_delete_image_menu)

            self.context_menu.popup(None, None, None, None, event.button, event.time)

    def show_delete_image_dialog(self, event):
        # отримання виділеного рядка
        _, selected = self.images_view.get_selection().get_selected_rows()
        row = self.images_store[selected[0]]

        # відображення вікна видалення образу
        dialog = DeleteImageDialog(self.get_toplevel(), row[0])
        dialog.connect('destroy', self.update_store)
        dialog.show_all()
        return True
