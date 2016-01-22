from gi.repository import Gtk

from store import store


class ContainerInfo(Gtk.Bin):
    def __init__(self):
        Gtk.Bin.__init__(self)
        self.set_border_width(5)

        builder = Gtk.Builder()
        builder.add_from_file('./ui/container_info.glade')
        builder.connect_signals(self)

        self.add(builder.get_object('container_info'))

        # labels
        self.name_label = builder.get_object('name_label')
        self.image_label = builder.get_object('image_label')
        self.status_label = builder.get_object('status_label')

        # stores
        self.ev_store = builder.get_object('ev_store')
        """ :type : Gtk.ListStore """
        self.mounts_store = builder.get_object('mounts_store')
        """ :type : Gtk.ListStore """

    def update_info(self, container_id):
        container_info = store.get_container_info(container_id)
        self.name_label.set_label(
            '<span size="xx-large">' + container_info['Name'] + '</span>')
        self.image_label.set_label(container_info['Config']['Image'])
        for i in container_info:
            print(i, container_info[i], '\n')

        # status
        status = container_info['State']['Status'].upper()
        if status[0] == 'E':
            status = '<span color="red">'+status+'</span>'
        elif status[0] == 'R':
            status = '<span color="green">'+status+'</span>'
        self.status_label.set_label(status)

        # environment variables store
        self.ev_store.clear()
        env = container_info['Config']['Env']
        for i in env:
            i = i.split('=')
            self.ev_store.append(i)

        # mounts store
        self.mounts_store.clear()
        mounts = container_info['Mounts']
        for i in mounts:
            self.mounts_store.append([
                i['Destination'],
                i['Source'],
                i['Mode']
            ])
