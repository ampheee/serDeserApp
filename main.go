package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
	f "testJsonSerDeser/serializationFuncs"
	s "testJsonSerDeser/structs"
	"time"
)

const (
	w = 450
	h = 500
)

func resizeAndMove(object fyne.CanvasObject, mvPosX, mvPosy, rSizePosX, rSizePosY float32) {
	object.Resize(fyne.NewSize(rSizePosX, rSizePosY))
	object.Move(fyne.NewPos(mvPosX, mvPosy))
}

func doAction(title string, group *widget.RadioGroup,
	window fyne.Window, juniors []s.Junior, techLeads []s.TechLeader,
	entry *widget.Entry, a fyne.App) *widget.Button {
	btn := widget.NewButton(title, func() {
		switch title {
		case "hire techLead":
			if group.Selected == "To memory" {

			} else {
				dialog.ShowError(errors.New("choose 'To memory'"), window)

			}
		case "hire another junior":
			if group.Selected == "To memory" {
				juniorForm := a.NewWindow("hireJuniorForm")
				juniorForm.Resize(fyne.NewSize(w*0.65, h*0.6))
				eEntry, sEntry, tEntry, pEntry, rEntry :=
					widget.NewMultiLineEntry(), widget.NewEntry(),
					widget.NewEntry(), widget.NewEntry(), widget.NewEntry()
				//eEntr := widget.NewFormItem("Email", temp)
				email, salary, team, position, resignation :=
					widget.NewLabel("Email:"), widget.NewLabel("Salary:"),
					widget.NewLabel("Team:"), widget.NewLabel("Position:"),
					widget.NewLabel("Resignation:")
				hire := widget.NewButton("HIRE!", func() {
					var newJun s.Junior
					//var err1, err2, err3 error
					intSalary, _ := strconv.Atoi(salary.Text)
					intTeam, _ := strconv.Atoi(team.Text)
					intResignation, _ := strconv.Atoi(resignation.Text)
					newJun.Email = email.Text
					newJun.Team = intTeam
					newJun.Salary = intSalary
					newJun.Resignation = intResignation
					juniors = append(juniors, newJun)
					fmt.Println(newJun, salary.Text, email.Text)
					//if err1 == nil && err2 == nil && err3 == nil && email.Text != "" && position.Text != "" {
					//
					//} else {
					//	var (
					//		errs     = []error{err1, err2, err3}
					//		warnings []string
					//		finalErr string
					//	)
					//	for i, err := range errs {
					//		if err != nil && i == 0 {
					//			warnings = append(warnings, "salary")
					//		}
					//		if err != nil && i == 1 {
					//			warnings = append(warnings, "number of team")
					//		}
					//		if err != nil && i == 2 {
					//			warnings = append(warnings, "percent of resignation")
					//		}
					//	}
					//	if len(warnings) == 1 {
					//		finalErr = fmt.Sprintf("please, enter valid %s", warnings[0])
					//	} else if len(warnings) == 2 {
					//		finalErr = fmt.Sprintf("please enter valid %s and %s", warnings[0], warnings[1])
					//	} else if len(warnings) == 3 {
					//		finalErr = fmt.Sprintf("please enter valid %s, %s and %s", warnings[0],
					//			warnings[1], warnings[3])
					//	}
					//	dialog.ShowError(errors.New(finalErr), window)
					//}
				})

				email.Move(fyne.NewPos(w*0.01, h*0.03))
				salary.Move(fyne.NewPos(w*0.01, h*0.11))
				team.Move(fyne.NewPos(w*0.01, h*0.19))
				position.Move(fyne.NewPos(w*0.01, h*0.27))
				resignation.Move(fyne.NewPos(w*0.01, h*0.37))

				resizeAndMove(hire, w*0.013, h*0.475, w*0.6, h*0.09)
				resizeAndMove(eEntry, w*0.18, h*0.04, w*0.425, h*0.075)
				resizeAndMove(sEntry, w*0.18, h*0.12, w*0.425, h*0.075)
				resizeAndMove(tEntry, w*0.18, h*0.20, w*0.425, h*0.075)
				resizeAndMove(pEntry, w*0.18, h*0.28, w*0.425, h*0.075)
				resizeAndMove(rEntry, w*0.30, h*0.36, w*0.305, h*0.075)

				juniorForm.SetContent(container.NewWithoutLayout(email, salary, team, position, resignation,
					eEntry, sEntry, tEntry, pEntry, rEntry,
					hire))
				juniorForm.SetFixedSize(true)
				juniorForm.CenterOnScreen()
				juniorForm.Show()
			} else {
				dialog.ShowError(errors.New("choose 'To memory'"), window)
			}
		case "parse [from]":
			if group.Selected == "JSON" {
				err := f.Deserialize("json", entry, window)
				if err != nil {
					dialog.ShowError(err, window)
				}
			} else if group.Selected == "XML" {
				err := f.Deserialize("xml", entry, window)
				if err != nil {
					dialog.ShowError(err, window)
				}
			} else {
				dialog.ShowError(errors.New("choose XML or JSON"), window)
			}
		case "load [to folder]":
			if group.Selected == "JSON" {
				dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
					err = f.Serialize(uri.Path(), "json", juniors, techLeads)
					if err != nil {
						dialog.ShowError(err, window)
					}
				},
					window,
				)
			} else if group.Selected == "XML" {
				dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
					err = f.Serialize(uri.Path(), "xml", juniors, techLeads)
					if err != nil {
						dialog.ShowError(err, window)
					}
				},
					window,
				)
			} else {
				dialog.ShowError(errors.New("choose XML/JSON"), window)
			}
		}
	})
	btn.Resize(fyne.NewSize(w*0.4, h*0.045))
	return btn
}

