import QtQuick 2.5
import QtQuick.Layouts 1.1
import QtQuick.Controls 1.4 as Controls

import Material 0.2
import Material.ListItems 0.1 as ListItem

import DockerGUI 1.0 as Docker

Item {
    id: containerInfo
    property string containerID
    property var container: Docker.Containers.inspect(containerID)

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
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.top: parent.top
            anchors.topMargin: margings
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
                    iconName: "av/play_arrow"
                    visible: !container.running
                }
                IconButton {
                    iconName: "av/stop"
                    visible: container.running
                }
                IconButton {
                    iconName: "av/replay"
                    visible: container.running
                }
                IconButton {
                    iconName: "action/delete"
                    color: Palette.colors["red"][400]
                }
            }
            Label {
                style: "body1"
                text: qsTr("Image: ") + container.image
                width: parent.width
                elide: Text.ElideRight
            }
            Label {
                style: "body1"
                text: qsTr("Created: ") + container.created
                width: parent.width
                elide: Text.ElideRight
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
