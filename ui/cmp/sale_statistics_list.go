package cmp

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"ondrejmaksi.com/db2project/rdg"
)

type SaleStatisticTable struct {
	*tview.Table
}

var (
	saleStatisticTable *SaleStatisticTable = nil
)

func NewSaleStatisticsTable(saleStatistics []rdg.SaleStatistic, onDone func()) *SaleStatisticTable {
	if saleStatisticTable == nil {
		saleStatisticTable = &SaleStatisticTable{
			Table: tview.NewTable(),
		}
		saleStatisticTable.SetBorders(false).
			SetBordersColor(tcell.ColorYellow).
			SetSelectable(true, false)
		saleStatisticTable.SetFixed(1, 0)
	}

	saleStatisticTable.SetProps(saleStatistics, onDone)

	return saleStatisticTable
}

func (sst *SaleStatisticTable) SetProps(saleStatistics []rdg.SaleStatistic, onDone func()) {

	sst.SetDoneFunc(func(key tcell.Key) {
		onDone()
	})

	sst.ShowSaleStatistics(saleStatistics)
}

func (sst *SaleStatisticTable) ShowSaleStatistics(saleStatistics []rdg.SaleStatistic) {
	sst.Clear()
	sst.SetFixed(1, 0)

	sst.SetCell(0, 0, tview.NewTableCell("Attribute Count").SetTextColor(tcell.ColorYellow).SetSelectable(false))
	sst.SetCell(0, 1, tview.NewTableCell("Total Sold").SetTextColor(tcell.ColorYellow).SetSelectable(false))
	for i, s := range saleStatistics {
		sst.SetCell(i+1, 0, tview.NewTableCell(fmt.Sprintf("%d", s.AttributeCount)).SetTextColor(tcell.ColorYellow))
		sst.SetCell(i+1, 1, tview.NewTableCell(fmt.Sprintf("%d", s.TotalSold)))
	}
}
