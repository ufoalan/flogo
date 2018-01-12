package coco_object_classifier

import (
	"errors"
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-coco-object-classifier")

const (
	ivModelFilePath   = "model_file_path"
	ivInputImageFile  = "input_image_file"
	ivOutputImagePath = "output_image_path"
	ivLabelsFile      = "labels_file"

	ivOutput= "output"
)

type GPIOActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new GPIOActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &GPIOActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *GPIOActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *GPIOActivity) Eval(context activity.Context) (done bool, err error) {
	//getmethod
	log.Debug("Running COCO object classifier activity.")
	fmt.Println("Running COCO object classifier activity.")
	modelFilePathInput := context.GetInput(ivModelFilePath)
	inputImageFileInput := context.GetInput(ivInputImageFile)
	outputImagePathInput := context.GetInput(ivOutputImagePath)
	labelsFileInput := context.GetInput(ivLabelsFile)

	model_file_path, ok := modelFilePathInput.(string)
	if !ok {
		return true, errors.New("model_file_path field not set.")
	}

	input_image_file, ok := inputImageFileInput.(string)
	if !ok {
		return true, errors.New("input_image_file field not set.")
	}

	output_image_path, ok := outputImagePathInput.(string)
	if !ok {
		return true, errors.New("output_image_path field not set.")
	}

	// Get last index, searching from right to left.
	i := strings.LastIndex(input_image_file, "/")
	inputfile := input_image_file[i:]
	outputFile := strings.Replace(inputfile, ".jpg", "_out.jpg", -1)
	outputpath := output_image_path + outputFile
	fmt.Println(outputpath)

	labels_file, ok := labelsFileInput.(string)
	if !ok {
		return true, errors.New("labels_file field not set.")
	}

	log.Debugf("Model file path : '%s', input image file : '%s', output image path : '%s', outputpath : '%s', labels file : '%s'", model_file_path, input_image_file, output_image_path, outputpath, labels_file)
	fmt.Println("Model file path : '%s', input image file : '%s', output image path : '%s', outputpath : '%s', labels file : '%s'", model_file_path, input_image_file, output_image_path, outputpath, labels_file)
	classify(model_file_path, input_image_file, outputpath, labels_file)

	context.SetOutput(ivOutput, outputpath)
	return true, nil
}
