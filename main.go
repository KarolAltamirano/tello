package main

import (
	"github.com/KarolAltamirano/tello/utils/drone"
	"github.com/KarolAltamirano/tello/utils/joystick"
)

func main() {
	joystick := joystick.Init()
	drone.Init(joystick)
}
