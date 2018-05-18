# Execute OS command
This activity provides your flogo application the ability to execute an os command


## Installation

```bash
flogo install github.com/ufoalan/flogo/activity/executecommand
```
Link for flogo web:
```
https://github.com/ufoalan/flogo/activity/executecommand
```

## Schema
Inputs and Outputs:

```json
{
  "input":[
    {
      "name": "command",
      "type": "string",
      "required": true,
      "value": "start"
    },
    {
      "name": "background",
      "type": "boolean",
      "value": true
    },
    {
      "name": "arg1",
      "type": "string"
    },
    {
      "name": "arg2",
      "type": "string"
    },
    {
      "name": "arg3",
      "type": "string"
    },
    {
      "name": "arg4",
      "type": "string"
    },
    {
      "name": "arg5",
      "type": "string"
    },
    {
      "name": "arg6",
      "type": "string"
    },
    {
      "name": "arg7",
      "type": "string"
    },
    {
      "name": "arg8",
      "type": "string"
    }
  ],
  "output": [
    {
      "name": "result",
      "type": "string"
    },
    {
      "name": "pid",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting     | Description    |
|:------------|:---------------|
| command     | OS command to be ececuted |
| background  | If set to true, the command will be executed at background without waiting for result |
| arg1        | 1st input argument of command |
| arg2        | 2st input argument of command |
| arg3        | 3st input argument of command |
| arg4        | 4st input argument of command |
| arg5        | 5st input argument of command |
| arg6        | 6st input argument of command |
| arg7        | 7st input argument of command |
| arg8        | 8st input argument of command |


## Configuration Examples
### Simple
Configure a task in flow to echo "hello world" :

```json
{
  "input":[
    {
      "name": "command",
      "type": "string",
      "required": true,
      "value": "echo"
    },
    {
      "name": "background",
      "type": "boolean",
      "value": false 
    },
    {
      "name": "arg1",
      "type": "string"
      "value": "Hello"
    },
    {
      "name": "arg2",
      "type": "string"
      "value": "World"
    },
    {
      "name": "arg3",
      "type": "string"
    },
    {
      "name": "arg4",
      "type": "string"
    },
    {
      "name": "arg5",
      "type": "string"
    },
    {
      "name": "arg6",
      "type": "string"
    },
    {
      "name": "arg7",
      "type": "string"
    },
    {
      "name": "arg8",
      "type": "string"
    }
  ],
  "output": [
    {
      "name": "result",
      "type": "string"
    },
    {
      "name": "pid",
      "type": "string"
    }
  ]
}
```
