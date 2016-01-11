from gi.repository import Gtk


class HeaderBar(Gtk.HeaderBar):
    def __init__(self):
        Gtk.HeaderBar.__init__(self)

        # показати кнопки керування вікном
        self.set_show_close_button(True)

        # кнопка назад
        self.btn_back = Gtk.Button(image=Gtk.Image(stock=Gtk.STOCK_GO_BACK))
        self.btn_back.set_sensitive(False)
        self.pack_start(self.btn_back)

        # кнопки перемикання сторінок Gtk.Stack
        self.stack_switcher = Gtk.StackSwitcher()
        self.set_custom_title(self.stack_switcher)

        # кнопка оновити
        btn_update = Gtk.Button(image=Gtk.Image(stock=Gtk.STOCK_REFRESH))
        self.pack_end(btn_update)

        # кнопка додати
        btn_add = Gtk.Button(image=Gtk.Image(stock=Gtk.STOCK_ADD))
        self.pack_end(btn_add)
