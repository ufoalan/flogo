# tibco-face-detection-demo
This is a simple demo to verify the face_detector_trigger and the face_detector_trigger is to demonstrate the capability of integration with OpenCV. With the face_detector_trigger, flogo application could detect a human face using Macbook/Raspberry Pi camera. Once a human face is detected, the image will be stored in provided file path.

## Pre-requisite
```bash
1. Get the OpenCV 3.3.1/3.4 and GoCV installed. Please refer to the face_detector_trigger (https://github.com/ufoalan/flogo/trigger/face_detector) READme regarding the installation.
2. Download haarcascade_frontalface_default.xml classifier file and face_detection_demo.json file from this repository. Please noted that haarcascade_frontalface_default.xml file can be retrieved from GoCV folder as well.
3. If Raspberry Pi 3 is used, connect the camera (RPi camera v2) with RPi 3.
4. Since the GoCV is based on OpenCV which is C/C++ library, CGO is required to build the application. However, cross-compilation is not support by CGO yet. Please do this installation on your Raspberry PI directly. That means you have to install Golang and Flogo-CLI on your Raspberry before installation step below.
```

## Installation

```bash
1. Execute "flogo create -f face_detection_demo.json" command on your terminal.
2. Build the application by "flogo build"
3. Start the application by "bin/your_binary"
4. Once the application is started, you will see the webcam image showing on a window and start to record your face once detected.
```
