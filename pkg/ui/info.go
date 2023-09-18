package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const logo = `
███████╗██╗███╗   ██╗    ████████╗██████╗  █████╗  ██████╗██╗  ██╗
██╔════╝██║████╗  ██║    ╚══██╔══╝██╔══██╗██╔══██╗██╔════╝██║ ██╔╝
█████╗  ██║██╔██╗ ██║       ██║   ██████╔╝███████║██║     █████╔╝ 
██╔══╝  ██║██║╚██╗██║       ██║   ██╔══██╗██╔══██║██║     ██╔═██╗ 
██║     ██║██║ ╚████║       ██║   ██║  ██║██║  ██║╚██████╗██║  ██╗
╚═╝     ╚═╝╚═╝  ╚═══╝       ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝
                                                                  
`

func (u *Ui) createInfoScreen() tview.Primitive {
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorBlue).
		SetTextAlign(tview.AlignCenter)

	fmt.Fprint(logoBox, logo)

	return center(100, 10, logoBox)
}
