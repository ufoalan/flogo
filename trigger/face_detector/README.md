# tibco-raspberrypi3-cameraV2
This trigger provides face detection capability to your flogo application. When the camera detect a human face, it will store the image in the given file path and the file path will be passed to next action as output parameter.

## Pre-requisite
In order to install this trigger, you have to install OpenCV and GoCV.
1. Install the RaspberryPi 3 Camera-V2. Please refer to https://projects.raspberrypi.org/en/projects/getting-started-with-picamera about how to connect the camera to RPi.
2. Install GoCV and OpenCV 3.3.1. Please refer to https://gocv.io/getting-started/linux/ for installation.
3. Add below lines in the end of ~/.bashrc file. The first line is to set the CGO envirnoment variables for OpenCV. The second line is to enable the driver for Raspberry PI 3 Camera V2. (Please noted that step 1 will only enable the camera. In order to allow OpenCV to interact with the camera, bcm2835-v4l2 driver is required)
	source $GOROOT/src/gocv.io/x/gocv/env.sh
	sudo modprobe bcm2835-v4l2	
4. Execute "source ~/.bashrc"
## Installation

```bash
flogo install trigger github.com/ufoalan/trigger/face_detector
```

## Schema
Inputs and Outputs:

```json
  "settings":[
    {
      "name": "cameraID",
      "type": "integer",
      "required": "true"
    },
    {
      "name": "classifier_file",
      "type": "string",
      "required": "true"
    },
    {
      "name": "save_image_path",
      "type": "string",
      "required": "true"
    }
  ],
  "outputs": [
    {
      "name": "output",
      "type": "string"
    },
    {
      "name": "image",
      "type": "string"
    }
  ],
```
## Settings
| Setting         | Description    |
|:----------------|:---------------|
| cameraID        | Default is 0 |
| classifier_file | Classification model file |
| save_image_path | The recorded image will be stored in this path |


## Configuration Examples
```json
  "triggers": [
    {
      "id": "face_detector_trigger",
      "ref": "github.com/ufoalan/flogo/trigger/face_detector",
      "settings": {
        "cameraID": "0",
        "classifier_file": "/home/pi/go/src/gocv.io/x/gocv/data/haarcascade_frontalface_default.xml"
        "save_image_path": "/home/pi/flogo_workspace/images"
      },
```
