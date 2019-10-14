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
	stick := joystick.Init()
	defer stick.Close()

	stick.Emitter.On("JoyButtonEvent", func(args ...interface{}) {
		buttonID := args[0].(uint8)
		state := args[1].(bool)

		switch buttonID {
		case joystick.Button08ID:
			if state {
				fmt.Println("Drone TakeOff")
				drone.TakeOff()
			}
		case joystick.Button07ID:
			if state {
				fmt.Println("Drone Land")
				drone.Land()
			}
		case joystick.Button10ID:
			if state {
				fmt.Println("Drone Hover")
				drone.Hover()
			}
		}
	})

	work := func() {
		for {
			stick.Run()
			if stick.Ready {
				drone.SetVector(stick.Pitch, stick.Roll, stick.Thrust, stick.Yaw)
			}
			time.Sleep(20 * time.Millisecond)
		}
	}

	robot := gobot.NewRobot("tello", []gobot.Connection{}, []gobot.Device{drone}, work)

	robot.Start()
}
