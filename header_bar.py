from gi.repository import Gtk


class HeaderBar(Gtk.HeaderBar):
    def __init__(self):
        Gtk.HeaderBar.__init__(self)

        # показати кнопки керування вікном
        self.set_show_close_button(True)

        # кнопка назад
        self.btn_go_back = Gtk.Button(image=Gtk.Image(stock=Gtk.STOCK_GO_BACK))
        self.btn_go_back.set_sensitive(False)
        self.pack_start(self.btn_go_back)

        # кнопки перемикання сторінок Gtk.Stack
        self.stack_switcher = Gtk.StackSwitcher()
        self.set_custom_title(self.stack_switcher)

        # кнопка процесу
        self.btn_background = Gtk.Button()
        self.btn_background.add(Gtk.Spinner(active=True))
        self.pack_end(self.btn_background)
        self.btn_background_popover_label = Gtk.Label('Background task')
        self.btn_background_popover = Gtk.Popover()
        self.btn_background_popover.set_border_width(10)
        self.btn_background_popover.set_relative_to(self.btn_background)
        self.btn_background_popover.add(self.btn_background_popover_label)
        self.btn_background.connect('clicked', self.on_background_btn_clicked)

        # кнопка оновити
        btn_update = Gtk.Button(image=Gtk.Image(stock=Gtk.STOCK_REFRESH))
        self.pack_end(btn_update)

        # кнопка додати
        btn_add = Gtk.Button(image=Gtk.Image(stock=Gtk.STOCK_ADD))
        self.pack_end(btn_add)

    def on_background_btn_clicked(self, event):
        self.btn_background_popover.show_all()
