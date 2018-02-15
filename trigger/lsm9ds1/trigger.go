package lsm9ds1

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

const DT float64 = 0.02
const AA float64 = 0.97
const A_GAIN float64 = 0.0573
const G_GAIN float64 = 0.070
const RAD_TO_DEG float64 = 57.29578
const M_PI float64 = 3.14159265358979323846

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
	// start the trigger

        var rate_gyr_y float64 = 0.0   // [deg/s]
        var rate_gyr_x float64 = 0.0   // [deg/s]
        var rate_gyr_z  float64 = 0.0   // [deg/s]
        var gyroXangle float64 = 0.0
        var gyroYangle float64 = 0.0
        var gyroZangle float64 = 0.0
        var AccYangle float64 = 0.0
        var AccXangle float64 = 0.0
        var CFangleX float64 = 0.0
        var CFangleY float64 = 0.0

        //var accRaw [3]int
        //var gyrRaw [3]int

	tmp := t.config.GetSetting("interval")
	interval, _ := strconv.Atoi(tmp)

	fmt.Println("interval : %d\n", interval)
        //lsm := lsm9ds1.NewLSM9DS1(embd.NewI2CBus(1))
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

                //Convert Gyro raw to degrees per second
                rate_gyr_x = float64(gyrRaw[0]) * G_GAIN
                rate_gyr_y = float64(gyrRaw[1]) * G_GAIN
                rate_gyr_z = float64(gyrRaw[2]) * G_GAIN

                //Calculate the angles from the gyro
                gyroXangle+=rate_gyr_x*DT
                gyroYangle+=rate_gyr_y*DT
                gyroZangle+=rate_gyr_z*DT

                //Convert Accelerometer values to degrees
                AccXangle = math.Atan2(float64(accRaw[1]),float64(accRaw[2])+M_PI)*RAD_TO_DEG
                AccYangle = math.Atan2(float64(accRaw[2]),float64(accRaw[0])+M_PI)*RAD_TO_DEG

                //If IMU is up the correct way, use these lines
                AccXangle -= float64(180.0)
                if (AccYangle > 90) {
                                AccYangle -= float64(270)
                } else {
                        AccYangle += float64(90)
                }

                //Complementary filter used to combine the accelerometer and gyro values.
                CFangleX=AA*(CFangleX+rate_gyr_x*DT) +(1 - AA) * AccXangle
                CFangleY=AA*(CFangleY+rate_gyr_y*DT) +(1 - AA) * AccYangle

                fmt.Printf("   GyroX  %7.3f, AccXangle  %7.3f, CFangleX  %7.3f, GyroY  %7.3f, AccYangle  %7.3f, CFangleY  %7.3f\n",gyroXangle,AccXangle,CFangleX,gyroYangle,AccYangle,CFangleY)


                // Init handlers
                for _, handlerCfg := range t.config.Handlers {

                        if handlerIsValid(handlerCfg) {
                                log.Debugf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)
                                fmt.Printf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)

//                                t1 := time.Now()

                                data := map[string]interface{}{
                                        "accX":  AccXangle,
                                        "accY":  AccYangle,
                                        "gyrX":  gyroXangle,
                                        "gyrY":  gyroYangle,
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

                time.Sleep(1000 * time.Millisecond)
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

