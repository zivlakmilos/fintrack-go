package ui

import (
	"github.com/gdamore/tcell/v2"
)

type menuItem struct {
	id    string
	title string
	key   tcell.Key
}

var mainMenu = []menuItem{
	{id: "F1", key: tcell.KeyF1, title: "Info"},
	{id: "F2", key: tcell.KeyF2, title: "New Income"},
	{id: "F3", key: tcell.KeyF3, title: "New Expense"},
	{id: "F4", key: tcell.KeyF4, title: "Repports"},
	{id: "F5", key: tcell.KeyF5, title: "Graphs"},
	{id: "F6", key: tcell.KeyF6, title: "Accounts"},
	{id: "F9", key: tcell.KeyF9, title: "Open Year"},
	{id: "F10", key: tcell.KeyF10, title: "Exit"},
}
