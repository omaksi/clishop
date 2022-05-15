package cmp

import "github.com/rivo/tview"

func NewStatusBar() (*tview.Flex, *tview.TextView, *tview.TextView) {
	statusbar := tview.NewFlex()
	statusbar.SetDirection(tview.FlexColumn)

	status := tview.NewTextView().
		SetDynamicColors(true).
		SetText("This is the status bar!")

	status.SetText("Hello, world!")

	keyhints := tview.NewTextView()

	statusbar.AddItem(status, 0, 1, false)
	statusbar.AddItem(keyhints, 0, 4, false)

	return statusbar, status, keyhints
}
