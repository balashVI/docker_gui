import QtQuick 2.5

import Material.ListItems 0.1 as ListItem

Rectangle {
    property color backgroundColor: Qt.rgba(0.65, 0.6, 0.59, 0.1)

    property alias expanded: header.expanded
    property alias text: header.text

    width: parent.width
    height: header.height

    color: header.expanded ? backgroundColor : "transparent"

    ListItem.SectionHeader {
        id: header
    }
}
