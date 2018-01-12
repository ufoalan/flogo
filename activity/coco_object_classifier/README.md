# tibco-coco-object-classifier-activity
This activity provides your flogo application the ability to leverage tensorflow pre-trained deep learning model based on COCO dataset.

## Installation

```bash
flogo install activity github.com/ufoalan/flogo/activity/coco_object_classifier
```

## Schema
Inputs and Outputs:

```json
  "inputs":[
    {
      "name": "model_file_path",
      "type": "string",
      "required": true
    },
    {
      "name": "input_image_file",
      "type": "string",
      "required": true
    },
    {
      "name": "output_image_path",
      "type": "string",
      "required": true
    },
    {
      "name": "labels_file",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "output",
      "type": "string"
    }
  ]
```
## Input
| Input             | Description                                                     |
|:------------------|:----------------------------------------------------------------|
| model_file_path   | This is the model file full path of Tensorflow model file (.pb) |
| input_image_file  | This is the input image file full path                          |
| output_image_path | This is full path of output image                               |
| labels_file       | This is full path of label file                                 |


## Configuration Examples
```json
                "attributes": [
                  {
                    "name": "model_file_path",
                    "value": "/home/pi/tf_test/ssd_mobilenet_v1_coco_11_06_2017/frozen_inference_graph.pb",
                    "type": "string"
                  },
                  {
                    "name": "input_image_file",
                    "value": "/home/pi/tf_test/test.jpg",
                    "type": "string"
                  },
                  {
                    "name": "output_image_path",
                    "value": "/home/pi/tf_test/output",
                    "type": "string"
                  },
                  {
                    "name": "labels_file",
                    "value": "/home/pi/tf_test/labels.txt",
                    "type": "string"
                  }
                ]
```
