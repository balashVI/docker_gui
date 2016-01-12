from datetime import datetime
from gi.repository import Gtk

from docker_client import cli
from store import store


class ContainersList(Gtk.ScrolledWindow):
    def __init__(self):
        Gtk.ScrolledWindow.__init__(self)

        # view
        tree_view = Gtk.TreeView(store.containers_store)
        tree_view.set_grid_lines(Gtk.TreeViewGridLines.HORIZONTAL)
        self.add(tree_view)

        # налаштування стовпців
        for i, column_title in enumerate(['Status', 'Name', 'Image']):
            renderer = Gtk.CellRendererText()
            column = Gtk.TreeViewColumn(column_title, renderer, text=i)
            column.set_sort_column_id(i)
            tree_view.append_column(column)
        tree_view.get_column(1).set_expand(True)
