package cmp

import "github.com/rivo/tview"

func NewLog() *tview.List {
	log := tview.NewList()
	log.ShowSecondaryText(false).SetTitle("Log").SetBorder(true)
	// log.AddItem("This is the log", "", 0, nil)
	return log
}
