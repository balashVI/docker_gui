import QtQuick 2.5
import Material 0.2
import Material.ListItems 0.1 as ListItem

import DockerGUI 1.0 as Docker

Item {
    id: images

    property int selectedImage: 0

    Sidebar {
        Column {
            width: parent.width
            Repeater {
                model: Docker.Images.list.len
                delegate: ListItem.Subtitled {
                    property var data: Docker.Images.list.get(modelData)
                    text: data.name
                    subText: data.tag

                    selected: modelData == images.selectedImage
                    onClicked: {
                        images.selectedImage = modelData
                    }
                }
            }
        }
    }
}
