package ui

import (
	"github.com/rivo/tview"
	"ondrejmaksi.com/db2project/rdg"
	"ondrejmaksi.com/db2project/ui/state"
)

/* Struct holding all of the application UI state */
type AppState struct {
	keyHints         *state.KeyHintsState
	message          *state.MessageState
	focus            *state.FocusState
	content          *state.ContentState
	users            *state.UsersState
	user             *state.UserState
	products         *state.ProductsState
	product          *state.ProductState
	orders           *state.OrdersState
	order            *state.OrderState
	searchStatistics *state.SearchStatisticsState
	saleStatistics   *state.SaleStatisticsState
	title            *state.TitleState
}

var (
	s AppState = AppState{
		keyHints:         state.NewKeyHintsState("Foo"),
		message:          state.NewMessageState("Foo"),
		focus:            state.NewFocusState("menu"),
		content:          state.NewContentState("userList"),
		users:            state.NewUsersState([]rdg.User{}),
		user:             state.NewUserState(rdg.User{}),
		products:         state.NewProductsState([]rdg.Product{}),
		product:          state.NewProductState(rdg.Product{}),
		orders:           state.NewOrdersState([]rdg.Order{}),
		order:            state.NewOrderState(rdg.Order{}),
		searchStatistics: state.NewSearchStatisticsState([]rdg.SearchStatistic{}),
		saleStatistics:   state.NewSaleStatisticsState([]rdg.SaleStatistic{}),
		title:            state.NewTitleState("Welcome!"),
	}
	a Actions = Actions{}
)

/*  Application initialization */
func Init() *tview.Application {
	app := NewApp()
	s.focus.AddHandler(func(newFocus string) {
		// s.message.SetMessage("Focus: " + newFocus)
		app.SetFocus(s.focus.GetPrimitives()[newFocus])
	})

	// db.Subscribe(s.message.SetMessage)

	s.focus.Menu()
	return app
}
