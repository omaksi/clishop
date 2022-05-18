package cmp

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"ondrejmaksi.com/db2project/rdg"
)

type ProductTable struct {
	*tview.Table
}

var (
	productTable *ProductTable = nil
)

func NewProductsTable(products []rdg.Product, onEdit func(product rdg.Product), onDelete func(product rdg.Product), onAdd func(user rdg.Product), onDone func()) *ProductTable {
	if productTable == nil {
		productTable = &ProductTable{
			Table: tview.NewTable(),
		}
		productTable.SetBorders(false).
			SetBordersColor(tcell.ColorYellow).
			SetSelectable(true, false)
		productTable.SetFixed(1, 0)
	}

	productTable.SetProps(products, onDelete, onEdit, onAdd, onDone)

	return productTable
}

func (pt *ProductTable) SetProps(products []rdg.Product, onDelete func(user rdg.Product), onEdit func(user rdg.Product), onAdd func(user rdg.Product), onDone func()) {
	pt.SetSelectedFunc(func(row int, col int) {
		// messageState.SetState(fmt.Sprint(row))
	})
	pt.SetDoneFunc(func(key tcell.Key) {
		onDone()
	})

	pt.SetInputCapture(func(eventKey *tcell.EventKey) *tcell.EventKey {
		if eventKey.Key() == tcell.KeyRune && eventKey.Rune() == 'd' {
			row, _ := pt.GetSelection()
			onDelete(products[row-1])
			return nil
		}
		if eventKey.Key() == tcell.KeyRune && eventKey.Rune() == 'e' {
			row, _ := pt.GetSelection()
			onEdit(products[row-1])
			return nil
		}
		if eventKey.Key() == tcell.KeyRune && eventKey.Rune() == 'a' {
			if (products == nil) || (len(products) == 0) {
				onAdd(rdg.Product{})
				return nil
			}
			row, _ := pt.GetSelection()
			onAdd(products[row-1])
			return nil
		}
		return eventKey
	})

	pt.ShowProducts(products)
}

func (pt *ProductTable) ShowProducts(products []rdg.Product) {
	pt.Clear()
	pt.SetFixed(1, 0)

	pt.SetCell(0, 0, tview.NewTableCell("ID").SetTextColor(tcell.ColorBlue).SetSelectable(false))
	pt.SetCell(0, 1, tview.NewTableCell("Name").SetTextColor(tcell.ColorBlue).SetSelectable(false))
	pt.SetCell(0, 2, tview.NewTableCell("Price").SetTextColor(tcell.ColorBlue).SetSelectable(false))
	pt.SetCell(0, 3, tview.NewTableCell("Quantity").SetTextColor(tcell.ColorBlue).SetSelectable(false))

	for i, product := range products {
		pt.SetCell(i+1, 0, tview.NewTableCell(fmt.Sprintf("%d", product.Id)).SetTextColor(tcell.ColorBlue))
		pt.SetCell(i+1, 1, tview.NewTableCell(product.Name))
		pt.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprintf("%.2f", product.Price)))
		pt.SetCell(i+1, 3, tview.NewTableCell(fmt.Sprintf("%d", product.Quantity)))
	}
}
