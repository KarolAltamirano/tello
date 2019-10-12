package drone

import (
	"fmt"
	"github.com/KarolAltamirano/tello/utils/joystick"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"time"
)

// Init drone
func Init(joystick *joystick.State) {
	drone := tello.NewDriver("8888")

	work := func() {
		drone.On(tello.TakeoffEvent, func(data interface{}) {
			fmt.Println("### TakeOff Successful ###")
		})

		drone.On(tello.LandingEvent, func(data interface{}) {
			fmt.Println("### Land Successful ###")
		})

		for {
			// roll
			if joystick.Axis.Roll < 0 {
				drone.Left(joystick.Axis.Roll * -1)
			} else {
				drone.Right(joystick.Axis.Roll)
			}

			// pitch
			if joystick.Axis.Pitch < 0 {
				drone.Forward(joystick.Axis.Pitch * -1)
			} else {
				drone.Backward(joystick.Axis.Pitch)
			}

			// yaw
			if joystick.Axis.Yaw < 0 {
				drone.Clockwise(joystick.Axis.Yaw * -1)
			} else {
				drone.CounterClockwise(joystick.Axis.Yaw)
			}

			if joystick.Buttons.Button12 {
				drone.Up(20)
			}

			if joystick.Buttons.Button11 {
				drone.Down(20)
			}

			if joystick.Buttons.Button10 {
				drone.Hover()
			}

			if joystick.Buttons.Button08 {
				fmt.Println("### TakeOff In Progress ... ###")
				if err := drone.TakeOff(); err != nil {
					fmt.Println("### TakeOff Failed ###")
					fmt.Printf("%#v", err)
				}
			}

			if joystick.Buttons.Button07 {
				fmt.Println("### Land In Progress ... ###")
				if err := drone.Land(); err != nil {
					fmt.Println("### Land Failed ###")
					fmt.Printf("%#v", err)
				}
			}

			time.Sleep(1 * time.Millisecond)
		}
	}

	robot := gobot.NewRobot("tello", []gobot.Connection{}, []gobot.Device{drone}, work)

	robot.Start()
}
