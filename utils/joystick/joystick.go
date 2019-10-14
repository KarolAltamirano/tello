package joystick

import (
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
	joystick := &Joystick{}
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
			}
		case *sdl.JoyDeviceRemovedEvent:
			if stick := sdl.JoystickFromInstanceID(event.Which); stick != nil {
				stick.Close()
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
		case *sdl.JoyButtonEvent:
			j.Emitter.Emit("JoyButtonEvent", event.Button, event.State == 1)
		}
	}
}
