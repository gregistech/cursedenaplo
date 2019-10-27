package main

import gc "github.com/gbin/goncurses"

func InitAllWindows() ([]*gc.Window, error) {
	scr, err := initMainWindow()
	infoWin, err := initWindow("info", scr)
	commandWin, err := initWindow("command", scr)
	return []*gc.Window{infoWin, commandWin, scr}, err
}

func initMainWindow() (*gc.Window, error) {
	scr, err := gc.Init()
	defer gc.End()
	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)
	return scr, err
}

func initWindow(name string, scr *gc.Window) (*gc.Window, error) {
	y_size, x_size := scr.MaxYX()
	switch name {
	case "info":
		infoWin, err := gc.NewWindow(1, x_size-1, 0, 0)
		infoWin.SetBackground(gc.A_REVERSE)
		infoWin.Refresh()
		return infoWin, err
	case "command":
		return gc.NewWindow(1, x_size-1, y_size-1, 0)
	}
	return scr, nil
}

func ToggleUserInput(isUserInput bool) (bool, string) {
	if isUserInput {
		gc.Cursor(0)
	} else {
		gc.Cursor(1)
	}
	return !isUserInput, ""
}
