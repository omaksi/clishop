package state

import (
	"ondrejmaksi.com/db2project/rdg"
	"ondrejmaksi.com/db2project/ui/lib"
)

type SearchStatisticsState struct {
	lib.GenericState[[]rdg.SearchStatistic]
}

func NewSearchStatisticsState(ss []rdg.SearchStatistic) *SearchStatisticsState {
	return &SearchStatisticsState{}
}
