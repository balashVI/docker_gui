import QtQuick 2.5
import Material 0.2

Rectangle {
    id: root
    property var getLogs
    onGetLogsChanged: refresh()

    color: Qt.darker(Theme.primaryColor)
    Flickable {
        id: logs_flickable
        boundsBehavior: Flickable.StopAtBounds
        anchors {
            fill: parent
            margins: Units.dp(5)
        }
        clip: true
        contentHeight: logs_text.height
        contentWidth: width
        Label {
            id: logs_text
            color: "white"
        }
    }
    ModdedScrollbar {
        flickableItem: logs_flickable
        color: "white"
        scrollBarOpacity: 0.5
    }

    Timer {
        id: log_timer
        running: root.visible
        repeat: true
        interval: 3000
        onTriggered: {
            var log = getLogs(false);
            if (log!==""){
                logs_text.text += log;
                logs_flickable.contentY = logs_flickable.contentHeight-logs_flickable.height;
            }
        }
    }

    function refresh(){
        log_timer.restart();
        logs_text.text = getLogs(true);
        logs_flickable.contentY = logs_flickable.contentHeight-logs_flickable.height;
    }
}
