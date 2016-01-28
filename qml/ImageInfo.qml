import QtQuick 2.0

Item {
    property string imageID: null
    property var image: DockerImages.get(imageID)
}
