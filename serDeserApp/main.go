package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	f "testJsonSerDeser/serializationFuncs"
	s "testJsonSerDeser/structs"
	"time"
)

const (
	w = 450
	h = 500
)

func doAction(title string, w, h float32, group *widget.RadioGroup,
	window fyne.Window, s1 []s.Junior, s2 []s.TechLeader, entry *widget.Entry, a fyne.App) *widget.Button {
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
				juniorForm.Resize(fyne.NewSize(w*0.4, h*0.4))
				email, salary, team, position :=
					widget.NewEntry(), widget.NewEntry(),
					widget.NewEntry(), widget.NewEntry()
				juniorForm.SetContent(container.NewWithoutLayout(
					email, salary, team, position))
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
					err = f.Serialize(uri.Path(), "json", s1, s2)
					if err != nil {
						dialog.ShowError(err, window)
					}
				},
					window,
				)
			} else if group.Selected == "XML" {
				dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
					err = f.Serialize(uri.Path(), "xml", s1, s2)
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
	btn.Resize(fyne.NewSize(w, h))
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
		},
			{Programmer: s.Programmer{Email: "check2@mail.ru", Salary: 30000, Team: 3, Position: "Junior-"},
				Resignation: 5,
			}}
		tLStaff []s.TechLeader

		entry = widget.NewEntry()

		newTechLeadBtn = doAction("hire techLead", w*0.4, h*0.045, radio, window, jStaff, tLStaff, entry, mainApp)
		newJuniorBtn   = doAction("hire another junior", w*0.4, h*0.045, radio, window, jStaff, tLStaff, entry, mainApp)
		parseBtn       = doAction("parse [from]", w*0.4, h*0.045, radio, window, jStaff, tLStaff, entry, mainApp)
		loadBtn        = doAction("load [to folder]", w*0.4, h*0.045, radio, window, jStaff, tLStaff, entry, mainApp)
		itsMe          = widget.NewLabel("telegram @ampheee\n\t\tmgtu stankin")

		total = widget.NewLabel("")
	)
	var check []byte
	check, _ = json.Marshal(jStaff)
	fmt.Println(string(check))

	window.CenterOnScreen()
	window.Resize(fyne.NewSize(w, h))
	window.SetFixedSize(true)
	total.Resize(fyne.NewSize(w*0.5, h*0.1))
	entry.Resize(fyne.NewSize(w*0.8, h*0.25))
	itsMe.Resize(fyne.NewSize(w*0.5, h*0.1))

	entry.Move(fyne.NewPos(w*0.1, h*0.5))
	radio.Move(fyne.NewPos(w*0.10, h*0.15))
	total.Move(fyne.NewPos(w*0.265, h*0.8))
	itsMe.Move(fyne.NewPos(w*0.30, h*0.85))

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
