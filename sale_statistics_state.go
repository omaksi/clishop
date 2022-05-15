package state

import (
	"ondrejmaksi.com/db2project/rdg"
	"ondrejmaksi.com/db2project/ui/lib"
)

type SaleStatisticsState struct {
	lib.GenericState[[]rdg.SaleStatistic]
}

func NewSaleStatisticsState(ss []rdg.SaleStatistic) *SaleStatisticsState {
	return &SaleStatisticsState{}
}
