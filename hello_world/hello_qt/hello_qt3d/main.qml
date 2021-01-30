import QtQuick
import QtQuick.Window
import QtQuick3D.Model

Window {
    width: 640
    height: 480
    visible: true
    title: qsTr("Hello World")
}

Model {
    position: Qt.vector3d(0, -200, 0)
    source: "#Cylinder"
    scale: Qt.vector3d(2, 0.2, 1)
    materials: [ DefaultMaterial {
            diffuseColor: "red"
        }
    ]
}

PerspectiveCamera {
    position: Qt.vector3d(0, 200, 300)
    eulerRotation.x: -30
}

DirectionalLight {
    eulerRotation.x: -30
    eulerRotation.y: -70
}