func totalStuff(w *widget.Label, slice1 []s.Junior, slice2 []s.TechLeader) {
	w.SetText(fmt.Sprintf("Juniors [%d] | techLeads [%d]",
		len(slice1),
		len(slice2)),
	)
}

func main() {
	var (
		mainApp = app.New()
		window  = mainApp.NewWindow("First Lab")

		radio = widget.NewRadioGroup([]string{
			"XML", "JSON", "To memory",
		}, func(s string) {})

		jStaff = []s.Junior{{
			Programmer:  s.Programmer{Email: "check@mail.ru", Salary: 15000, Team: 3, Position: "Junior+"},
			Resignation: 99,
		}}
		tLStaff []s.TechLeader

		entry = widget.NewEntry()

		newTechLeadBtn = doAction("hire techLead", radio, window, jStaff, tLStaff, entry, mainApp)
		newJuniorBtn   = doAction("hire another junior", radio, window, jStaff, tLStaff, entry, mainApp)
		parseBtn       = doAction("parse [from]", radio, window, jStaff, tLStaff, entry, mainApp)
		loadBtn        = doAction("load [to folder]", radio, window, jStaff, tLStaff, entry, mainApp)
		itsMe          = widget.NewLabel("telegram @ampheee\n\t\tmgtu stankin")

		total = widget.NewLabel("")
	)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(w, h))
	window.SetFixedSize(true)
	resizeAndMove(total, w*0.265, h*0.8, w*0.5, h*0.1)
	resizeAndMove(entry, w*0.1, h*0.5, w*0.8, h*0.25)
	resizeAndMove(itsMe, w*0.3, h*0.85, w*0.5, h*0.1)
	radio.Move(fyne.NewPos(w*0.10, h*0.15))

	newTechLeadBtn.Move(fyne.NewPos(w*0.5, h*0.16))
	newJuniorBtn.Move(fyne.NewPos(w*0.5, h*0.22))
	parseBtn.Move(fyne.NewPos(w*0.5, h*0.28))
	loadBtn.Move(fyne.NewPos(w*0.5, h*0.34))

	window.SetContent(container.NewWithoutLayout(
		radio,
		newTechLeadBtn,
		newJuniorBtn,
		parseBtn,
		loadBtn,
		entry,
		total,
		itsMe,
	),
	)

	go func() {
		for range time.Tick(time.Millisecond * 500) {
			totalStuff(total, jStaff, tLStaff)
		}
	}()

	window.ShowAndRun()
}
