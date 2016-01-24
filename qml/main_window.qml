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
                iconName: "awesome/tasks"
                name: "Background tasks"
                hoverAnimation: true
                onTriggered: bottomSheet.open()
            },
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
    Snackbar{
        id: snackbar
        Timer{
            property string lastEvent: DockerEvents.lastEvent
            property var eventsList: []
            onLastEventChanged: {
                eventsList.unshift(lastEvent);
                running = true
            }
            interval: 2000
            triggeredOnStart: true
            onTriggered: {
                var event = eventsList.pop();
                if (event)
                    snackbar.open(event);
                if(eventsList.length > 0)
                    running = true;
            }
        }
    }
}
