package ui

import (
	"fmt"
	"github.com/gdamore/tcell"
	"os"
)

// InitUI Initialize UI
func InitUI() tcell.Screen {
	screen, err := tcell.NewScreen()

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	if err = screen.Init(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))

	screen.Clear()

	return screen
}
