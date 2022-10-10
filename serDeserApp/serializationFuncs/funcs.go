package serializationFuncs

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testJsonSerDeser/errs"
	s "testJsonSerDeser/structs"
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
			jS, err = json.Marshal(arrTL)
			_, err = file.Write(jS)
			err = file.Close()
		case "xml":
			fPath := filepath.Join(folderPath, "Juniors"+".xml")
			file, err = os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE, 0644)
			jS, err = xml.Marshal(arrTL)
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
						s := fmt.Sprintf("Email: %s, Salary %d, Team: %d, Position: %s, Subordinates: %d",
							employee.Email, employee.Salary, employee.Team, employee.Position, employee.Subordinates)
						sendToEntry = append(sendToEntry, s)
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

func addNewJunior(arr []s.Junior) (err error) {

	return nil
}
