# tibco-lego-ev3
This activity provides your flogo application the ability to control LEGO mindstorm EV3 motor

## Installation

```bash
flogo add activity github.com/ufoalan/activity/control_ev3
```

## Schema
Inputs and Outputs:

```json
  "inputs":[
    {
      "name": "method",
      "type": "string",
      "required": true,
      "allowed" : ["start", "stop", "auto", "sleep"]
    },
    {
      "name": "value",
      "type": "integer",
      "required": true
    },
    {
      "name": "port",
      "type": "string",
      "allowed" : ["outA", "outB", "outC", "outD"]
    },
    {
      "name": "state",
      "type": "string",
      "allowed" : ["High", "Low"]
    },

    {
      "name": "Pull",
      "type": "string",
      "allowed" : ["Up", "Down", "Off"]
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "integer"
    }
  ]
```
## Settings
| Setting     | Description    |
|:------------|:---------------|
| method      | The method to take action for LEGO mindstorm EV3 large motor|         
| value       | The duration for "sleep" method |
| port        | The output port on EV3 you want to control |
| state       | Reserved |
| Pull        | Reserved |


## Configuration Examples
### Start motor
Start large motor on port outA
```json
  "attributes": [
          {
            "name": "method",
            "value": "start",
            "type": "string"
          },
          {
            "name": "port",
            "value": "outA",
            "type": "string"
          }
        ]
```
### Stop motor
Stop large motor on port outA
```json
  "attributes": [
          {
            "name": "method",
            "value": "stop",
            "type": "string"
          },
          {
            "name": "port",
            "value": "outA",
            "type": "string"
          }
        ]
```
### Sleep function
Sleep function
```json
  "attributes": [
          {
            "name": "method",
            "value": "sleep",
            "type": "string"
          },
          {
            "name": "value",
            "value": "2",
            "type": "integer"
          }
        ]
```
### Start large motor with auto mode
Start large motor on port outA with auto mode (turn -> sleep 2 sec -> turn again -> stop)
```json
  "attributes": [
          {
            "name": "method",
            "value": "auto",
            "type": "string"
          },
          {
            "name": "port",
            "value": "outA",
            "type": "string"
          }
        ]
```
