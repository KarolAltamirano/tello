package drone

import (
	"fmt"
	"github.com/KarolAltamirano/tello/utils/joystick"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"time"
)

// Init drone
func Init(stick *joystick.State) {
	drone := tello.NewDriver("8888")

	work := func() {
		var previousStickState joystick.State

		drone.On(tello.TakeoffEvent, func(data interface{}) {
			fmt.Println("### TakeOff Successful ###")
		})

		drone.On(tello.LandingEvent, func(data interface{}) {
			fmt.Println("### Land Successful ###")
		})

		for {
			// roll
			if stick.Axis.Roll != previousStickState.Axis.Roll {
				if stick.Axis.Roll < 0 {
					fmt.Printf("Left: %v\n", stick.Axis.Roll*-1)
					drone.Left(stick.Axis.Roll * -1)
				} else {
					fmt.Printf("Right: %v\n", stick.Axis.Roll)
					drone.Right(stick.Axis.Roll)
				}
			}

			// pitch
			if stick.Axis.Pitch != previousStickState.Axis.Pitch {
				if stick.Axis.Pitch < 0 {
					fmt.Printf("Forward: %v\n", stick.Axis.Pitch*-1)
					drone.Forward(stick.Axis.Pitch * -1)
				} else {
					fmt.Printf("Backward: %v\n", stick.Axis.Pitch)
					drone.Backward(stick.Axis.Pitch)
				}
			}

			// yaw
			if stick.Axis.Yaw != previousStickState.Axis.Yaw {

				if stick.Axis.Yaw < 0 {
					fmt.Printf("Clockwise: %v\n", stick.Axis.Yaw*-1)
					drone.Clockwise(stick.Axis.Yaw * -1)
				} else {
					fmt.Printf("CounterClockwise: %v\n", stick.Axis.Yaw)
					drone.CounterClockwise(stick.Axis.Yaw)
				}
			}

			// button 12
			if stick.Buttons.Button12 && stick.Buttons.Button12 != previousStickState.Buttons.Button12 {
				fmt.Println("Up: 20")
				drone.Up(20)
			}

			if !stick.Buttons.Button12 && stick.Buttons.Button12 != previousStickState.Buttons.Button12 {
				fmt.Println("Up: 0")
				drone.Up(0)
			}

			// button 11
			if stick.Buttons.Button11 && stick.Buttons.Button11 != previousStickState.Buttons.Button11 {
				fmt.Println("Down: 20")
				drone.Down(20)
			}

			if !stick.Buttons.Button11 && stick.Buttons.Button11 != previousStickState.Buttons.Button11 {
				fmt.Println("Down: 0")
				drone.Down(0)
			}

			// button 10
			if stick.Buttons.Button10 && stick.Buttons.Button10 != previousStickState.Buttons.Button10 {
				fmt.Println("Hover")
				drone.Hover()
			}

			// button 8
			if stick.Buttons.Button08 && stick.Buttons.Button08 != previousStickState.Buttons.Button08 {
				fmt.Println("### TakeOff In Progress ... ###")
				if err := drone.TakeOff(); err != nil {
					fmt.Println("### TakeOff Failed ###")
					fmt.Printf("%#v", err)
				}
			}

			// button 7
			if stick.Buttons.Button07 && stick.Buttons.Button07 != previousStickState.Buttons.Button07 {
				fmt.Println("### Land In Progress ... ###")
				if err := drone.Land(); err != nil {
					fmt.Println("### Land Failed ###")
					fmt.Printf("%#v", err)
				}
			}

			// copy stick status
			previousStickState = *stick

			// timeout
			time.Sleep(1 * time.Millisecond)
		}
	}

	robot := gobot.NewRobot("tello", []gobot.Connection{}, []gobot.Device{drone}, work)

	robot.Start()
}
