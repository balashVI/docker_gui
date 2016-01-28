import QtQuick 2.5
import QtQuick.Layouts 1.1
import Material 0.2

Dialog {
    id: downloadImageDialog
    title: qsTr("Download new image")
    hasActions: false

    RowLayout {
        width: parent.width
        spacing: 5
        TextField {
            id: search
            Layout.fillWidth: true
            placeholderText: qsTr("Search")
            anchors.verticalCenter: parent.verticalCenter
        }
        Button {
            anchors.verticalCenter: parent.verticalCenter
            text: qsTr("Search")
        }
    }
}
