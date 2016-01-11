#!/usr/bin/python3
import gi

gi.require_version('Gtk', '3.0')
from gi.repository import Gtk, GObject
from main_window import MainWindow

if __name__ == '__main__':
    GObject.threads_init()
    win = MainWindow()
    win.connect('delete-event', Gtk.main_quit)
    win.show_all()
    Gtk.main()
