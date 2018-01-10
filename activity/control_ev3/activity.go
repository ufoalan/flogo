package control_ev3

import (
	"errors"
	"fmt"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/ev3go/ev3dev"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-rest")

const (
	method         = "method"
	pinNumber      = "pinNumber"
	value          = "value"
	port           = "port"
	directionState = "direction"
	state          = "state"
	direction      = "Direction"
	setState       = "Set State"
	readState      = "Read State"
	pull           = "Pull"
	start          = "start"
	stop           = "stop"
	auto           = "auto"
	sleep          = "sleep"

	input = "Input"
	//output = "Output"

	high = "High"
	//low = "Low"

	up   = "Up"
	down = "Down"
	//off = "off"

	//ouput

	result = "result"
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
	log.Debug("Running control_ev3 activity.")
	methodInput := context.GetInput(method)

	ivmethod, ok := methodInput.(string)
	if !ok {
		return true, errors.New("Method field not set.")
	}

	//get pinNumber
	sleepDuration, ok := context.GetInput(value).(int)

	if !ok {
		return true, errors.New("Value must exist")
	}

	portInput := context.GetInput(port)
	ivport, ok := portInput.(string)
	if !ok {
		return true, errors.New("Port field not set.")
	}

	log.Debugf("Method '%s', Port '%s' and pin number '%d'", methodInput, portInput, value)

        out, err := ev3dev.TachoMotorFor(ivport, "lego-ev3-l-motor")
        if err != nil {
                log.Debugf("failed to find large motor on %s: %v", ivport, err)
        }
        err = out.SetStopAction("brake").Err()
        if err != nil {
                log.Debugf("failed to set brake stop for large motor on %s: %v", ivport, err)
        }
        maxMedium := out.MaxSpeed()


	switch ivmethod {
	case start:
                out.SetSpeedSetpoint(50 * maxMedium / 100).Command("run-forever")
                checkErrors(out)
	case stop:
                out.Command("stop")
                checkErrors(out)
	case sleep:
                time.Sleep(time.Second * time.Duration(sleepDuration))
	case auto:
                for i := 0; i < 2; i++ {

                        // Run medium motor on outA at speed 50, wait for 0.5 second and then brake.
                        out.SetSpeedSetpoint(50 * maxMedium / 100).Command("run-forever")
                        time.Sleep(time.Second / 2)
                        out.Command("stop")
                        checkErrors(out)

                        // Run medium motor on outA at speed -75, wait for 0.5 second and then brake.
                        out.SetSpeedSetpoint(-75 * maxMedium / 100).Command("run-forever")
                        time.Sleep(time.Second / 2)
                        out.Command("stop")
                        checkErrors(out)
                }
	default:
		log.Errorf("Cannot found method %s ", ivmethod)
		return true, errors.New("Cannot found method %s " + ivmethod)
	}

	context.SetOutput(result, 0)
	return true, nil
}

func checkErrors(devs ...ev3dev.Device) {
        for _, d := range devs {
                err := d.(*ev3dev.TachoMotor).Err()
                if err != nil {
                        drv, dErr := ev3dev.DriverFor(d)
                        if dErr != nil {
                                drv = fmt.Sprintf("(missing driver name: %v)", dErr)
                        }
                        addr, aErr := ev3dev.AddressOf(d)
                        if aErr != nil {
                                drv = fmt.Sprintf("(missing port address: %v)", aErr)
                        }
                        log.Debugf("motor error for %s:%s on port %s: %v", d, drv, addr, err)
                }
        }
}
