package cmp

import (
	"github.com/rivo/tview"
	"ondrejmaksi.com/db2project/rdg"
)

func NewUserForm(user rdg.User, onSubmit func(user rdg.User), onCancel func()) *tview.Form {
	email := user.Email
	userForm := tview.NewForm().
		AddInputField("Email", user.Email, 20, nil, func(text string) {
			email = text
			// messageState.SetState(email)
		}).
		AddButton("Save", func() {
			newUser := user
			newUser.Email = email
			onSubmit(newUser)
		}).
		SetCancelFunc(func() {
			onCancel()
		})

	return userForm
}
