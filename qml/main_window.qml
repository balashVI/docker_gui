import QtQuick 2.5
import Material 0.2

ApplicationWindow {
    title: "Docker GUI"

    theme {
        backgroundColor: "white"
        primaryColor: "#8d6e63"
        accentColor: "red"
        tabHighlightColor: "white"
    }

    initialPage: TabbedPage{
        title: "Docker GUI"

        actions: [
            Action {
                iconName: "action/settings"
                name: "Settings"
                hoverAnimation: true
            }
        ]

        Tab{
            title: "Images"

            Images{
                anchors.fill: parent
            }
        }

        Tab{
            title: "Containers"

            Containers{
                anchors.fill: parent
            }
        }

        Tab{
            title: "Projects"

            Projects{
                anchors.fill: parent
            }
        }
    }
}
