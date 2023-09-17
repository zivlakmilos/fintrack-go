package ui

import (
	"github.com/rivo/tview"
)

func createIncomeScreen() tview.Primitive {
	form := tview.NewForm().
		AddDropDown("Income", []string{"Test 1", "Test 2"}, 0, nil).
		AddDropDown("Account", []string{"Test adflkj", "dfa"}, 0, nil).
		AddInputField("Amount", "test", 20, nil, nil).
		AddButton("Save", nil).
		AddButton("Cancel", nil)

	setFormColors(form)

	form.SetBorder(true)
	form.SetTitle(" New Income ")

	return center(30, 12, form)
}
