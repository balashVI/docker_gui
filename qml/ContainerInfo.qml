import QtQuick 2.5
import QtQuick.Layouts 1.1
import QtQuick.Controls 1.4 as Controls

import Material 0.2
import Material.ListItems 0.1 as ListItem

Item {
    id: containerInfo
    property string containerID: null
    property var container: DockerContainers.inspect(containerID)
    onContainerChanged: logs.refresh()

    Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: column.height + 2 * column.margings
        contentWidth: width
        interactive: contentHeight > height
        clip: true

        Column {
            id: column
            property int margings: Units.dp(16)
            width: parent.width - 2 * margings
            anchors {
                horizontalCenter: parent.horizontalCenter
                top: parent.top
                topMargin: margings
            }
            spacing: Units.dp(10)

            RowLayout {
                width: parent.width
                spacing: Units.dp(10)
                Label {
                    style: "title"
                    color: Palette.colors['brown'][500]
                    text: container.name
                    Layout.fillWidth: true
                    elide: Text.ElideRight
                }
                IconButton {
                    visible: container.status === "stopped"
                    action: Action {
                        name: qsTr("Start container")
                        iconName: "av/play_arrow"
                    }
                    onClicked: DockerTasks.startContainer(container.id)
                }
                IconButton {
                    visible: container.status === "running"
                    hoverAnimation: true
                    action: Action {
                        name: qsTr("Stop container")
                        iconName: "av/stop"
                    }
                    onClicked: DockerTasks.stopContainer(container.id)
                }
                IconButton {
                    visible: container.status === "running"
                    hoverAnimation: true
                    action: Action {
                        name: qsTr("Restart container")
                        iconName: "av/replay"
                    }
                }
                IconButton {
                    action: Action {
                        name: qsTr("Delete container")
                        iconName: "action/delete"
                    }
                    visible: container.status !== "destroying"
                    color: Palette.colors["red"][400]
                    onClicked: {
                        container.status = "destroying"
                        DockerTasks.deleteContainer(container.id)
                    }
                }
            }
            GridLayout {
                rowSpacing: Units.dp(10)
                columnSpacing: Units.dp(10)
                columns: 2
                Label {
                    text: qsTr("Image: ")
                }
                Label {
                    text: container.image
                }
                Label {
                    text: qsTr("Status: ")
                }
                Label {
                    text: container.status.toUpperCase()
                    style: "body2"
                    color: {
                        if (container.status === "running")
                            return Palette.colors["green"][800]
                        else if (container.status === "destroying")
                            return Palette.colors["red"][800]
                        else
                            return Theme.light.textColor
                    }
                }
                Label {
                    text: qsTr("Created: ")
                }
                Label {
                    text: container.created
                }
            }

            ModdedSectionHeader {
                id: logs_header
                expanded: true
                text: qsTr("Logs")
            }
            Rectangle {
                visible: logs_header.expanded
                width: parent.width
                height: width / 2
                color: Qt.darker(Theme.primaryColor)
                Flickable {
                    id: logs_flickable
                    anchors {
                        fill: parent
                        margins: Units.dp(5)
                    }
                    clip: true
                    contentHeight: logs.height
                    contentWidth: width
                    Label {
                        id: logs
                        color: "white"
                        function refresh(){
                            log_timer.restart();
                            text = container.getLogs(true);
                            logs_flickable.contentY = logs_flickable.contentHeight-logs_flickable.height;
                        }

                        Timer {
                            id: log_timer
                            running: logs_header.expanded
                            repeat: true
                            interval: 5000
                            onTriggered: {
                                var log = container.getLogs(false);
                                if (log!==""){
                                    parent.text += log;
                                    logs_flickable.contentY = logs_flickable.contentHeight-logs_flickable.height;
                                }
                            }
                        }
                    }
                }
                ModdedScrollbar {
                    flickableItem: logs_flickable
                    color: "white"
                    scrollBarOpacity: 0.5
                }
            }

            ModdedSectionHeader {
                id: env_header
                text: qsTr("Environment variables")
            }
            Table {
                visible: env_header.expanded
                width: parent.width
                columns: [{
                        title: "Key",
                        role: "key",
                        width: 1
                    }, {
                        title: "Value",
                        role: "value",
                        width: 2
                    }]
                source: container.env
            }

            ModdedSectionHeader {
                id: ports_header
                text: qsTr("Ports")
            }

            Table {
                visible: ports_header.expanded
                width: parent.width
                columns: [{
                        title: "Container",
                        role: "containerPort",
                        width: 1
                    }, {
                        title: "Host",
                        role: "hostPort",
                        width: 2
                    }]
                source: container.ports
            }

            ModdedSectionHeader {
                id: mounts_header
                text: qsTr("Mounts")
            }
            Table {
                visible: mounts_header.expanded
                width: parent.width
                columns: [{
                        title: "Destination",
                        role: "destination",
                        width: 1
                    }, {
                        title: "Source",
                        role: "source",
                        width: 2
                    }]
                source: container.mounts
            }
        }
    }
    Scrollbar {
        flickableItem: flickable
    }
}
