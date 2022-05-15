package state

import "ondrejmaksi.com/db2project/ui/lib"

type ContentState struct {
	lib.GenericState[string]
}

func NewContentState(value string) *ContentState {
	return &ContentState{}
}
