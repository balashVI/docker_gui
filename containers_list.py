from datetime import datetime
from gi.repository import Gtk

from docker_client import cli


class ContainersList(Gtk.ScrolledWindow):
    def __init__(self):
        Gtk.ScrolledWindow.__init__(self)

        # store
        self.containers_store = Gtk.ListStore(str, str, str)
        self.update_store()

        # view
        tree_view = Gtk.TreeView(self.containers_store)
        tree_view.set_grid_lines(Gtk.TreeViewGridLines.HORIZONTAL)
        self.add(tree_view)

        # налаштування стовпців
        for i, column_title in enumerate(['Status', 'Name', 'Image']):
            renderer = Gtk.CellRendererText()
            column = Gtk.TreeViewColumn(column_title, renderer, text=i)
            column.set_sort_column_id(i)
            tree_view.append_column(column)
        tree_view.get_column(1).set_expand(True)

    def update_store(self):
        containers = cli.containers(all=True)
        for container in containers:
            # print(container)
            status = container['Status'].split()[0]
            self.containers_store.append([status, container['Names'][0], container['Image']])
