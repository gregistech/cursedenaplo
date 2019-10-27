package main

import gc "github.com/gbin/goncurses"

func startInputLoop(windows []*gc.Window) {
	var widgetWins []*gc.Window
	in := make(chan gc.Char)
	ready := make(chan bool)
	go func(w *gc.Window, ch chan<- gc.Char) {
		for {
			<-ready
			ch <- gc.Char(w.GetChar())
		}
	}(windows[1], in)
	isUserInput := false
	var currentCommand string
	for {
		var c gc.Char
		select {
		case c = <-in:
			if isUserInput {
				if c == 10 || c == gc.KEY_ENTER {
					windows[1].Erase()
					widgetWins = executeCommand(currentCommand, windows, widgetWins)
					isUserInput, currentCommand = ToggleUserInput(isUserInput)
				} else if c == gc.KEY_BACKSPACE || c == 127 || c == 8 {
					if len(currentCommand) > 0 {
						currentCommand = currentCommand[:len(currentCommand)-1]
						curs_y, curs_x := windows[1].CursorYX()
						windows[1].MoveDelChar(curs_y, curs_x-1)
						windows[1].Refresh()
					}
				} else {
					windows[1].Print(string(c))
					windows[1].Refresh()
					currentCommand += string(c)
				}
			} else {
				if c == gc.Char(':') {
					windows[1].Erase()
					windows[1].Print(":")
					windows[1].Refresh()
					isUserInput, currentCommand = ToggleUserInput(isUserInput)
				}
			}
		case ready <- true:
		}
	}
}
