package state

import "ondrejmaksi.com/db2project/ui/lib"

type MessageState struct {
	*lib.GenericState[string]
}

func (s *MessageState) SetMessage(message string) {
	s.SetState(message)
}
func (s *MessageState) Success(message string) {
	s.SetState("✅ " + message)
}
func (s *MessageState) Fail(message string) {
	s.SetState("❌ " + message)
}

func NewMessageState(value string) *MessageState {
	return &MessageState{
		lib.NewGenericState(value),
	}
}
