package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"local.package/model"
)

func main() {

	model.NewModel().GetUserToken("mimoc")

	// box := tview.NewBox().SetBorder(true).SetTitle(userlist[0].Name)
	// if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	// 	panic(err)
	// }

	app := tview.NewApplication()

	textView := tview.NewTextView()
	textView.SetTitle("textView")
	textView.SetBorder(true)

	inputField := tview.NewInputField()
	inputField.SetLabel("input: ")
	inputField.SetTitle("inputField").
		SetBorder(true)

	inputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			textView.SetText(textView.GetText(true) + inputField.GetText() + "\n")
			inputField.SetText("")
			return nil
		}
		return event
	})

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		AddItem(inputField, 3, 0, true).
		AddItem(textView, 0, 1, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}

}
