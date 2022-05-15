package cmp

import (
	"strconv"

	"github.com/rivo/tview"
	"ondrejmaksi.com/db2project/rdg"
)

func NewPaymentForm(order rdg.Order, onSubmit func(order rdg.Order, amount float64), onCancel func()) *tview.Form {
	parseError := false
	amount := 0.0
	orderForm := tview.NewForm().
		AddInputField("Amount", "0.0", 20, nil, func(text string) {
			f, err := strconv.ParseFloat(text, 64)
			if err != nil {
				parseError = true
			}
			parseError = false
			amount = f

		}).
		AddButton("Save", func() {
			if !parseError {
				onSubmit(order, amount)
			}
		}).
		SetCancelFunc(func() {
			onCancel()
		})

	return orderForm
}
