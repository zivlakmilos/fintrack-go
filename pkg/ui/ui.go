package ui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Ui struct {
	app   *tview.Application
	pages *tview.Pages
	menu  *tview.TextView
}

func NewUi() (*Ui, error) {
	ui := Ui{}

	ui.app = tview.NewApplication()
	ui.pages = tview.NewPages()

	ui.createPages()

	layout := ui.createLayout()

	ui.app.SetRoot(layout, true)

	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for idx, item := range mainMenu {
			if event.Key() == item.key {
				ui.menu.Highlight(strconv.Itoa(idx)).ScrollToHighlight()
				return nil
			}
		}

		return event
	})

	ui.menu.Highlight("0")

	return &ui, nil
}

func (u *Ui) Run() error {
	return u.app.Run()
}

func (u *Ui) createLayout() *tview.Grid {
	header := u.createHeader()
	footer := u.createFooter()

	u.menu = footer

	grid := tview.NewGrid().
		SetRows(3, 0, 1).
		SetColumns(0).
		AddItem(header, 0, 0, 1, 1, 0, 0, false).
		AddItem(u.pages, 1, 0, 1, 1, 0, 0, true).
		AddItem(footer, 2, 0, 1, 1, 0, 0, false).
		SetBorders(true)

	return grid
}

func (u *Ui) createHeader() *tview.Grid {
	title := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("[yellow]FinTrack")
	copyright := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText(fmt.Sprintf("[darkcyan]Copyright (c) 2023 - %d Milos Zivlak", time.Now().Year()))

	header := tview.NewGrid().
		SetRows(0, 1).
		SetColumns(0).
		AddItem(title, 0, 0, 1, 1, 0, 0, false).
		AddItem(copyright, 1, 0, 1, 1, 0, 0, false)

	return header
}

func (u *Ui) createFooter() *tview.TextView {
	footer := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			u.pages.SwitchToPage(added[0])
		})

	for idx, item := range mainMenu {
		fmt.Fprintf(footer, `%s ["%d"][darkcyan]%s[white][""]  `, item.id, idx, item.title)
	}

	return footer
}

func (u *Ui) createPages() {
	u.pages.AddPage("1", createIncomeScreen(), true, false)
}
