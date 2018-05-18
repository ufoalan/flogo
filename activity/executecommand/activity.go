package executecommand

import (
	"fmt"
	"os/exec"
	"bytes"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-ufoalan-executecommand")

const (
	command		= "command"
	background	= "background"
	arg1		= "arg1"
	arg2		= "arg2"
	arg3		= "arg3"
	arg4		= "arg4"
	arg5		= "arg5"
	arg6		= "arg6"
	arg7		= "arg7"
	arg8		= "arg8"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval

	commandInput := context.GetInput(command)
	ivCommand, ok := commandInput.(string)
	if !ok {
		context.SetOutput("result", "Command Not Set")
		return true, fmt.Errorf("command not set")
	}

	backgroundInput := context.GetInput(background)
	ivBackground, ok := backgroundInput.(bool)

	arg1Input := context.GetInput(arg1)
	ivArg1, ok := arg1Input.(string)

	arg2Input := context.GetInput(arg2)
	ivArg2, ok := arg2Input.(string)

	arg3Input := context.GetInput(arg3)
	ivArg3, ok := arg3Input.(string)

	arg4Input := context.GetInput(arg4)
	ivArg4, ok := arg4Input.(string)

	arg5Input := context.GetInput(arg5)
	ivArg5, ok := arg5Input.(string)

	arg6Input := context.GetInput(arg6)
	ivArg6, ok := arg6Input.(string)

	arg7Input := context.GetInput(arg7)
	ivArg7, ok := arg7Input.(string)

	arg8Input := context.GetInput(arg8)
	ivArg8, ok := arg8Input.(string)


	cmd := exec.Command(ivCommand, ivArg1, ivArg2, ivArg3, ivArg4, ivArg5, ivArg6, ivArg7, ivArg8)
	fmt.Printf("%s, %s, %s, %s, %s, %s, %s, %s, %s\n", ivCommand, ivArg1, ivArg2, ivArg3, ivArg4, ivArg5, ivArg6, ivArg7, ivArg8)

	var outb, errb bytes.Buffer
	var err1 error
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	if (ivBackground) {
		err1 = cmd.Start()
	} else {
		err1 = cmd.Run()
	}
	var msg = ""
	if err1 != nil {
        	log.Error(err1)
		msg = errb.String()
	} else {
		msg = outb.String()
	}
	outStr, errStr := string(outb.Bytes()), string(errb.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	fmt.Printf("pid=%d\n", cmd.Process.Pid)
	log.Info("pid=%d", cmd.Process.Pid)
	context.SetOutput("result", msg)
	context.SetOutput("pid", cmd.Process.Pid)

	return true, nil
}
