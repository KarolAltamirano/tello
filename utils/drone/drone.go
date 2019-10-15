package drone

import (
	"fmt"
	"github.com/KarolAltamirano/tello/utils/joystick"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"time"
)

// Start drone
func Start() {
	drone := tello.NewDriver("8888")

	work := func() {
		var flightData *tello.FlightData

		stick := joystick.Init()
		defer stick.Close()

		drone.On(tello.FlightDataEvent, func(data interface{}) {
			flightData = data.(*tello.FlightData)
		})

		stick.Emitter.On("JoyButtonEvent", func(args ...interface{}) {
			buttonID := args[0].(uint8)
			state := args[1].(bool)

			switch buttonID {
			case joystick.Button12ID:
				if state {
					fmt.Println("Drone TakeOff")
					drone.TakeOff()
				}
			case joystick.Button11ID:
				if state {
					fmt.Println("Drone Land")
					drone.Land()
				}
			case joystick.Button10ID:
				if state {
					fmt.Println("Drone Hover")
					drone.Hover()
				}
			case joystick.Button09ID:
				if state && flightData != nil {
					fmt.Println("------------")
					fmt.Printf("Battery: %v\n", flightData.BatteryPercentage)
					fmt.Printf("Height: %v\n", flightData.Height)
					fmt.Printf("DownVisualState: %v\n", flightData.DownVisualState)
					fmt.Printf("DroneBatteryLeft: %v\n", flightData.DroneBatteryLeft)
					fmt.Printf("DroneFlyTimeLeft: %v\n", flightData.DroneFlyTimeLeft)
					fmt.Println("------------")
				}
			}
		})

		for {
			stick.Run()
			if stick.Ready {
				drone.SetVector(stick.Pitch*-1, stick.Roll, stick.Thrust*-1, stick.Yaw)
			}
			time.Sleep(20 * time.Millisecond)
		}
	}

	robot := gobot.NewRobot("tello", []gobot.Connection{}, []gobot.Device{drone}, work)

	robot.Start()
}
