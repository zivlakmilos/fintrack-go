package ui

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/zivlakmilos/fintrack-go/pkg/core"
	"github.com/zivlakmilos/fintrack-go/pkg/db"
)

type Ui struct {
	app     *tview.Application
	pages   *tview.Pages
	menu    *tview.TextView
	db      *sql.DB
	screens []func() tview.Primitive
}

func NewUi() (*Ui, error) {
	ui := Ui{}

	err := ui.loadData()
	if err != nil {
		return nil, err
	}

	ui.app = tview.NewApplication()
	ui.pages = tview.NewPages()

	ui.createPages()

	layout := ui.createLayout()

	ui.app.SetRoot(layout, true)

	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for idx, item := range mainMenu {
			if event.Key() == item.key {
				ui.showPage(idx)
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

func (u *Ui) loadData() error {
	config, err := core.LoadConfig()
	if err != nil {
		return err
	}

	dbPath, err := core.GetDBPath(config.Year)
	if err != nil {
		return err
	}

	u.db, err = db.Open(dbPath)
	if err != nil {
		return err
	}

	err = db.InitDB(u.db)
	if err != nil {
		return err
	}

	return nil
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
			idx, _ := strconv.Atoi(added[0])
			u.pages.AddPage(added[0], u.screens[idx](), true, false)
			u.pages.SwitchToPage(added[0])
		})

	for idx, item := range mainMenu {
		fmt.Fprintf(footer, `%s ["%d"][darkcyan]%s[white][""]  `, item.id, idx, item.title)
	}

	return footer
}

func (u *Ui) createPages() {
	u.screens = []func() tview.Primitive{
		u.createInfoScreen,
		u.createIncomeScreen,
		u.createExpenseScreen,
		nil,
		nil,
		u.createAccountsScreen,
		nil,
		nil,
	}
}

func (u *Ui) showPage(idx int) {
	u.menu.Highlight(strconv.Itoa(idx)).ScrollToHighlight()
}
