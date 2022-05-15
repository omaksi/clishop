package cmp

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"ondrejmaksi.com/db2project/rdg"
)

type OrderTable struct {
	*tview.Table
}

var (
	orderTable *OrderTable = nil
)

/*
Creates a new table for displaying a list of orders
*/
func NewOrdersTable(orders []rdg.Order, onCancel func(order rdg.Order), onPay func(order rdg.Order), onExpedite func(order rdg.Order), onDone func()) *OrderTable {
	if orderTable == nil {
		orderTable = &OrderTable{
			Table: tview.NewTable(),
		}
		orderTable.SetBorders(false).
			SetBordersColor(tcell.ColorYellow).
			SetSelectable(true, false)
		orderTable.SetFixed(1, 0)
	}

	orderTable.setProps(orders, onCancel, onPay, onExpedite, onDone)

	return orderTable
}

func (pt *OrderTable) setProps(orders []rdg.Order, onCancel func(order rdg.Order), onPay func(order rdg.Order), onExpedite func(order rdg.Order), onDone func()) {
	pt.SetSelectedFunc(func(row int, col int) {
		// messageState.SetState(fmt.Sprint(row))
	})
	pt.SetDoneFunc(func(key tcell.Key) {
		onDone()
	})

	pt.SetInputCapture(func(eventKey *tcell.EventKey) *tcell.EventKey {
		if eventKey.Key() == tcell.KeyRune && eventKey.Rune() == 'c' {
			row, _ := pt.GetSelection()
			onCancel(orders[row-1])
			return nil
		}
		if eventKey.Key() == tcell.KeyRune && eventKey.Rune() == 'p' {
			row, _ := pt.GetSelection()
			onPay(orders[row-1])
			return nil
		}
		if eventKey.Key() == tcell.KeyRune && eventKey.Rune() == 'e' {
			row, _ := pt.GetSelection()
			onExpedite(orders[row-1])
			return nil
		}
		return eventKey
	})

	pt.showOrders(orders)
}

func (pt *OrderTable) showOrders(orders []rdg.Order) {
	pt.Clear()
	pt.SetFixed(1, 0)

	pt.SetCell(0, 0, tview.NewTableCell("ID").SetTextColor(tcell.ColorRed).SetSelectable(false))
	pt.SetCell(0, 1, tview.NewTableCell("UserId").SetTextColor(tcell.ColorRed).SetSelectable(false))
	pt.SetCell(0, 2, tview.NewTableCell("Address").SetTextColor(tcell.ColorRed).SetSelectable(false))
	pt.SetCell(0, 3, tview.NewTableCell("Total").SetTextColor(tcell.ColorRed).SetSelectable(false))
	pt.SetCell(0, 4, tview.NewTableCell("Status").SetTextColor(tcell.ColorRed).SetSelectable(false))

	for i, order := range orders {
		pt.SetCell(i+1, 0, tview.NewTableCell(fmt.Sprintf("%d", order.Id)).SetTextColor(tcell.ColorRed))
		pt.SetCell(i+1, 1, tview.NewTableCell(fmt.Sprintf("%d", order.UserId)))
		pt.SetCell(i+1, 2, tview.NewTableCell(order.Address))
		pt.SetCell(i+1, 3, tview.NewTableCell(fmt.Sprintf("%.2f", order.Total)))
		pt.SetCell(i+1, 4, tview.NewTableCell(order.Status))
	}
}
