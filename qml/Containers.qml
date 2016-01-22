import QtQuick 2.5
import QtGraphicalEffects 1.0
import Material 0.2
import Material.ListItems 0.1 as ListItem

import DockerGUI 1.0 as Docker

Item {
    id: containers

    property int selectedContainer: 0

    Sidebar {
        id: sidebar
        Column {
            width: parent.width
            Repeater {
                model: Docker.Containers.list.len
                delegate: ListItem.Subtitled {
                    id: delegate
                    property var data: Docker.Containers.list.get(modelData)
                    text: data.name
                    subText: data.image

                    action: Icon{
                        anchors.centerIn: parent
                        size: parent.height
                        name: delegate.data.isRunning ? "av/pause_circle_outline" : "av/play_circle_outline"
                        color: delegate.data.isRunning ? Palette.colors["green"][400] : Palette.colors["red"][400]
                    }

                    selected: modelData == containers.selectedContainer
                    onClicked: {
                        containers.selectedContainer = modelData
                    }
                }
            }
        }
    }

    ContainerInfo{
        anchors.top: parent.top
        anchors.right: parent.right
        anchors.bottom: parent.bottom
        anchors.left: sidebar.right
        containerID: Docker.Containers.list.get(parent.selectedContainer).id
    }
}
