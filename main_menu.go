package cmp

import (
	"github.com/rivo/tview"
)

func NewMainMenu(onMenuSelect func(int, string)) *tview.List {
	mainMenu := tview.NewList()
	mainMenu.ShowSecondaryText(false).SetTitle("MainMenu").SetBorder(true)
	mainMenu.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		onMenuSelect(i, s1)
	})

	mainMenu.AddItem("Users", "", 0, nil)
	mainMenu.AddItem("New User", "", 0, nil)
	mainMenu.AddItem("Products", "", 0, nil)
	mainMenu.AddItem("New Product", "", 0, nil)
	mainMenu.AddItem("Orders", "", 0, nil)
	mainMenu.AddItem("Search", "", 0, nil)
	mainMenu.AddItem("Search Statistics", "", 0, nil)
	mainMenu.AddItem("Sales", "", 0, nil)

	mainMenu.AddItem("---------------", "", 0, nil)

	mainMenu.AddItem("Create Script", "", 0, nil)
	mainMenu.AddItem("Generate Script", "", 0, nil)
	mainMenu.AddItem("Create & Generate", "", 0, nil)

	return mainMenu
}
