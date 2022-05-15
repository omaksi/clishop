package ui

import (
	"ondrejmaksi.com/db2project/rdg"
	"ondrejmaksi.com/db2project/ts"
)

/* Actions struct, holds actions modifying state, actions beginning with Show change screen content*/
type Actions struct {
}

func Noop[T any](T) {
}

func (c *Actions) ShowUserList() {
	users := rdg.GetUsers()
	s.users.SetState(users)
	s.title.SetState("User List")
	s.content.SetState("userList")
}

func (c *Actions) ShowNewUser() {
	s.title.SetState("New User")
	s.content.SetState("newUser")
}

func (c *Actions) CreateUser(user rdg.User) {
	s.message.SetMessage("Create user: " + user.Email)
	rdg.CreateUser(user)
	a.ShowUserList()
}

func (c *Actions) UpdateUser(user rdg.User) {
	// s.message.SetMessage("Update user: " + user.ToJson())
	rdg.UpdateUser(user)
	a.ShowUserList()
}

func (c *Actions) DeleteUser(user rdg.User) {
	s.message.SetMessage("Delete user: " + user.Email)
	rdg.DeleteUser(user)
	a.ShowUserList()
}

func (c *Actions) ShowEditUser(user rdg.User) {
	s.message.SetMessage("Edit user: " + user.Email)
	s.user.SetState(user)
	s.title.SetState("Edit User " + user.Email)
	s.content.SetState("editUser")
}

func (c *Actions) SearchByAttribute(searchText string) {
	s.message.SetMessage("Search Text: " + searchText)
	products, err := rdg.SearchForProductByAttribute(searchText)
	if err != nil {
		s.message.Fail("Error: " + err.Error())
		return
	}
	s.products.SetState(products)
	s.title.SetState("Search")
	s.content.SetState("productList")
}

func (c *Actions) GetSearchStatistics(searchText string) {
	s.message.SetMessage("GetSearchStatistics Text: " + searchText)
	searchStatistics := rdg.GetSearchStatistics(searchText)
	s.searchStatistics.SetState(searchStatistics)
	s.title.SetState("Search Statistics")
	s.content.SetState("searchStatistics")
}

func (c *Actions) ShowProductList() {
	products := rdg.GetProducts()
	s.products.SetState(products)
	s.title.SetState("Product list")
	s.content.SetState("productList")
}

func (c *Actions) ShowBasket(user rdg.User) {
	s.user.SetState(user)
	products := rdg.ListProductsInBasket(user.Id)
	s.products.SetState(products)
	s.title.SetState("Basket for " + user.Email)
	s.content.SetState("basketList")
}

func (c *Actions) ShowEditProduct(product rdg.Product) {
	s.message.SetMessage("Edit product: " + product.Name)
	s.product.SetState(product)
	s.title.SetState("Edit Product " + product.Name)
	s.content.SetState("editProduct")
}

func (c *Actions) CreateProduct(product rdg.Product) {
	s.message.SetMessage("Create product: " + product.Name)
	rdg.CreateProduct(product)
	a.ShowProductList()
}

func (c *Actions) UpdateProduct(product rdg.Product) {
	s.message.SetMessage("Update product: " + product.ToJson())
	rdg.UpdateProduct(product)
	a.ShowProductList()
}

func (c *Actions) DeleteProduct(product rdg.Product) {
	s.message.SetMessage("Delete product: " + product.Name)
	rdg.DeleteProduct(product)
	a.ShowProductList()
}

func (c *Actions) ShowAddBasketItem(product rdg.Product) {
	products := rdg.GetProducts()
	s.products.SetState(products)
	s.title.SetState("Add item to Basket")
	s.content.SetState("addToBasket")
}

func (c *Actions) AddBasketItem(product rdg.Product) {
	s.message.SetMessage("Add basketItem: " + product.ToJson())
	rdg.AddProductToBasket(s.user.GetState(), product)
	a.ShowBasket(s.user.GetState())
}

func (c *Actions) ShowEditBasketItem(product rdg.Product) {
	s.message.SetMessage("Edit basketItem: " + product.Name)
	s.product.SetState(product)
	s.title.SetState("Edit basketItem " + product.Name)
	s.content.SetState("editBasketItem")
}

func (c *Actions) UpdateBasketItem(product rdg.Product) {
	s.message.SetMessage("Update basketItem: " + product.ToJson())
	rdg.UpdateBasketItemQuantity(product)
	a.ShowBasket(s.user.GetState())
}

func (c *Actions) DeleteBasketItem(product rdg.Product) {
	s.message.SetMessage("Delete basketItem: " + product.Name)
	rdg.DeleteBasketItem(product)
	a.ShowBasket(s.user.GetState())
}

func (c *Actions) CreateOrder(userId int, address string) {
	s.message.SetMessage("Create Order: " + address)
	ts.CreateOrder(userId, address)
	a.ShowOrderList()
}

func (c *Actions) ShowOrderList() {
	orders := rdg.ListOrders()
	s.orders.SetState(orders)
	s.title.SetState("Order List")
	s.content.SetState("orderList")
}

func (c *Actions) ShowCreateOrder(user rdg.User) {
	s.message.SetMessage("ShowCreateOrder: ")
	s.user.SetState(user)
	s.title.SetState("New Order")
	s.content.SetState("newOrder")
}

func (c *Actions) CancelOrder(order rdg.Order) {
	s.message.SetMessage("CancelOrder: ")
	ts.CancelOrder(order)
	a.ShowOrderList()
}

func (c *Actions) ExpediteOrder(order rdg.Order) {
	err := ts.ExpediteOrder(order)
	if err != nil {
		s.message.SetMessage("ExpediteOrder: " + err.Error())
	}
	a.ShowOrderList()
}

func (c *Actions) ShowPayOrder(order rdg.Order) {
	s.message.SetMessage("ShowPayOrder: ")
	s.order.SetState(order)
	s.title.SetState("Pay Order")
	s.content.SetState("payOrder")
}

func (c *Actions) PayOrder(order rdg.Order, amount float64) {
	s.message.SetMessage("PayOrder: ")
	err := ts.PayOrder(order, amount)
	if err != nil {
		s.message.SetMessage("PayOrder: " + err.Error())
	}
	a.ShowOrderList()
}

func (c *Actions) ShowSaleStatistics() {
	s.message.SetMessage("ShowOrder: ")
	saleStatistics := rdg.GetSaleStatistics()
	s.saleStatistics.SetState(saleStatistics)
	s.title.SetState("Sale Statistics")
	s.content.SetState("saleStatistics")
}
