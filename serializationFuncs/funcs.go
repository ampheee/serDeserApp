package serializationFuncs

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testJsonSerDeser/errs"
	s "testJsonSerDeser/structs"
)

const (
	w = 450
	h = 500
)

func Serialize(folderPath, mode string, arrJ []s.Junior, arrTL []s.TechLeader) (err error) {
	defer func() { err = errs.WrapIfErr("cant serialize", err) }()
	if len(arrTL) == 0 && len(arrJ) == 0 {
		return errors.New("there`s no jun`s and tl`s")
	}
	if len(arrTL) != 0 {
		var (
			file *os.File
			tlS  []byte
		)
		switch mode {
		case "json":
			fPath := filepath.Join(folderPath, "TechLeads"+".json")
			file, err = os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE, 0644)
			tlS, err = json.Marshal(arrTL)
			_, err = file.Write(tlS)
			err = file.Close()
		case "xml":
			fPath := filepath.Join(folderPath, "TechLeads"+".xml")
			file, err = os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE, 0644)
			tlS, err = xml.Marshal(arrTL)
			_, err = file.Write(tlS)
			err = file.Close()
		}
	}
	if len(arrJ) != 0 {
		var (
			file *os.File
			jS   []byte
		)
		switch mode {
		case "json":
			fPath := filepath.Join(folderPath, "Juniors"+".json")
			file, err = os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE, 0644)
			jS, err = json.Marshal(arrJ)
			_, err = file.Write(jS)
			err = file.Close()
		case "xml":
			fPath := filepath.Join(folderPath, "Juniors"+".xml")
			file, err = os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE, 0644)
			jS, err = xml.Marshal(arrJ)
			_, err = file.Write(jS)
			err = file.Close()
		}
	}
	return err
}

func Deserialize(mode string, entry *widget.Entry, window fyne.Window) (err error) {
	defer func() { err = errs.WrapIfErr("something went wrong:", err) }()
	dialog.ShowFileOpen(
		func(reader fyne.URIReadCloser, err error) {
			switch mode {
			case "xml":
				if !strings.HasSuffix(reader.URI().Path(), ".xml") {
					dialog.ShowError(errors.New("it`s not a XML file"), window)
				} else if strings.HasSuffix(reader.URI().Path(), "Juniors.xml") {
					var (
						data, _     = io.ReadAll(reader)
						juniors     []s.Junior
						sendToEntry []string
					)
					err = xml.Unmarshal(data, &juniors)
					for _, employee := range juniors {
						str := fmt.Sprintf("Email: %s, Salary %d, Team: %d, Position: %s, Resignation: %d",
							employee.Email, employee.Salary, employee.Team, employee.Position, employee.Resignation)
						sendToEntry = append(sendToEntry, str)
					}
					final := strings.Join(sendToEntry, "\n")
					entry.SetText(final)
				} else if strings.HasSuffix(reader.URI().Path(), "TechLeads.xml") {
					var (
						data, _     = io.ReadAll(reader)
						teamLeaders []s.TechLeader
						sendToEntry []string
					)
					err = xml.Unmarshal(data, &teamLeaders)
					if err != nil {
						return
					}
					for _, employee := range teamLeaders {
						str := fmt.Sprintf("Email: %s, Salary: %d, Team: %d, CurrentProject: %s,"+
							"Position: %s, Subordinates: %d",
							employee.Email, employee.Salary, employee.Team, employee.CurrentProject,
							employee.Position, employee.Subordinates)
						sendToEntry = append(sendToEntry, str)
					}
					final := strings.Join(sendToEntry, "\n")
					entry.SetText(final)
				}
			case "json":
				if !strings.HasSuffix(reader.URI().Path(), ".json") {
					dialog.ShowError(errors.New("it`s not a JSON file"), window)
				} else if strings.HasSuffix(reader.URI().Path(), "Juniors.json") {
					var (
						data, _     = io.ReadAll(reader)
						juniors     []s.Junior
						sendToEntry []string
					)
					err = json.Unmarshal(data, &juniors)
					for _, employee := range juniors {
						str := fmt.Sprintf("Email: %s, Salary %d, Team: %d, Position: %s, Resignation: %d",
							employee.Email, employee.Salary, employee.Team, employee.Position, employee.Resignation)
						sendToEntry = append(sendToEntry, str)
					}
					final := strings.Join(sendToEntry, "\n")
					entry.SetText(final)
				} else if strings.HasSuffix(reader.URI().Path(), "TechLeads.json") {
					var (
						data, _     = io.ReadAll(reader)
						teamLeaders []s.TechLeader
						sendToEntry []string
					)
					err = json.Unmarshal(data, &teamLeaders)
					if err != nil {
						return
					}
					for _, employee := range teamLeaders {
						str := fmt.Sprintf("Email: %s, Salary %d, Team: %d, Position: %s, Subordinates: %d",
							employee.Email, employee.Salary, employee.Team, employee.Position, employee.Subordinates)
						sendToEntry = append(sendToEntry, str)
					}
					final := strings.Join(sendToEntry, "\n")
					entry.SetText(final)
				}
			}
		},
		window,
	)
	return err
}

