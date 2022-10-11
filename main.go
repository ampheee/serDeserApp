package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	funcs "testJsonSerDeser/serializationFuncs"
	s "testJsonSerDeser/structs"
	"time"
)

const (
	w = 450
	h = 500
)

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

		jStaff  []s.Junior
		tLStaff []s.TechLeader

		entry = widget.NewEntry()

		newTechLeadBtn = funcs.DoAction("hire techLead", radio, window, &jStaff, &tLStaff, entry, mainApp)
		newJuniorBtn   = funcs.DoAction("hire another junior", radio, window, &jStaff, &tLStaff, entry, mainApp)
		parseBtn       = funcs.DoAction("parse [from]", radio, window, &jStaff, &tLStaff, entry, mainApp)
		loadBtn        = funcs.DoAction("load [to folder]", radio, window, &jStaff, &tLStaff, entry, mainApp)
		itsMe          = widget.NewLabel("telegram @ampheee\n\t\tmgtu stankin")

		total = widget.NewLabel("")
	)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(w, h))
	window.SetFixedSize(true)
	funcs.ResizeAndMove(total, w*0.275, h*0.775, w*0.5, h*0.1)
	funcs.ResizeAndMove(entry, w*0.1, h*0.5, w*0.8, h*0.25)
	funcs.ResizeAndMove(itsMe, w*0.31, h*0.825, w*0.5, h*0.1)
	radio.Move(fyne.NewPos(w*0.10, h*0.16))

	funcs.ResizeAndMove(newTechLeadBtn, w*0.5, h*0.16, w*0.4, h*0.07)
	funcs.ResizeAndMove(newJuniorBtn, w*0.5, h*0.23, w*0.4, h*0.07)
	funcs.ResizeAndMove(parseBtn, w*0.5, h*0.30, w*0.4, h*0.07)
	funcs.ResizeAndMove(loadBtn, w*0.5, h*0.37, w*0.4, h*0.07)

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
