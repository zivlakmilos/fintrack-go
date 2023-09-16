package ui

import (
	"github.com/gdamore/tcell/v2"
)

type menuItem struct {
	id    int
	title string
	key   tcell.Key
}

var mainMenu = []menuItem{
	{id: 1, key: tcell.KeyF1, title: "Info"},
	{id: 2, key: tcell.KeyF2, title: "New Income"},
	{id: 3, key: tcell.KeyF3, title: "New Expense"},
	{id: 4, key: tcell.KeyF4, title: "Repports"},
	{id: 5, key: tcell.KeyF5, title: "Graphs"},
	{id: 6, key: tcell.KeyF6, title: "Accounts"},
	{id: 9, key: tcell.KeyF9, title: "Open Year"},
	{id: 0, key: tcell.KeyF10, title: "Exit"},
}
