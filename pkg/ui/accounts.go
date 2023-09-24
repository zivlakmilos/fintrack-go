package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

type AccountsScreen struct {
}

func (u *Ui) createAccountsScreen() tview.Primitive {
	filter := tview.NewList().
		AddItem("All", "", '*', nil).
		AddItem("Assets", "", 'a', nil).
		AddItem("liability", "", 'l', nil).
		AddItem("Income", "", 'i', nil).
		AddItem("Expense", "", 'e', nil)

	filter.SetBorder(true)
	filter.SetTitle(" Filter (Ctrl+F) ")

	help := tview.NewTextView()

	fmt.Fprintf(help, "%s\n", "New account - Ctrl + N")
	fmt.Fprintf(help, "%s\n", "Filter      - Ctrl + F")
	fmt.Fprintf(help, "%s\n", "Remove      - Del")

	help.SetBorder(true)
	help.SetTitle(" Help ")

	header := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(filter, 25, 0, true).
		AddItem(help, 30, 0, false).
		AddItem(nil, 0, 1, false)

	tbl := tview.NewTable()

	tbl.SetBorder(true)
	tbl.SetTitle(" Accounts ")

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(header, 10, 0, true).
		AddItem(tbl, 0, 1, false)

	layout.SetBorder(true)
	layout.SetTitle(" Accounts ")

	return layout
}
