package state

import (
	"ondrejmaksi.com/db2project/rdg"
	"ondrejmaksi.com/db2project/ui/lib"
)

type OrdersState struct {
	lib.GenericState[[]rdg.Order]
}

func NewOrdersState(users []rdg.Order) *OrdersState {
	return &OrdersState{}
}

type OrderState struct {
	lib.GenericState[rdg.Order]
}

func NewOrderState(user rdg.Order) *OrderState {
	return &OrderState{}
}
