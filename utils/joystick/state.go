package joystick

// State of joystick
type State struct {
	Axis    Axis
	Buttons Buttons
}

// Axis of joystick
type Axis struct {
	Roll   int
	Pitch  int
	Yaw    int
	Thrust int
}

// Buttons of joystick
type Buttons struct {
	Button01 bool
	Button02 bool
	Button03 bool
	Button04 bool
	Button05 bool
	Button06 bool
	Button07 bool
	Button08 bool
	Button09 bool
	Button10 bool
	Button11 bool
	Button12 bool
}
