package main

import gc "github.com/gbin/goncurses"

func GetPositionByName(name string, windows []*gc.Window, widgetSize []int) []int {
	y, x := windows[2].MaxYX()
	positions := map[string][]int{
		"top-left":     []int{1, 0},
		"bottom-left":  []int{(y - 3) - widgetSize[0], 0},
		"top-right":    []int{1, x - widgetSize[1]},
		"bottom-right": []int{(y - 3) - widgetSize[0], x - widgetSize[1]},
		"center-point": []int{y / 2, x / 2},
		"center":       []int{(y / 2) - widgetSize[0], ((x / 2) - widgetSize[1]) + 15},
	}
	return positions[name]
}
