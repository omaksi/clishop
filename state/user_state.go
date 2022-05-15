package state

import (
	"ondrejmaksi.com/db2project/rdg"
	"ondrejmaksi.com/db2project/ui/lib"
)

type UsersState struct {
	lib.GenericState[[]rdg.User]
}

func NewUsersState(users []rdg.User) *UsersState {
	return &UsersState{}
}

type UserState struct {
	lib.GenericState[rdg.User]
}

func NewUserState(user rdg.User) *UserState {
	return &UserState{}
}
