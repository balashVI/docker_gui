from gi.repository import GObject, Gtk
from datetime import datetime
import threading

from docker_client import cli


class _Store(GObject.GObject):
    def __init__(self):
        GObject.GObject.__init__(self)

        self.images_store = Gtk.ListStore(str, str, str, str)
        self.update_images_store()

        self.containers_store = Gtk.ListStore(str, str, str)
        self.update_containers_store()

        # списки та контейнери що в процесі видалення
        self.deleted_images = set()
        self.deleted_containers = set()

    def update_images_store(self):
        self.images_store.clear()
        images = cli.images()
        for image in images:
            repo_tags = image['RepoTags'][0].split(':')
            repo, tags = repo_tags[0], repo_tags[1]
            created = str(datetime.fromtimestamp(image['Created']))
            self.images_store.append([image['Id'], repo, tags, created])

    def update_containers_store(self):
        self.containers_store.clear()
        containers = cli.containers(all=True)
        for container in containers:
            # print(container)
            status = container['Status'].split()[0]
            self.containers_store.append([status, container['Names'][0], container['Image']])

    def add_delete_image_task(self, image_id):
        # перевірка чи зображення вже видаляється
        if image_id not in self.deleted_images:
            self.deleted_images.add(image_id)
            thread = threading.Thread(target=self.delete_image,
                                      args=(image_id, self.delete_image_finish))
            thread.daemon = True
            thread.start()

    @staticmethod
    def delete_image(image_id, finish):
        cli.remove_image(image_id, force=True)
        GObject.idle_add(finish, image_id)

    def delete_image_finish(self, image_id):
        self.deleted_images.remove(image_id)

    def add_delete_container_task(self):
        pass

    def add_download_image_task(self):
        pass

    def delete_container(self):
        pass

    def download_image(self):
        pass

    def delete_container_finish(self):
        pass

    def download_image_finish(self, image_id):
        pass


# pseudo singleton
store = _Store()
