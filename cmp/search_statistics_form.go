package cmp

import (
	"github.com/rivo/tview"
)

func NewSearchStatisticsForm(onSubmit func(year string), onCancel func()) *tview.Form {
	year := ""
	userForm := tview.NewForm().
		AddInputField("Year", "", 20, nil, func(text string) {
			year = text
		}).
		AddButton("Save", func() {
			onSubmit(year)
		}).
		SetCancelFunc(func() {
			onCancel()
		})

	return userForm
}
