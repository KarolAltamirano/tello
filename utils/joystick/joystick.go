package joystick

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	roll = iota
	pitch
	yaw
	thrust
)

const (
	button01 = iota
	button02
	button03
	button04
	button05
	button06
	button07
	button08
	button09
	button10
	button11
	button12
)

func normalizeAxis(value int16) int {
	e := 15.0
	f := (float64(value) / 32767.0) * 100.0
	if f > e {
		f = ((f - e) / (100.0 - e)) * 100.0
	} else if f < -e {
		f = ((f + e) / (100.0 - e)) * 100.0
	} else {
		f = 0
	}
	return int(f)
}

func startEventLoop(state *State) {
	if err := sdl.Init(sdl.INIT_JOYSTICK); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	for {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch event := e.(type) {
			case *sdl.JoyDeviceAddedEvent:
				if stick := sdl.JoystickOpen(event.Which); stick != nil {
					r, _ := stick.AxisInitialState(roll)
					p, _ := stick.AxisInitialState(pitch)
					y, _ := stick.AxisInitialState(yaw)
					t, _ := stick.AxisInitialState(thrust)
					state.Axis.Roll = normalizeAxis(r)
					state.Axis.Pitch = normalizeAxis(p)
					state.Axis.Yaw = normalizeAxis(y)
					state.Axis.Thrust = normalizeAxis(t)
				}
			case *sdl.JoyDeviceRemovedEvent:
				if stick := sdl.JoystickFromInstanceID(event.Which); stick != nil {
					stick.Close()
				}
			case *sdl.JoyAxisEvent:
				switch event.Axis {
				case roll:
					state.Axis.Roll = normalizeAxis(event.Value)
				case pitch:
					state.Axis.Pitch = normalizeAxis(event.Value)
				case yaw:
					state.Axis.Yaw = normalizeAxis(event.Value)
				case thrust:
					state.Axis.Thrust = normalizeAxis(event.Value)
				}
			case *sdl.JoyButtonEvent:
				switch event.Button {
				case button01:
					state.Buttons.Button01 = event.State == 1
				case button02:
					state.Buttons.Button02 = event.State == 1
				case button03:
					state.Buttons.Button03 = event.State == 1
				case button04:
					state.Buttons.Button04 = event.State == 1
				case button05:
					state.Buttons.Button05 = event.State == 1
				case button06:
					state.Buttons.Button06 = event.State == 1
				case button07:
					state.Buttons.Button07 = event.State == 1
				case button08:
					state.Buttons.Button08 = event.State == 1
				case button09:
					state.Buttons.Button09 = event.State == 1
				case button10:
					state.Buttons.Button10 = event.State == 1
				case button11:
					state.Buttons.Button11 = event.State == 1
				case button12:
					state.Buttons.Button12 = event.State == 1
				}
			}

		}

		sdl.Delay(16)
	}
}

// Init joystick
func Init() *State {
	var state State
	go startEventLoop(&state)
	return &state
}
