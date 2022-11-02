package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Contact struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	state       string
	business    bool
}

var contacts []Contact

var app = tview.NewApplication()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(q) to quit")

var form = tview.NewForm()

func addContactForm() {
	contact := Contact{}

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		contact.firstName = firstName
	})

	form.AddInputField("Last Name", "", 20, nil, func(lastName string) {
		contact.lastName = lastName
	})

	form.AddInputField("Email", "", 20, nil, func(email string) {
		contact.email = email
	})

	form.AddInputField("Phone", "", 20, nil, func(phone string) {
		contact.phoneNumber = phone
	})

	form.AddDropDown("State", states, 0, func(state string, index int) {
		contact.state = state
	})

	form.AddCheckbox("Business", false, func(business bool) {
		contact.business = business
	})

	form.AddButton("Save", func() {
		contacts = append(contacts, contact)
	})
}

func main() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		}
		return event
	})

	if err := app.SetRoot(text, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
