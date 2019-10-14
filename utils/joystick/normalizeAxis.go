package joystick

func normalizeAxis(value int16) float32 {
	e := float32(15.0)
	f := float32(value) / 32767.0
	if f > e {
		f = (f - e) / (100.0 - e)
	} else if f < -e {
		f = (f + e) / (100.0 - e)
	} else {
		f = 0.0
	}
	return f
}
