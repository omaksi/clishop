package cmp

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"ondrejmaksi.com/db2project/rdg"
)

type UserTable struct {
	*tview.Table
}

var (
	userTable *UserTable = nil
)

func NewUsersTable(users []rdg.User, onDelete func(user rdg.User), onEdit func(user rdg.User), onBasket func(user rdg.User), onOrder func(user rdg.User), onDone func()) *UserTable {
	if userTable == nil {
		userTable = &UserTable{
			Table: tview.NewTable(),
		}
		userTable.SetBorders(false).
			SetBordersColor(tcell.ColorYellow).
			SetSelectable(true, false)
		userTable.SetFixed(1, 0)
	}

	userTable.SetProps(users, onDelete, onEdit, onBasket, onOrder, onDone)

	return userTable
}

func (ut *UserTable) SetProps(users []rdg.User, onDelete func(user rdg.User), onEdit func(user rdg.User), onBasket func(user rdg.User), onOrder func(user rdg.User), onDone func()) {
	ut.SetSelectedFunc(func(row int, col int) {
		// messageState.SetState(fmt.Sprint(row))
	})
	ut.SetDoneFunc(func(key tcell.Key) {
		onDone()
	})

	ut.SetInputCapture(func(eventKey *tcell.EventKey) *tcell.EventKey {
		if eventKey.Key() == tcell.KeyRune && eventKey.Rune() == 'd' {
			row, _ := ut.GetSelection()
			onDelete(users[row-1])
			return nil
		}
		if eventKey.Key() == tcell.KeyRune && eventKey.Rune() == 'e' {
			row, _ := ut.GetSelection()
			onEdit(users[row-1])
			return nil
		}
		if (eventKey.Key() == tcell.KeyRune && (eventKey.Rune() == 'b')) || eventKey.Key() == tcell.KeyEnter {
			row, _ := ut.GetSelection()
			onBasket(users[row-1])
			return nil
		}
		if eventKey.Key() == tcell.KeyRune && eventKey.Rune() == 'o' {
			row, _ := ut.GetSelection()
			onOrder(users[row-1])
			return nil
		}
		if eventKey.Key() == tcell.KeyEsc {
			onDone()
			return nil
		}
		return eventKey
	})

	ut.ShowUsers(users)
}

func (ut *UserTable) ShowUsers(users []rdg.User) {
	ut.Clear()
	ut.SetFixed(1, 0)

	ut.SetCell(0, 0, tview.NewTableCell("ID").SetTextColor(tcell.ColorYellow).SetSelectable(false))
	ut.SetCell(0, 1, tview.NewTableCell("Email").SetTextColor(tcell.ColorYellow).SetSelectable(false))
	ut.SetCell(0, 2, tview.NewTableCell("BasketCount").SetTextColor(tcell.ColorYellow).SetSelectable(false))
	for i, user := range users {
		ut.SetCell(i+1, 0, tview.NewTableCell(fmt.Sprintf("%d", user.Id)).SetTextColor(tcell.ColorYellow))
		ut.SetCell(i+1, 1, tview.NewTableCell(user.Email))
		ut.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprintf("%d", user.BasketItemCount)))
	}
}
