package state

import (
	"github.com/rivo/tview"
	"ondrejmaksi.com/db2project/ui/lib"
)

type FocusState struct {
	*lib.GenericState[string]
	primitives map[string]tview.Primitive
}

func (s *FocusState) Menu() {
	s.SetState("menu")
}
func (s *FocusState) Content() {
	s.SetState("content")
}
func (f *FocusState) AddFocusTarget(name string, p tview.Primitive) {
	f.primitives[name] = p
}
func (f *FocusState) GetPrimitives() map[string]tview.Primitive {
	return f.primitives
}

func NewFocusState(value string) *FocusState {
	newState := lib.NewGenericState(value)

	focusState := &FocusState{
		newState,
		make(map[string]tview.Primitive),
	}

	return focusState
}