func DoAction(title string, group *widget.RadioGroup,
	window fyne.Window, juniors *[]s.Junior, techLeads *[]s.TechLeader,
	entry *widget.Entry, a fyne.App) *widget.Button {
	btn := widget.NewButton(title, func() {
		switch title {
		case "hire techLead":
			if group.Selected == "To memory" {
				techLeadForm := a.NewWindow("hireTechLeadForm")
				techLeadForm.Resize(fyne.NewSize(w*0.65, h*0.675))
				eEntry, sEntry, tEntry, pEntry, cpEntry, sbEntry :=
					widget.NewEntry(), widget.NewEntry(),
					widget.NewEntry(), widget.NewEntry(), widget.NewEntry(), widget.NewEntry()

				email, salary, team, position, project, subordinates :=
					widget.NewLabel("Email:"), widget.NewLabel("Salary:"),
					widget.NewLabel("Team:"), widget.NewLabel("Position:"),
					widget.NewLabel("Project:"), widget.NewLabel("Subordinates:")
				hire := widget.NewButton("HIRE!", func() {
					var (
						newTLead         s.TechLeader
						err1, err2, err3 error
					)
					intSalary, err1 := strconv.Atoi(sEntry.Text)
					intTeam, err2 := strconv.Atoi(tEntry.Text)
					intSubordinates, err3 := strconv.Atoi(sbEntry.Text)
					if err1 == nil && err2 == nil && err3 == nil {
						newTLead.Email = eEntry.Text
						newTLead.Team = intTeam
						newTLead.Salary = intSalary
						newTLead.Position = pEntry.Text
						newTLead.CurrentProject = cpEntry.Text
						newTLead.Subordinates = intSubordinates
						*techLeads = append(*techLeads, newTLead)
						techLeadForm.Close()
					} else {
						var (
							allErrs  = []error{err1, err2, err3}
							warnings []string
							finalErr string
						)
						for i, err := range allErrs {
							if err != nil && i == 0 {
								warnings = append(warnings, "salary")
							}
							if err != nil && i == 1 {
								warnings = append(warnings, "number of team")
							}
							if err != nil && i == 2 {
								warnings = append(warnings, "subordinates")
							}
						}
						if len(warnings) == 1 {
							finalErr = fmt.Sprintf("please, enter valid %s", warnings[0])
						} else if len(warnings) == 2 {
							finalErr = fmt.Sprintf("please enter valid\n%s and %s", warnings[0], warnings[1])
						} else if len(warnings) == 3 {
							finalErr = fmt.Sprintf("please enter valid\n%s, \n%s and %s", warnings[0],
								warnings[1], warnings[3])
						}
						dialog.ShowError(errors.New(finalErr), techLeadForm)
					}
				})

				email.Move(fyne.NewPos(w*0.01, h*0.03))
				salary.Move(fyne.NewPos(w*0.01, h*0.11))
				team.Move(fyne.NewPos(w*0.01, h*0.19))
				position.Move(fyne.NewPos(w*0.01, h*0.27))
				project.Move(fyne.NewPos(w*0.01, h*0.35))
				subordinates.Move(fyne.NewPos(w*0.01, h*0.45))

				ResizeAndMove(hire, w*0.013, h*0.55, w*0.6, h*0.09)
				ResizeAndMove(eEntry, w*0.18, h*0.04, w*0.425, h*0.075)
				ResizeAndMove(sEntry, w*0.18, h*0.12, w*0.425, h*0.075)
				ResizeAndMove(tEntry, w*0.18, h*0.20, w*0.425, h*0.075)
				ResizeAndMove(pEntry, w*0.18, h*0.28, w*0.425, h*0.075)
				ResizeAndMove(cpEntry, w*0.18, h*0.36, w*0.425, h*0.075)
				ResizeAndMove(sbEntry, w*0.285, h*0.44, w*0.32, h*0.075)

				techLeadForm.SetContent(container.NewWithoutLayout(email, salary, team, position, project, subordinates,
					eEntry, sEntry, tEntry, pEntry, cpEntry, sbEntry,
					hire))
				techLeadForm.SetFixedSize(true)
				techLeadForm.CenterOnScreen()
				techLeadForm.Show()
			} else {
				dialog.ShowError(errors.New("choose 'To memory'"), window)
			}
		case "hire another junior":
			if group.Selected == "To memory" {
				juniorForm := a.NewWindow("hireJuniorForm")
				juniorForm.Resize(fyne.NewSize(w*0.65, h*0.6))
				eEntry, sEntry, tEntry, pEntry, rEntry :=
					widget.NewEntry(), widget.NewEntry(),
					widget.NewEntry(), widget.NewEntry(), widget.NewEntry()

				email, salary, team, position, resignation :=
					widget.NewLabel("Email:"), widget.NewLabel("Salary:"),
					widget.NewLabel("Team:"), widget.NewLabel("Position:"),
					widget.NewLabel("Resignation:")
				hire := widget.NewButton("HIRE!", func() {
					var (
						newJun           s.Junior
						err1, err2, err3 error
					)
					intSalary, err1 := strconv.Atoi(sEntry.Text)
					intTeam, err2 := strconv.Atoi(tEntry.Text)
					intResignation, err3 := strconv.Atoi(rEntry.Text)
					if err1 == nil && err2 == nil && err3 == nil {
						newJun.Email = eEntry.Text
						newJun.Team = intTeam
						newJun.Salary = intSalary
						newJun.Position = pEntry.Text
						newJun.Resignation = intResignation
						*juniors = append(*juniors, newJun)
						juniorForm.Close()
					} else {
						var (
							allErrs  = []error{err1, err2, err3}
							warnings []string
							finalErr string
						)
						for i, err := range allErrs {
							if err != nil && i == 0 {
								warnings = append(warnings, "salary")
							}
							if err != nil && i == 1 {
								warnings = append(warnings, "number of team")
							}
							if err != nil && i == 2 {
								warnings = append(warnings, "percent of resignation")
							}
						}
						if len(warnings) == 1 {
							finalErr = fmt.Sprintf("please, enter valid %s", warnings[0])
						} else if len(warnings) == 2 {
							finalErr = fmt.Sprintf("please enter valid\n%s and %s", warnings[0], warnings[1])
						} else if len(warnings) == 3 {
							finalErr = fmt.Sprintf("please enter valid\n%s, \n%s and %s", warnings[0],
								warnings[1], warnings[3])
						}
						dialog.ShowError(errors.New(finalErr), juniorForm)
					}
				})

				email.Move(fyne.NewPos(w*0.01, h*0.03))
				salary.Move(fyne.NewPos(w*0.01, h*0.11))
				team.Move(fyne.NewPos(w*0.01, h*0.19))
				position.Move(fyne.NewPos(w*0.01, h*0.27))
				resignation.Move(fyne.NewPos(w*0.01, h*0.37))

				ResizeAndMove(hire, w*0.013, h*0.475, w*0.6, h*0.09)
				ResizeAndMove(eEntry, w*0.18, h*0.04, w*0.425, h*0.075)
				ResizeAndMove(sEntry, w*0.18, h*0.12, w*0.425, h*0.075)
				ResizeAndMove(tEntry, w*0.18, h*0.20, w*0.425, h*0.075)
				ResizeAndMove(pEntry, w*0.18, h*0.28, w*0.425, h*0.075)
				ResizeAndMove(rEntry, w*0.30, h*0.36, w*0.305, h*0.075)

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
				err := Deserialize("json", entry, window)
				if err != nil {
					dialog.ShowError(err, window)
				}
			} else if group.Selected == "XML" {
				err := Deserialize("xml", entry, window)
				if err != nil {
					dialog.ShowError(err, window)
				}
			} else {
				dialog.ShowError(errors.New("choose XML or JSON"), window)
			}
		case "load [to folder]":
			if group.Selected == "JSON" {
				dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
					err = Serialize(uri.Path(), "json", *juniors, *techLeads)
					if err != nil {
						dialog.ShowError(err, window)
					}
				},
					window,
				)
			} else if group.Selected == "XML" {
				dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
					err = Serialize(uri.Path(), "xml", *juniors, *techLeads)
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

func ResizeAndMove(object fyne.CanvasObject, mvPosX, mvPosy, rSizePosX, rSizePosY float32) {
	object.Resize(fyne.NewSize(rSizePosX, rSizePosY))
	object.Move(fyne.NewPos(mvPosX, mvPosy))
}
