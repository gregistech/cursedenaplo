package main

import (
	gc "github.com/gbin/goncurses"
	"github.com/thegergo02/gokreta"
	"strconv"
	"strings"
)

func SwitchToTab(tabName string, windows []*gc.Window, oldWidgetWins []*gc.Window) ([]*gc.Window, error) {
	var widgetWins []*gc.Window
	inst_code, username, password := GetCredetinals()
	if inst_code == "" {
		windows[0].Erase()
		windows[0].Print("Invalid credetinals! Set CK_INST, CK_USERNAME and CK_PASSWORD environment variables! (CK_INST = klikXXXXXXXXXX)")
		windows[0].Refresh()
		return widgetWins, nil
	}
	authDetails, err := gokreta.GetAuthDetailsByCredetinals(inst_code, username, password)
	if err != nil {
		if err.Error() == "invalid_grant" {
			windows[0].Erase()
			windows[0].Print("Invalid credetinals! Set CK_INST, CK_USERNAME and CK_PASSWORD environment variables! (CK_INST = klikXXXXXXXXXX)")
			windows[0].Refresh()
			return widgetWins, err
		}
	}
	student, err := gokreta.GetStudentDetails(inst_code, authDetails.AccessToken)
	if err != nil {
		return widgetWins, err
	}
	DidSwitchHappen := true
	switch tabName {
	case "dashboard":
		lines := []string{
			"Student Informations:",
			"- Name: " + student.Name + " (" + strconv.Itoa(student.Id) + ")",
			"- Place of birth: " + student.PlaceOfBirth,
			"- Form teacher: ",
			"-- Name: " + student.FormTeacher.Name + " (" + strconv.Itoa(student.FormTeacher.Id) + ")",
			"-- Email: " + student.FormTeacher.Email,
			"-- Phone Number: " + student.FormTeacher.PhoneNumber,
		}
		widgetWin, err := CreateWidgetAtPos(GetPositionByName("center", windows, GetWidgetSize(lines)), lines)
		if err != nil {
			return widgetWins, err
		}
		widgetWins = append(widgetWins, widgetWin)
		break
	case "avg":
		lines := []string{}
		for i, e := range student.SubjectAverages {
			lines = append(lines, "- Subject: "+e.Subject+" ("+strconv.Itoa(i)+")")
			lines = append(lines, "-- Subject Category: "+e.SubjectCategoryName)
			lines = append(lines, "-- Your average: "+strconv.FormatFloat(e.Value, 'f', -1, 64))
			lines = append(lines, "-- Class average: "+strconv.FormatFloat(e.ClassValue, 'f', -1, 64))
			lines = append(lines, "-- Difference: "+strconv.FormatFloat(e.Difference, 'f', -1, 64))
			widgetWin, err := CreateWidgetAtPos(GetPositionByName("top-left", windows, GetWidgetSize(lines)), lines)
			if err != nil {
				return widgetWins, err
			}
			widgetWins = append(widgetWins, widgetWin)
		}
		break
	default:
		windows[1].Erase()
		windows[1].Print("This tab doesn't exist!")
		windows[1].Refresh()
		DidSwitchHappen = false
		widgetWins = oldWidgetWins
		break
	}
	if DidSwitchHappen {
		windows[0].Erase()
		windows[0].Print(student.InstituteName + " - " + student.Name + " - " + strings.Title(tabName))
		windows[0].Erase()
		windows[0].Refresh()
		for i := range oldWidgetWins {
			oldWidgetWins[i].Erase()
			oldWidgetWins[i].Delete()
		}
	}
	return widgetWins, err
}
