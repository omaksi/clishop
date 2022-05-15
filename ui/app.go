package ui

import (
	"github.com/rivo/tview"

	"ondrejmaksi.com/db2project/db"
	"ondrejmaksi.com/db2project/rdg"
	"ondrejmaksi.com/db2project/ui/cmp"
)

/* Application constructor, creates main application window and components, adds handling for state changes */
func NewApp() *tview.Application {

	var (
		app                         = tview.NewApplication()
		menu                        = cmp.NewMainMenu(onMenuItemSelected)
		content                     = cmp.NewContent()
		statusbar, status, keyhints = cmp.NewStatusBar()
		flex                        = tview.NewFlex()
		log                         = cmp.NewLog()
	)

	s.focus.AddFocusTarget("menu", menu)
	s.focus.AddFocusTarget("content", content)

	s.message.AddHandler(func(newMessage string) {
		status.SetText(newMessage)
		log.AddItem(newMessage, "", 0, nil)
	})

	s.keyHints.AddHandler(func(newKeyHints string) {
		keyhints.SetText(newKeyHints)
	})

	s.title.AddHandler(func(newTitle string) {
		content.SetTitle(" " + newTitle + " ")
	})

	flex.SetDirection(tview.FlexRow).AddItem(
		tview.NewFlex().
			AddItem(menu, 0, 1, true).
			AddItem(content, 0, 3, false).
			AddItem(log, 0, 1, false),
		0, 1, false).
		AddItem(statusbar, 1, 1, false)

	s.content.AddHandler(func(newContent string) {
		onContentChanged(newContent, content)
	})

	app.EnableMouse(true).SetRoot(flex, true)

	return app
}

/* Handles content switching, and sets appropriate components */
func onContentChanged(newContent string, content *cmp.Content) {
	s.keyHints.Clear()

	if newContent == "userList" {
		s.keyHints.SetForUser()
		content.SetContent(
			cmp.NewUsersTable(
				s.users.GetState(),
				a.DeleteUser,
				a.ShowEditUser,
				a.ShowBasket,
				a.ShowCreateOrder,
				s.focus.Menu,
			),
		)
	}

	if newContent == "newUser" {
		content.SetContent(
			cmp.NewUserForm(
				rdg.User{},
				a.CreateUser,
				s.focus.Menu,
			),
		)
	}

	if newContent == "editUser" {
		content.SetContent(
			cmp.NewUserForm(
				s.user.GetState(),
				a.UpdateUser,
				s.focus.Menu,
			),
		)
	}

	if newContent == "search" {
		content.SetContent(
			cmp.NewSearchForm(a.SearchByAttribute, s.focus.Menu),
		)
	}

	if newContent == "searchStatisticsForm" {
		content.SetContent(
			cmp.NewSearchStatisticsForm(a.GetSearchStatistics, s.focus.Menu),
		)
	}

	if newContent == "searchStatistics" {
		content.SetContent(
			cmp.NewSearchStatisticsTable(s.searchStatistics.GetState(), s.focus.Menu),
		)
	}

	if newContent == "basketList" {
		s.keyHints.SetForBasket()
		content.SetContent(
			cmp.NewProductsTable(
				s.products.GetState(),
				a.ShowEditBasketItem,
				a.DeleteBasketItem,
				a.ShowAddBasketItem,
				a.ShowUserList,
			),
		)
	}

	if newContent == "editBasketItem" {
		content.SetContent(
			cmp.NewBasketItemForm(
				s.product.GetState(),
				a.UpdateBasketItem,
				s.focus.Menu,
			),
		)
	}

	if newContent == "addToBasket" {
		s.keyHints.SetForAddToBasket()
		showBasket := func() { a.ShowBasket(s.user.GetState()) }
		content.SetContent(
			cmp.NewProductsTable(
				s.products.GetState(),
				Noop[rdg.Product],
				Noop[rdg.Product],
				a.AddBasketItem,
				showBasket,
			),
		)
	}

	if newContent == "productList" {
		s.keyHints.SetForProduct()
		content.SetContent(
			cmp.NewProductsTable(
				s.products.GetState(),
				a.ShowEditProduct,
				a.DeleteProduct,
				Noop[rdg.Product],
				s.focus.Menu,
			),
		)
	}

	if newContent == "saleStatistics" {
		s.keyHints.Clear()
		content.SetContent(
			cmp.NewSaleStatisticsTable(
				s.saleStatistics.GetState(),
				s.focus.Menu,
			),
		)
	}

	if newContent == "newProduct" {
		s.title.SetState("New Product")
		content.SetContent(
			cmp.NewProductForm(
				rdg.Product{},
				a.CreateProduct,
				s.focus.Menu,
			),
		)
	}

	if newContent == "editProduct" {
		content.SetContent(
			cmp.NewProductForm(
				s.product.GetState(),
				a.UpdateProduct,
				s.focus.Menu,
			),
		)
	}

	if newContent == "orderList" {
		s.keyHints.SetForOrder()
		content.SetContent(
			cmp.NewOrdersTable(
				s.orders.GetState(),
				a.CancelOrder,
				a.ShowPayOrder,
				a.ExpediteOrder,
				s.focus.Menu,
			),
		)
	}

	if newContent == "newOrder" {
		content.SetContent(
			cmp.NewOrderForm(
				s.user.GetState().Id,
				a.CreateOrder,
				s.focus.Menu,
			),
		)
	}

	if newContent == "payOrder" {
		content.SetContent(
			cmp.NewPaymentForm(
				s.order.GetState(),
				a.PayOrder,
				s.focus.Menu,
			),
		)
	}
}

/* Handles menu selection actions */
func onMenuItemSelected(menuId int, menuTitle string) {
	switch menuId {
	case 0:
		a.ShowUserList()
		s.focus.Content()
	case 1:
		a.ShowNewUser()
		s.focus.Content()
	case 2:
		a.ShowProductList()
		s.focus.Content()
	case 3:
		s.content.SetState("newProduct")
		s.focus.Content()
	case 4:
		a.ShowOrderList()
		s.focus.Content()
	case 5:
		s.content.SetState("search")
		s.focus.Content()
	case 6:
		s.content.SetState("searchStatisticsForm")
		s.focus.Content()
	case 7:
		a.ShowSaleStatistics()
		s.focus.Content()
	case 8:
		// --------------------
	case 9:
		s.message.SetMessage("Running Create Script...")
		db.RunScript("sql/create_script.sql")
		s.message.Success("Create Script executed")
	case 10:
		s.message.SetMessage("Running Generate Script...")
		db.RunScript("sql/generate_script.sql")
		s.message.Success("Generate Script executed")
	case 11:
		db.RunScript("sql/create_script.sql")
		s.message.Success("Create Script executed")
		db.RunScript("sql/generate_script.sql")
		s.message.Success("Generate Script executed")
	}
}
