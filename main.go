package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"testJsonSerDeser/funcs/mainFuncs"
	s "testJsonSerDeser/structs"
	"time"
)

const (
	w = 450
	h = 500
)

func main() {
	var (
		mainApp = app.New()
		window  = mainApp.NewWindow("First Lab")

		radio = widget.NewRadioGroup([]string{
			"XML", "JSON", "To memory",
		}, func(s string) {})

		jStaff  []s.Junior
		tLStaff []s.TechLeader

		entry = widget.NewEntry()

		newTechLeadBtn = mainFuncs.DoAction("hire techLead", radio, window, &jStaff, &tLStaff, entry, mainApp)
		newJuniorBtn   = mainFuncs.DoAction("hire another junior", radio, window, &jStaff, &tLStaff, entry, mainApp)
		parseBtn       = mainFuncs.DoAction("parse [from]", radio, window, &jStaff, &tLStaff, entry, mainApp)
		loadBtn        = mainFuncs.DoAction("load [to folder]", radio, window, &jStaff, &tLStaff, entry, mainApp)
		fio            = widget.NewLabel("Nazipov Rustam IDB-21-11")
		itsMe          = widget.NewLabel("telegram @ampheee\n\t\tmgtu stankin")

		total = widget.NewLabel("")
	)
	fio.Move(fyne.NewPos(w*0.275, h*0.05))
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(w, h))
	window.SetFixedSize(true)
	mainFuncs.ResizeAndMove(total, w*0.275, h*0.775, w*0.5, h*0.1)
	mainFuncs.ResizeAndMove(entry, w*0.1, h*0.5, w*0.8, h*0.25)
	mainFuncs.ResizeAndMove(itsMe, w*0.31, h*0.825, w*0.5, h*0.1)
	radio.Move(fyne.NewPos(w*0.10, h*0.16))

	mainFuncs.ResizeAndMove(newTechLeadBtn, w*0.5, h*0.16, w*0.4, h*0.07)
	mainFuncs.ResizeAndMove(newJuniorBtn, w*0.5, h*0.23, w*0.4, h*0.07)
	mainFuncs.ResizeAndMove(parseBtn, w*0.5, h*0.30, w*0.4, h*0.07)
	mainFuncs.ResizeAndMove(loadBtn, w*0.5, h*0.37, w*0.4, h*0.07)

	window.SetContent(container.NewWithoutLayout(
		fio,
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
			mainFuncs.TotalStuff(total, jStaff, tLStaff)
		}
	}()

	window.ShowAndRun()
}
