package state

import "ondrejmaksi.com/db2project/ui/lib"

type TitleState struct {
	lib.GenericState[string]
}

func NewTitleState(value string) *TitleState {
	return &TitleState{}
}
