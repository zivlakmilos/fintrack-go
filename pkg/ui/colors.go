package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func setFormColors(form *tview.Form) {
	form.SetFieldTextColor(tcell.ColorDarkBlue)
	form.SetButtonTextColor(tcell.ColorDarkBlue)
}
