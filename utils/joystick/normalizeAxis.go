package joystick

func normalizeAxis(value int16) float32 {
	e := float32(0.15)
	f := float32(value) / 32767
	if f > e {
		f = (f - e) / (1 - e)
	} else if f < -e {
		f = (f + e) / (1 - e)
	} else {
		f = 0
	}
	return f
}
