package udask_dio
// #cgo LDFLAGS: -L/home/ubuntu/usb-dask_101_i686/lib -lusb_dask
// #cgo CFLAGS: -I/home/ubuntu/usb-dask_101_i686/include
// #include "udask.h"

import "C"
import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

// MyTriggerFactory My Trigger factory
type MyTriggerFactory struct{
	metadata *trigger.Metadata
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MyTriggerFactory{metadata:md}
}

//New Creates a new trigger instance for a given id
func (t *MyTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &MyTrigger{metadata: t.metadata, config:config}
}

// MyTrigger is a stub for your Trigger implementation
type MyTrigger struct {
	metadata *trigger.Metadata
	runner   action.Runner
	config   *trigger.Config
}

// Init implements trigger.Trigger.Init
func (t *MyTrigger) Init(runner action.Runner) {
	t.runner = runner
}

// Metadata implements trigger.Trigger.Metadata
func (t *MyTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *MyTrigger) Start() error {
	// start the trigger
        var cardnum unit = 1;
        var card = C.UD_Register_Card(C.USB_2405, cardnum);
        fmt.Printf("%d\n", card);
	return nil
}

// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger
	return nil
}
