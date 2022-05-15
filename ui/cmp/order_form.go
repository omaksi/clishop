package cmp

import (
	"github.com/rivo/tview"
)

func NewOrderForm(userId int, onSubmit func(userId int, address string), onCancel func()) *tview.Form {
	address := ""
	orderForm := tview.NewForm().
		AddInputField("Address", "", 20, nil, func(text string) {
			address = text
		}).
		AddButton("Save", func() {
			onSubmit(userId, address)
		}).
		SetCancelFunc(func() {
			onCancel()
		})

	return orderForm
}
