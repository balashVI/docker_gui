from gi.repository import Gtk


class TwoColTable(Gtk.ScrolledWindow):
    def __init__(self, title1, title2, items):
        Gtk.ScrolledWindow.__init__(self)

        builder = Gtk.Builder()
        builder.add_from_file('./ui/two_col_table.glade')
        builder.connect_signals(self)
        self.add(builder.get_object('view'))

        # init model
        self.model = builder.get_object('model')
        self.mo
        for item in items:


        # rename columns
        builder.get_object('column1').set_title(title1)
        builder.get_object('column2').set_title(title2)
