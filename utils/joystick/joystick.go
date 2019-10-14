package joystick

import (
	"fmt"
	"github.com/KarolAltamirano/tello/utils/emitter"
	"github.com/veandco/go-sdl2/sdl"
)

// Joystick struct
type Joystick struct {
	Emitter emitter.Emitter
	Ready   bool
	Roll    float32
	Pitch   float32
	Yaw     float32
	Thrust  float32
}

// Init joystick
func Init() *Joystick {
	if err := sdl.Init(sdl.INIT_JOYSTICK); err != nil {
		panic(err)
	}
	joystick := &Joystick{
		Emitter: emitter.NewEmitter(),
		Ready:   false,
		Roll:    0,
		Pitch:   0,
		Yaw:     0,
		Thrust:  0,
	}
	return joystick
}

// Close joystick
func (j Joystick) Close() {
	sdl.Quit()
}

// Run joystick
func (j *Joystick) Run() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch event := e.(type) {
		case *sdl.JoyDeviceAddedEvent:
			if !j.Ready {
				if stick := sdl.JoystickOpen(event.Which); stick != nil {
					r, _ := stick.AxisInitialState(RollID)
					p, _ := stick.AxisInitialState(PitchID)
					y, _ := stick.AxisInitialState(YawID)
					t, _ := stick.AxisInitialState(ThrustID)
					j.Roll = normalizeAxis(r)
					j.Pitch = normalizeAxis(p)
					j.Yaw = normalizeAxis(y)
					j.Thrust = normalizeAxis(t)
					j.Ready = true
					fmt.Println("Joystick Connected")
					fmt.Printf("Joystick Axis: %+.4f | %+.4f | %+.4f | %+.4f\n", j.Roll, j.Pitch, j.Yaw, j.Thrust)
				}
			}
		case *sdl.JoyDeviceRemovedEvent:
			if stick := sdl.JoystickFromInstanceID(event.Which); stick != nil {
				stick.Close()
				j.Ready = false
				fmt.Println("Joystick Removed")
			}
		case *sdl.JoyAxisEvent:
			switch event.Axis {
			case RollID:
				j.Roll = normalizeAxis(event.Value)
			case PitchID:
				j.Pitch = normalizeAxis(event.Value)
			case YawID:
				j.Yaw = normalizeAxis(event.Value)
			case ThrustID:
				j.Thrust = normalizeAxis(event.Value)
			}
			fmt.Printf("Joystick Axis: %+.4f | %+.4f | %+.4f | %+.4f\n", j.Roll, j.Pitch, j.Yaw, j.Thrust)
		case *sdl.JoyButtonEvent:
			j.Emitter.Emit("JoyButtonEvent", event.Button, event.State == 1)
			fmt.Printf("Joystick Button: %2d | %v\n", event.Button, event.State == 1)
		}
	}
}
