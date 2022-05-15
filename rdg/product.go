package rdg

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"ondrejmaksi.com/db2project/db"
)

/* A product row with Name, Price and Quantity*/
type Product struct {
	Jsonable
	Id       uint
	Name     string
	Price    float64
	Quantity uint
}

func rowsToProducts(rows *sql.Rows) []Product {
	res := []Product{}
	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			panic(err)
		}
		res = append(res, product)
	}

	return res
}

func GetProducts() []Product {
	sql := `SELECT p.id, p.name, p.price, p.quantity FROM products p`
	rows, err := db.GetDatabase().Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rowsToProducts(rows)
}

func CreateProduct(product Product) {
	sql := `
		INSERT INTO products (name, price, quantity) VALUES ($1, $2, $3)
		`
	_, err := db.GetDatabase().Exec(sql, product.Name, product.Price, product.Quantity)
	if err != nil {
		panic(err)
	}

}

func UpdateProduct(product Product) {
	sql := `UPDATE products SET name=$1, price=$2, quantity=$3 WHERE id=$4`
	_, err := db.GetDatabase().Exec(sql, product.Name, product.Price, product.Quantity, product.Id)
	if err != nil {
		log.Println(product)
		panic(err)
	}
}

func ListProductsInBasket(userId int) []Product {
	sql := `SELECT bi.id, p.name, p.price, bi.quantity from basket_items bi JOIN products p ON bi.product_id = p.id WHERE bi.user_id=$1`
	rows, err := db.GetDatabase().Query(sql, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rowsToProducts(rows)
}

func AddProductToBasket(user User, product Product) {
	sql := `INSERT INTO basket_items (user_id, product_id, quantity) VALUES ($1, $2, $3) ON CONFLICT (user_id, product_id) DO NOTHING`
	_, err := db.GetDatabase().Exec(sql, user.Id, product.Id, 1)
	if err != nil {
		panic(err)
	}
}

func UpdateBasketItemQuantity(product Product) {
	log.Println("UpdateBasketItemQuantity: ", product.Id)
	log.Println(product)
	sql := `UPDATE basket_items SET quantity=$1 WHERE id=$2 `
	_, err := db.GetDatabase().Exec(sql, product.Quantity, product.Id)
	if err != nil {
		panic(err)
	}
}

func DeleteBasketItem(product Product) {
	sql := `DELETE FROM basket_items WHERE id=$1`
	_, err := db.GetDatabase().Exec(sql, product.Id)
	if err != nil {
		panic(err)
	}
}

func EmptyBasket(userId int) {
	sql := `DELETE FROM basket_items WHERE user_id=$1`
	_, err := db.GetDatabase().Exec(sql, userId)
	if err != nil {
		panic(err)
	}
}

func DeleteProduct(product Product) {
	sql := `DELETE FROM products WHERE id=$1`
	_, err := db.GetDatabase().Exec(sql, product.Id)
	if err != nil {
		panic(err)
	}
}

/* Fulltext search for product by string input. Searches in all product attributes and product name */
func SearchForProductByAttribute(searchText string) ([]Product, error) {

	ctx := context.Background()
	tx, err := db.GetDatabase().BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	defer tx.Rollback()

	if err != nil {
		return nil, errors.New("could not create tx")
	}

	sql := `
	insert into search_log (query) 
	values ($1) 
	on conflict do nothing
	`
	_, err = tx.Exec(sql, searchText)

	if err != nil {
		return nil, errors.New("could not insert into search_log")
	}

	sql = `with select_log as (
		select id 
		from search_log
		where query = $1
	), log_timestamp as (
		insert into 
			search_log_timestamps (search_log_id, searched_at) 
		select  
			sl.id as search_log_id,
			 NOW() as searched_at
		from 
			select_log sl
	)
	select 
		distinct(p.id), p.name, p.price, p.quantity 
	from 
		products p
	join 
		product_attributes pa on p.id = pa.product_id
	join 
		attribute_values av on av.id = pa.attribute_value_id
	where 
		av.value like '%' || $1 || '%' or p.name like '%' || $1 || '%' 
	;`

	rows, err := tx.Query(sql, searchText)
	if err != nil {
		return nil, errors.New("search failed")
	}
	defer rows.Close()

	return rowsToProducts(rows), nil
}
