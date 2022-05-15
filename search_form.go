package cmp

import (
	"github.com/rivo/tview"
)

func NewSearchForm(onSubmit func(searchText string), onCancel func()) *tview.Form {
	searchText := ""
	userForm := tview.NewForm().
		AddInputField("Search", "", 20, nil, func(text string) {
			searchText = text
		}).
		AddButton("Save", func() {
			onSubmit(searchText)
		}).
		SetCancelFunc(func() {
			onCancel()
		})

	return userForm
}
