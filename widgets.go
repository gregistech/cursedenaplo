package main

import gc "github.com/gbin/goncurses"

func GetWidgetSize(lines []string) []int {
	var x_size int
	for i := range lines {
		if len(lines[i]) > x_size {
			x_size = len(lines[i])
		}
	}
	return []int{len(lines), x_size}
}

func CreateWidgetAtPos(position []int, lines []string) (*gc.Window, error) {
	widgetSize := GetWidgetSize(lines)
	widgetWin, err := gc.NewWindow(len(lines)+2, widgetSize[1], position[0], position[1])
	widgetWin.Println()
	for i := range lines {
		widgetWin.Print(" ")
		widgetWin.Println(lines[i])
	}
	widgetWin.Box(0, 0)
	widgetWin.Refresh()
	return widgetWin, err
}
