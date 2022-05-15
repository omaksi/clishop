package ts

import (
	"context"
	"database/sql"
	"errors"

	"ondrejmaksi.com/db2project/db"
	"ondrejmaksi.com/db2project/rdg"
)

/*

	Používateľ si zvolí, že chce objednať tovar, ktorý má vložený v košíku.
	Zadá miesto dodania.
	Vypočíta sa celková cena (produkty + doprava).
	V systéme sa vytvorí objednávka v stave vytvorená.
	Tovar sa na sklade rezervuje.
	Ak nie je všetok tovar na sklade, objednávka sa stale vytvorí, no používateľ je na to upozornený.

*/
func CreateOrder(userId int, address string) {

	sql := `
		with basket_subtotal as (
			select bi.quantity * p.price as subtotal from basket_items bi
			join products p on bi.product_id = p.id
			where bi.user_id = $1
		),
		basket_total as (
			select sum(bs.subtotal) + 3 as basket_sum from basket_subtotal bs
		),
		new_order as (
			insert into orders (user_id, address, status, total)
			select $1, $2, 'pending', bt.basket_sum
			from basket_total bt
			returning id
		),
		copy_basket_items as (
			insert into order_items (order_id, product_id, quantity)
			select n_o.id, bi.product_id, bi.quantity
			from new_order n_o
			join basket_items bi on bi.user_id = $1
			returning id
		)
		delete from basket_items bi
		where bi.user_id = $1
		;
		`

	_, err := db.GetDatabase().Exec(sql, userId, address)
	if err != nil {
		panic(err)
	}

}

/*
	Objednávky, ktoré ešte neboli zaplatené, sa dajú zrušiť.
	Zarezervovaný tovar sa musí odblokovať.
	Zaplatené ani expedované objednávky nie je možné zrušiť.
*/
func CancelOrder(order rdg.Order) {
	sql := `
	with updated_order as (
		update orders set status = 'cancelled' 
		where id = $1 and status = 'pending' 
		returning id
	),
	updated_order_items as (
		select * 
		from order_items oi 
		where oi.order_id = (select id from updated_order)
	),
	update_products_quantity as (
		update products p set quantity = p.quantity + uoi.quantity
		from updated_order_items uoi
		where p.id = uoi.product_id
		returning uoi.id
	)
	delete from order_items oi 
	where oi.order_id = $1
	;
	`

	_, err := db.GetDatabase().Exec(sql, order.Id)
	if err != nil {
		panic(err)
	}
}

/*
	Zamestnanec vyberie id objednávky, ktorú chce expedovať.
	Ak je objednávka zaplatená a je dostatok tovarov na sklade,
	objednávka sa presunie do stavu expedovaná.
	Zároveň sa upravia skladové zásoby tovarov.
*/
func ExpediteOrder(order rdg.Order) error {

	ctx := context.Background()
	tx, err := db.GetDatabase().BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	defer tx.Rollback()

	if err != nil {
		return errors.New("could not create tx")
	}

	sql := `select status = 'paid' from orders where id = $1`
	var paid bool
	err = tx.QueryRow(sql, order.Id).Scan(&paid)

	if err != nil {
		return errors.New("no such order")
	}

	if !paid {
		return errors.New("order is not paid")
	}

	sql = `
	update orders set status = 'expedited' 
	where id = $1 
	;
	`
	_, err = tx.Exec(sql, order.Id)

	if err != nil {
		return errors.New("could not update order status")
	}

	if err = tx.Commit(); err != nil {
		return errors.New("could not commit tx")
	}

	return nil
}

/*
	Po zaplatení sa v systéme eviduje platba.
	V prípade, že výška platby nesedí so sumou objednávky,
	platba musí byť vrátená (toto simulujte nejakým výpisom do konzoly).
	Po úspešnej platbe sa objednávka presúva do stavu zaplatená.
	V prípade, že na sklade nie je dostatočné množstvo nejakého tovaru,
	tento tovar sa automaticky objedná z veľkoskladu a doplní v
	požadovanom počte (simulujte výpisom do konzoly).
	Dvakrát zaplatiť tú istú objednávku možné nie je.
	Systém to musí odmietnuť a oznámiť chybu.
*/
func PayOrder(order rdg.Order, amount float64) error {
	ctx := context.Background()
	tx, err := db.GetDatabase().BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	defer tx.Rollback()

	if err != nil {
		return errors.New("could not create tx")
	}

	sql := `select status = 'pending' from orders where id = $1`
	var paid bool
	err = tx.QueryRow(sql, order.Id).Scan(&paid)

	if err != nil {
		return errors.New("no such order")
	}

	if !paid {
		return errors.New("order is not pending")
	}

	sql = `select total = $2 from orders where id = $1`
	var goodAmount bool
	err = tx.QueryRow(sql, order.Id, amount).Scan(&goodAmount)

	if err != nil {
		return errors.New("no such order")
	}

	if !goodAmount {
		return errors.New("payment amount does not match order total")
	}

	sql = `
	update orders set status = 'paid' 
	where id = $1 
	;
	`
	_, err = tx.Exec(sql, order.Id)

	if err != nil {
		return errors.New("could not update order status")
	}

	if err = tx.Commit(); err != nil {
		return errors.New("could not commit tx")
	}

	return nil
}
