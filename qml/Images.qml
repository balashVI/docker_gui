import QtQuick 2.5
import Material 0.2
import Material.ListItems 0.1 as ListItem

Item {
    id: images

    property int selectedImage: 0

    Sidebar {
        Column {
            width: parent.width
            Repeater {
                model: DockerImages.list.len
                delegate: ListItem.Subtitled {
                    property var data: DockerImages.list.get(modelData)
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
