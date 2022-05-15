package cmp

import (
	"strconv"

	"github.com/rivo/tview"
	"ondrejmaksi.com/db2project/rdg"
)

func NewProductForm(product rdg.Product, onSubmit func(product rdg.Product), onCancel func()) *tview.Form {
	parseError := false
	productForm := tview.NewForm().
		AddInputField("Name", product.Name, 20, nil, func(text string) {
			product.Name = text
		}).
		AddInputField("Price", strconv.FormatFloat(product.Price, 'f', 2, 64), 20, nil, func(text string) {
			f, err := strconv.ParseFloat(text, 64)
			if err != nil {
				parseError = true
			}
			parseError = false
			product.Price = f
		}).
		AddInputField("Quantity", strconv.FormatUint(uint64(product.Quantity), 10), 20, nil, func(text string) {
			u, err := strconv.ParseUint(text, 10, 64)
			if err != nil {
				parseError = true
			}
			parseError = false
			product.Quantity = uint(u)
		}).
		AddButton("Save", func() {
			if !parseError {
				onSubmit(product)
			}
		}).
		SetCancelFunc(func() {
			onCancel()
		})

	return productForm
}
