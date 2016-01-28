import QtQuick 2.5
import Material 0.2
import Material.ListItems 0.1 as ListItem

Item {
    id: root

    property var list: DockerImages.list
    property int selectedImage: 0

    Sidebar {
        id: sidebar
        height: parent.height
        Column {
            width: parent.width
            Repeater {
                model: list.len
                delegate: ListItem.Subtitled {
                    property var data: list.get(modelData)
                    text: data.name
                    subText: data.tag

                    selected: modelData == root.selectedImage
                    onClicked: {
                        root.selectedImage = modelData
                    }
                }
            }
        }

    }

    ActionButton {
        anchors {
            right: sidebar.right
            bottom: parent.bottom
            margins: Units.dp(20)
        }
        isMiniSize: true

        action: Action {
            name: qsTr("Add new")
            onTriggered: downloadDialog.show()
            iconName: "content/add"
        }
    }

    DownloadImageDialog{
        id: downloadDialog
        width: parent.width*0.6
    }

    ImageInfo{
        anchors {
            top: parent.top
            right: parent.right
            bottom: parent.bottom
            left: sidebar.right
        }
        imageID: list.get(selectedImage).id
    }
}
