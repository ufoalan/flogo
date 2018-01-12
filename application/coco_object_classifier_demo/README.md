# tibco-coco-object-classifier-demo
This demo demonstrate the possibility to apply Computer Vision with Tensorflow deep learning in Flogo application. The face_detector trigger will first detect a human face using OpenCV. Once a face is detected, it will save the image to provided file path. The saved file path will be passed to the next action. The next action is COCO object classifier and this classifier is a pre-trained tensorflow model based on COCO dataset. The coco-object-classifier-activity takes the output of face_detector trigger (saved image path) and it applies the classifier to predict the objects in the input image (output image of face_detector trigger). After the prediction, the output image will be saved in the provided folder.

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
1. Download all files from https://github.com/ufoalan/flogo/tree/master/application/coco_object_classifier_demo
2. Create a folder (E.g. /home/pi/coco_object_classifier_demo) and put all files into the folder
3. flogo create -f coco_object_classifier_demo.json
4. Go to the coco_object_classifier_demo folder
5. Execute "flogo build" command to build the flogo application
6. Go to bin folder and change corresponding file path in the flogo.json file
7. Execute the binary file coco_object_classifer_demo in bin directory
```
