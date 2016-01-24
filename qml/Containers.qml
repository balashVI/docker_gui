import QtQuick 2.5
import QtGraphicalEffects 1.0
import Material 0.2
import Material.ListItems 0.1 as ListItem

Item {
    id: containers

    property var list: DockerContainers.list
    property int selectedContainer: 0

    Sidebar {
        id: sidebar
        Column {
            width: parent.width
            Repeater {
                model: list.len

                onModelChanged: selectedContainer = 0

                delegate: ListItem.Subtitled {
                    id: delegate
                    property var data: list.get(modelData)
                    text: data.name
                    subText: data.image

                    action: IconButton{
                        anchors.centerIn: parent
                        size: parent.height
                        color: delegate.data.isRunning ? Palette.colors["green"][400] : Palette.colors["red"][400]
                        action: Action{
                            iconName: delegate.data.isRunning ? "av/pause_circle_outline" : "av/play_circle_outline"
                            name: delegate.data.isRunning ? qsTr("Stop container") : qsTr("Start container")
                        }

                        onClicked: {
                            if(delegate.data.isRunning)
                                DockerTasks.stopContainer(delegate.data.id);
                            else
                                DockerTasks.startContainer(delegate.data.id);
                        }
                    }

                    selected: modelData === containers.selectedContainer
                    onClicked: {
                        containers.selectedContainer = modelData
                    }
                }
            }
        }
    }

    ContainerInfo{
        id: container_info
        anchors {
            top: parent.top
            right: parent.right
            bottom: parent.bottom
            left: sidebar.right
        }
        containerID: list.get(selectedContainer).id
    }
}
