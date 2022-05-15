package cmp

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"ondrejmaksi.com/db2project/rdg"
)

type SearchStatisticTable struct {
	*tview.Table
}

var (
	searchStatisticTable *SearchStatisticTable = nil
)

func NewSearchStatisticsTable(searchStatistics []rdg.SearchStatistic, onDone func()) *SearchStatisticTable {
	if searchStatisticTable == nil {
		searchStatisticTable = &SearchStatisticTable{
			Table: tview.NewTable(),
		}
		searchStatisticTable.SetBorders(false).
			SetBordersColor(tcell.ColorYellow).
			SetSelectable(true, false)
		searchStatisticTable.SetFixed(1, 0)
	}

	searchStatisticTable.SetProps(searchStatistics, onDone)

	return searchStatisticTable
}

func (sst *SearchStatisticTable) SetProps(searchStatistics []rdg.SearchStatistic, onDone func()) {

	sst.SetDoneFunc(func(key tcell.Key) {
		onDone()
	})

	sst.ShowSearchStatistics(searchStatistics)
}

func (sst *SearchStatisticTable) ShowSearchStatistics(searchStatistics []rdg.SearchStatistic) {
	sst.Clear()
	sst.SetFixed(1, 0)

	sst.SetCell(0, 0, tview.NewTableCell("Month").SetTextColor(tcell.ColorYellow))
	sst.SetCell(0, 1, tview.NewTableCell("Share").SetTextColor(tcell.ColorYellow))
	for i, s := range searchStatistics {
		sst.SetCell(i+1, 0, tview.NewTableCell(s.Month.Format("2006-01")).SetTextColor(tcell.ColorYellow))
		if s.Share.Valid {
			sst.SetCell(i+1, 1, tview.NewTableCell(fmt.Sprintf("%.2f", s.Share.Float64)))
		} else {
			sst.SetCell(i+1, 1, tview.NewTableCell("null"))
		}
		// sst.SetCell(i+1, 1, tview.NewTableCell(strconv.FormatFloat(s.Share, 'f', 2, 64)))
	}
}
