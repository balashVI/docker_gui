import threading
from gi.repository import Gtk, GObject

from docker_client import cli


class DeleteImageDialog(Gtk.Window):
    def __init__(self, parent, image_id):
        Gtk.Window.__init__(self)
        # налаштування вікна
        self.set_transient_for(parent)
        self.set_modal(True)
        self.set_type_hint(Gtk.WindowType.POPUP)
        self.set_default_size(150, 100)
        self.set_border_width(10)
        self.set_title('')
        self.set_resizable(False)

        self.image_id = image_id

        builder = Gtk.Builder()
        builder.add_from_file('./ui/dialog_delete_image.glade')
        builder.connect_signals({
            'on_close': self.on_close,
            'on_delete': self.on_delete,
        })
        self.stack = builder.get_object('stack')
        self.add(self.stack)

    def on_close(self, event):
        self.close()

    def on_delete(self, event):
        self.stack.set_visible_child_name('process')
        thread = threading.Thread(target=self.delete_image,
                                  args=(self.image_id, self.finish))
        thread.daemon = True
        thread.start()

    def finish(self):
        self.stack.set_visible_child_name('finish')

    @staticmethod
    def delete_image(image_id, finish):
        cli.remove_image(image_id, force=True)
        GObject.idle_add(finish)
