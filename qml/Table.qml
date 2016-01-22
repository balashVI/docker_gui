import QtQuick 2.5
import Material 0.2

Column {
    id: table
    property var columns: []
    property var source
    // інколи не оновлює len якщо задати model:source.len, тому задається вручну
    onSourceChanged: rows.model = source.len

    property int fullWidth: {
        var res = 0;
        for(var i =0;i<columns.length;i++)
            res += columns[i].width;
        if (res < 1)
            return 1;
        else
            return res;
    }

    Row {
        Repeater {
            model: columns
            delegate: Item {
                height: Units.dp(38)
                width: modelData.width/fullWidth * table.width

                Label {
                    style: "body2"
                    color: Theme.light.subTextColor
                    elide: Text.ElideRight
                    text: modelData.title.toUpperCase()

                    anchors {
                        verticalCenter: parent.verticalCenter
                        left: parent.left
                        right: parent.right
                        margins: Units.dp(16)
                    }

                }
            }
        }
    }

    ThinDivider{}

    Item{width: parent.width; height: Units.dp(7)}

    Repeater{
        id: rows

        delegate: Row{
            id: row
            property int number: modelData

            Repeater{
                model: columns

                delegate: Item {
                    height: label.height + Units.dp(10)
                    width: modelData.width/fullWidth * table.width
                    anchors.verticalCenter: parent.verticalCenter

                    Label {
                        id: label
                        style: "body1"
                        color: Theme.light.subTextColor
                        wrapMode: Text.WrapAnywhere
                        text: source.get(row.number)[modelData.role]

                        anchors {
                            left: parent.left
                            right: parent.right
                            margins: Units.dp(16)
                        }

                    }
                }
            }
        }
    }
}
