# tibco-raspberrypi3-cameraV2
This trigger provides your flogo application the ability to detect human face and classify objects using rPi camera V2. The trigger first open the RPi camera and save an image once it detects a human face. Then the trigger will pass the image to Tensorflow for object classification using deep learning model.

## Pre-requisite
In order to install this trigger, you have to install OpenCV, GoCV and Tensorflow Go first.
1. Install the RaspberryPi 3 Camera-V2. Please refer to https://projects.raspberrypi.org/en/projects/getting-started-with-picamera about how to connect the camera to RPi.
2. Install GoCV and OpenCV 3.3.1. Please refer to https://gocv.io/getting-started/linux/ for installation.
3. Install Tensorflow Go. Please refer to https://www.tensorflow.org/install/install_go for installation.
4. Add below lines in the end of ~/.bashrc file. The first line is to set the CGO envirnoment variables for OpenCV. The second line is to enable the driver for Raspberry PI 3 Camera V2. (Please noted that step 1 will only enable the camera. In order to allow OpenCV to interact with the camera, bcm2835-v4l2 driver is required)
	source $GOROOT/src/gocv.io/x/gocv/env.sh
	sudo modprobe bcm2835-v4l2	
5. Execute "source ~/.bashrc"
## Installation

```bash
flogo install trigger github.com/ufoalan/trigger/object_detect
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


## Configuration Examples
```json
  "triggers": [
    {
      "id": "my_rest_trigger",
      "ref": "github.com/ufoalan/flogo/trigger/object_detect",
      "settings": {
        "cameraID": "0",
        "classifier_file": "/home/pi/go/src/gocv.io/x/gocv/data/haarcascade_frontalface_default.xml"
      },
```
