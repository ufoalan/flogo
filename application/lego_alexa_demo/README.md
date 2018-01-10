# tibco-lego-ev3-alexa-demo
This demo is to demonstrate how to control Lego Mindstorm EV3 motor using Flogo. In order to make the demo looks more interesting, it leverages the Flogo MQTT Raspberry demo that control LED using Alexa. When the user says voice command to Alexa, the command will be sent to Flogo resides on Lego EV3 and start/stop the motor.

## Pre-requisite
```bash
1. Get the LEGO Mindstorm EV3 boxset. It comes with central controller, 2 large motors, 2 medium motors, 1 color sensor, 1 Infra-red sensor and 1 Infra-red transceiver as remote control
2. Download ev3dev that is Debian Linux-based OS could run on Lego Mindstorm EV3 and copy the image to microSD
3. Put the microSD into EV3 and such it will boot up from microSD
4. Connect EV3 to internet using wifi-dongle or cable to your labtop with internet sharing enabled
```

## Installation

```bash
1. Download lego_alexa_demo.json app from https://github.com/ufoalan/flogo/demo/lego_alexa_demo.json
2. Pull the Flogo docker (docker run -it -p 3303:3303 flogo/flogo-docker:latest eula-accept) and start the Flogo-web
3. Import the lego_alexa_demo.json json file
4. If it prompts you error that mqtt and control_ev3 activity are missing, you have to new a new app and new flow, and then import the activity manually first. Two activities are required and they are
	- github.com/anshulsharmas/flogo-contrib/trigger/mqtt
	- github.com/ufoalan/activity/control_ev3
5. Build the application for ARM/Linux
6. Upload the binary, device.pem.crt, device.pem.key, root-CA.pem (Please refer to Flog Raspberry demo about how to get these certs from AWS) on Lego EV3
7. Issue "chmod +x your_binary" command
8. Start the application by "./your_binary"
```
