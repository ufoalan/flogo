package tcp_server

import (
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/firstrow/tcp_server"
	"fmt"
        "time"
        "math"
        "os"
)

var log = logger.GetLogger("trigger-lsm9ds1")

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
        port := t.config.GetSetting("port")
	url := "localhost:" + port

	// start the trigger
        //server := tcp_server.New("localhost:9999")
        server := tcp_server.New(url)

        server.OnNewClient(func(c *tcp_server.Client) {
                // new client connected
                // lets send some message
                fmt.Println("New client connected")

                // Init handlers
                for _, handlerCfg := range t.config.Handlers {
                        
                        if handlerIsValid(handlerCfg) {
                                log.Debugf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)
                                fmt.Printf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)

                                t1 := time.Now()
                                t2 := t1.Format("2006-01-02 15:04:05.000000")
                                //out := fmt.Sprintf("%s,%f,%f,%f,%f", t2, AccXangle, AccYangle, gyroXangle, gyroYangle)
                                data := map[string]interface{}{
                                        "status": "new client connected",
                                        "eventTime" : t2,
                                }
                                
                                //todo handle error
                                startAttrs, err := t.metadata.OutputsToAttrs(data, false)
                                if err != nil {
                                        log.Errorf("After run error' %s'\n", err)
                                        return err
                                }

                                
                                // run next action
                                act := action.Get(handlerCfg.ActionId)
                                ctx := trigger.NewInitialContext(startAttrs, handlerCfg)
                                //results, err := t.runner.RunAction(ctx, act, nil)
                                _ , err = t.runner.RunAction(ctx, act, nil)
                                
                                if err != nil {
                                        log.Debugf("Object Detection Trigger Error: %s", err.Error())
                                        fmt.Printf("Object Detection Trigger Error: %s", err.Error())
                                        return nil
                                }
                        } else {
                                panic(fmt.Sprintf("Invalid handler: %v", handlerCfg))
                        }
                }

//                c.Send("Hello")
        })
        server.OnNewMessage(func(c *tcp_server.Client, message string) {
                // new message received
                fmt.Println(message)

                // Init handlers
                for _, handlerCfg := range t.config.Handlers {

                        if handlerIsValid(handlerCfg) {
                                log.Debugf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)
                                fmt.Printf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)

                                t1 := time.Now()
                                t2 := t1.Format("2006-01-02 15:04:05.000000")
                                //out := fmt.Sprintf("%s,%f,%f,%f,%f", t2, AccXangle, AccYangle, gyroXangle, gyroYangle)
                                data := map[string]interface{}{
                                        "output": message,
                                        "eventTime" : t2,
                                }

                                //todo handle error
                                startAttrs, err := t.metadata.OutputsToAttrs(data, false)
                                if err != nil {
                                        log.Errorf("After run error' %s'\n", err)
                                        return err
                                }


                                // run next action
                                act := action.Get(handlerCfg.ActionId)
                                ctx := trigger.NewInitialContext(startAttrs, handlerCfg)
                                //results, err := t.runner.RunAction(ctx, act, nil)
                                _ , err = t.runner.RunAction(ctx, act, nil)

                                if err != nil {
                                        log.Debugf("Object Detection Trigger Error: %s", err.Error())
                                        fmt.Printf("Object Detection Trigger Error: %s", err.Error())
                                        return nil
                                }
                        } else {
                                panic(fmt.Sprintf("Invalid handler: %v", handlerCfg))
                        }
                }
        })
        server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
                // connection with client lost
                fmt.Println("client connection closed")
                // Init handlers
                for _, handlerCfg := range t.config.Handlers {

                        if handlerIsValid(handlerCfg) {
                                log.Debugf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)
                                fmt.Printf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)

                                t1 := time.Now()
                                t2 := t1.Format("2006-01-02 15:04:05.000000")
                                //out := fmt.Sprintf("%s,%f,%f,%f,%f", t2, AccXangle, AccYangle, gyroXangle, gyroYangle)
                                data := map[string]interface{}{
                                        "status": "client disconnected",
                                        "eventTime" : t2,
                                }

                                //todo handle error
                                startAttrs, err := t.metadata.OutputsToAttrs(data, false)
                                if err != nil {
                                        log.Errorf("After run error' %s'\n", err)
                                        return err
                                }


                                // run next action
                                act := action.Get(handlerCfg.ActionId)
                                ctx := trigger.NewInitialContext(startAttrs, handlerCfg)
                                //results, err := t.runner.RunAction(ctx, act, nil)
                                _ , err = t.runner.RunAction(ctx, act, nil)

                                if err != nil {
                                        log.Debugf("Object Detection Trigger Error: %s", err.Error())
                                        fmt.Printf("Object Detection Trigger Error: %s", err.Error())
                                        return nil
                                }
                        } else {
                                panic(fmt.Sprintf("Invalid handler: %v", handlerCfg))
                        }
                }
        })

        server.Listen()

	return nil
}

// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger
	return nil
}
