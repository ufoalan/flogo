package acc_gyr_trigger

import (
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
        "fmt"
        "time"
        "math"
        "os"
	"strconv"
        "github.com/kidoman/embd"
        _ "github.com/kidoman/embd/host/all"
//        "./lsm9ds1"
)

var log = logger.GetLogger("acc-gyr-trigger")

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

	tmp := t.config.GetSetting("sampleRate")
	tmpInt, _ := strconv.Atoi(tmp)
	interval := 1000/tmpInt

	fmt.Println("interval : %d\n", interval)
        lsm := NewLSM9DS1(embd.NewI2CBus(1))
        if (!lsm.DetectIMU()) {
                os.Exit(1)
        }
        lsm.EnableIMU()

        magRaw := make([]int16, 3)
        accRaw := make([]int16, 3)
        gyrRaw := make([]int16, 3)

        for {
                lsm.ReadMAG(magRaw)
                fmt.Printf("Mag : %v\n", magRaw)
                lsm.ReadACC(accRaw)
                fmt.Printf("Acc : %v\n", accRaw)
                lsm.ReadGYR(gyrRaw)
                fmt.Printf("Gyr : %v\n", gyrRaw)

                // Init handlers
                for _, handlerCfg := range t.config.Handlers {

                        if handlerIsValid(handlerCfg) {
                                log.Debugf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)
                                fmt.Printf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)

				t1 := time.Now()
				t2 := t1.Format("2006-01-02 15:04:05.000000")
				out := fmt.Sprintf("%s,%d,%d,%d,%d,%d,%d,%d,%d,%d", t2, accRaw[0], accRaw[1], accRaw[2], gyrRaw[0], gyrRaw[1], gyrRaw[2], magRaw[0], magRaw[1], magRaw[2])
                                data := map[string]interface{}{
                                        "output": out,
					"eventTime" : t2,
                                        "accX":  accRaw[0],
                                        "accY":  accRaw[1],
                                        "accZ":  accRaw[2],
                                        "gyrX":  gyrRaw[0],
                                        "gyrY":  gyrRaw[1],
                                        "gyrZ":  gyrRaw[2],
                                        "magX":  magRaw[0],
                                        "magY":  magRaw[1],
                                        "magZ":  magRaw[2],
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

                time.Sleep(time.Duration(interval) * time.Millisecond)
        }

	return nil
}

// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger
	return nil
}

func handlerIsValid(handler *trigger.HandlerConfig) bool {
        //validate path

        return true
}

