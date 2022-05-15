package state

import (
	"ondrejmaksi.com/db2project/rdg"
	"ondrejmaksi.com/db2project/ui/lib"
)

type ProductsState struct {
	lib.GenericState[[]rdg.Product]
}

func NewProductsState(users []rdg.Product) *ProductsState {
	return &ProductsState{}
}

type ProductState struct {
	lib.GenericState[rdg.Product]
}

func NewProductState(user rdg.Product) *ProductState {
	return &ProductState{}
}
