package main

import (
	"fmt"
	gc "github.com/gbin/goncurses"
	"github.com/thegergo02/gokreta"
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
			fmt.Sprintf("- Name: %s (%d)", student.Name, student.Id),
			fmt.Sprintf("- Place of birth: %s", student.PlaceOfBirth),
			"- Form teacher: ",
			fmt.Sprintf("-- Name: %s (%d)", student.FormTeacher.Name, student.FormTeacher.Id),
			fmt.Sprintf("-- Email: %s", student.FormTeacher.Email),
			fmt.Sprintf("-- Phone Number: %s", student.FormTeacher.PhoneNumber),
		}
		widgetWin, err := CreateWidgetAtPos(GetPositionByName("center", windows, GetWidgetSize(lines)), lines)
		if err != nil {
			return widgetWins, err
		}
		widgetWins = append(widgetWins, widgetWin)
		break
	case "avg":
		for i, e := range student.SubjectAverages {
			lines := []string{}
			lines = append(lines, fmt.Sprintf("- Subject: %s (%d)", e.Subject, i))
			lines = append(lines, fmt.Sprintf("-- Subject Category: %s", e.SubjectCategoryName))
			lines = append(lines, fmt.Sprintf("-- Your average: %.2f", e.Value))
			lines = append(lines, fmt.Sprintf("-- Class average: %.2f", e.ClassValue))
			lines = append(lines, fmt.Sprintf("-- Difference: %.2f", e.Difference))
			widgetWin, err := CreateWidgetAtPos(GetPositionByName("top-left", windows, GetWidgetSize(lines)), lines)
			if err != nil {
				return widgetWins, err
			}
			widgetWins = append(widgetWins, widgetWin)
		}
		break
	case "evals":
		var y_lines int
		for i := range student.Evaluations {
			lines := []string{}
			e := student.Evaluations[i]
			lines = append(lines, fmt.Sprintf("- Type: %s", e.TypeName))
			lines = append(lines, fmt.Sprintf("- Subject: %s", e.Subject))
			lines = append(lines, fmt.Sprintf("- Mode: %s", e.Mode))
			lines = append(lines, fmt.Sprintf("- Weight: %s", e.Weight))
			lines = append(lines, fmt.Sprintf("- Subject Category: %s", e.SubjectCategory))
			lines = append(lines, fmt.Sprintf("- Form: %s", e.FormName))
			lines = append(lines, fmt.Sprintf("- Theme: %s", e.Theme))
			lines = append(lines, fmt.Sprintf("- Affects Average?: %t", e.DoesCountIntoAvg))
			lines = append(lines, fmt.Sprintf("- Value: %s (%d)", e.Value, e.NumberValue))
			lines = append(lines, fmt.Sprintf("- Teacher: %s", e.Teacher))
			lines = append(lines, fmt.Sprintf("- Date: %s", e.Date))
			lines = append(lines, "- Nature:")
			lines = append(lines, fmt.Sprintf("-- Name: %s", e.Nature.Name))
			lines = append(lines, fmt.Sprintf("-- Description: %s", e.Nature.Description))
			pos := GetPositionByName("top-left", windows, GetWidgetSize(lines))
			pos[0] += y_lines
			y_lines += len(lines) + 2
			widgetWin, err := CreateWidgetAtPos(pos, lines)
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
			oldWidgetWins[i].Refresh()
			oldWidgetWins[i].Delete()
		}
	}
	return widgetWins, err
}
