package main

import (
	gc "github.com/gbin/goncurses"
	"os"
	"strings"
)

func cleanupCurses() {
	gc.Echo(true)
	gc.CBreak(false)
	gc.Cursor(1)
}

func exitCurses() {
	cleanupCurses()
	os.Exit(0)
	return
}

func executeCommand(currentCommand string, windows []*gc.Window, oldWidgetWins []*gc.Window) []*gc.Window {
	var widgetWins []*gc.Window
	var err error
	command := strings.Split(currentCommand, " ")
	switch command[0] {
	case "q":
		exitCurses()
	case "tab":
		widgetWins, err = SwitchToTab(command[1], windows, oldWidgetWins)
		if err != nil {
			return widgetWins
		}
	default:
		windows[1].Erase()
		windows[1].Print("This command doesn't exist!")
	}
	return widgetWins
}

func main() {
	windows, err := InitAllWindows()
	if err != nil {
		panic(err)
	}
	startInputLoop(windows)
}
