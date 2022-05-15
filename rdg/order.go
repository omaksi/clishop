package rdg

import (
	"database/sql"
	"fmt"

	"ondrejmaksi.com/db2project/db"
)

/*
An order, containing an address, Total price and status.
*/
type Order struct {
	Id      int
	UserId  int
	Address string
	Total   float64
	Status  string
}

/*
An item of an order, containing the product id reference, the order id reference and the quantity.
*/
type OrderItem struct {
	Id        int
	UserId    int
	ProductId int
	OrderId   int
	Quantity  int
}

func rowsToOrders(rows *sql.Rows) []Order {
	res := []Order{}
	for rows.Next() {
		order := Order{}
		err := rows.Scan(&order.Id, &order.UserId, &order.Address, &order.Total, &order.Status)
		if err != nil {
			panic(err)
		}
		res = append(res, order)
	}
	return res
}

func ListOrders() []Order {
	sql := `SELECT * FROM orders`
	rows, err := db.GetDatabase().Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rowsToOrders(rows)
}

func ListOrderItems(orderId int) {
	sql := `SELECT * FROM order_items WHERE order_id=$1`
	rows, err := db.GetDatabase().Query(sql, orderId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		orderItem := OrderItem{}
		err := rows.Scan(&orderItem.Id, &orderItem.UserId, &orderItem.ProductId, &orderItem.OrderId, &orderItem.Quantity)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", orderItem)
	}
}

func getIdFromRows(rows *sql.Rows) int {
	var id int
	err := rows.Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func CreateOrder(order Order) int {
	sql := `INSERT INTO orders (user_id, address, status) VALUES ($1, $2, $3) RETURNING id`
	rows, err := db.GetDatabase().Query(sql, order.UserId, order.Address, order.Status)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return getIdFromRows(rows)
}

func CreateOrderItem(orderItem OrderItem) {
	sql := `INSERT INTO order_items (user_id, product_id, order_id, quantity) VALUES ($1, $2, $3, $4)`
	_, err := db.GetDatabase().Query(sql, orderItem.UserId, orderItem.ProductId, orderItem.OrderId, orderItem.Quantity)
	if err != nil {
		panic(err)
	}
}
