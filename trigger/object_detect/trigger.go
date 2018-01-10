package object_detect

// #cgo LDFLAGS: -lstdc++
import "C"
import (
        "fmt"
        "image"
        "image/color"
	"strconv"
	"time"

        "gocv.io/x/gocv"
	"gopkg.in/gomail.v2"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("trigger-object-detect")

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
	cameraID := t.config.GetSetting("cameraID")
	deviceID, _ := strconv.Atoi(cameraID)
	xmlFile := t.config.GetSetting("classifier_file")
	// start the trigger
        // open webcam
        webcam, err := gocv.VideoCaptureDevice(int(deviceID))
        if err != nil {
                fmt.Printf("error opening video capture device: %v\n", deviceID)
                return nil
        }
        defer webcam.Close()

        // open display window
        window := gocv.NewWindow("Face Detect")
        defer window.Close()

        // prepare image matrix
        img := gocv.NewMat()
        defer img.Close()

        // color for the rect when faces detected
        blue := color.RGBA{0, 0, 255, 0}

        // load classifier to recognize faces
        classifier := gocv.NewCascadeClassifier()
        defer classifier.Close()

        classifier.Load(xmlFile)

        fmt.Printf("start reading camera device: %v\n", deviceID)
        for {
                if ok := webcam.Read(img); !ok {
                        fmt.Printf("cannot read device %d\n", deviceID)
                        return nil
                }
                if img.Empty() {
                        continue
                }

                // detect faces
                rects := classifier.DetectMultiScale(img)
                fmt.Printf("found %d faces\n", len(rects))

                // draw a rectangle around each face on the original image,
                // along with text identifing as "Human"
                for _, r := range rects {
                        gocv.Rectangle(img, r, blue, 3)

                        size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
                        pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
                        gocv.PutText(img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)

        		// Init handlers
        		for _, handlerCfg := range t.config.Handlers {

                		if handlerIsValid(handlerCfg) {
                        		log.Debugf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)
                        		fmt.Printf("Object Detection Trigger: Registering handler for Action Id: [%s]", handlerCfg.ActionId)

					t1 := time.Now()
					s := t1.Format("20060102150405")
					saveFile := "/home/pi/Downloads/tmp/" + s + ".jpg"
					gocv.IMWrite(saveFile, img)

					dir := "/home/pi/tf_test/ssd_mobilenet_v1_coco_11_06_2017"
					jpg := saveFile
					outjpg := "/home/pi/Downloads/tmp/" + s + "_out.jpg"
					labelsFile := "/home/pi/tf_test/labels.txt"
					classify(dir, jpg, outjpg, labelsFile)
					//sendMail(saveFile)
                			data := map[string]interface{}{
                        			"output":  "human detected",
                        			"image":   saveFile,
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
                }

                // show the image in the window, and wait 1 millisecond
                window.IMShow(img)
                window.WaitKey(1)
        }

	return nil
}

// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////
// Utils

func handlerIsValid(handler *trigger.HandlerConfig) bool {
//        if handler.Settings == nil {
//                return false
//        }

//        if handler.Settings["cameraID"] == "" {
//                return false
//        }

//        if handler.Settings["classifier_file"] == "" {
//                return false
//        }

        //validate path

        return true
}

func stringInList(str string, list []string) bool {
        for _, value := range list {
                if value == str {
                        return true
                }
        }
        return false
}

func sendMail(saveFile string) {
       	m := gomail.NewMessage()
       	m.SetHeader("From", "flogo@example.com")
       	//m.SetHeader("To", "flogo@example.com", "test@example.com", "test1@example.com")
       	m.SetHeader("To", "test@example.com")
       	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
       	m.SetHeader("Subject", "Hello!")
       	m.SetBody("text/html", "Hello, It's <b>flogo</b> test for face detecting</b>!")
       	m.Attach(saveFile)

       	d := gomail.NewDialer("smtp.gmail.com", 465, "user@example.com", "passwd")

       	// Send the email to Bob, Cora and Dan.
       	if err := d.DialAndSend(m); err != nil {
        	panic(err)
       	}
}
