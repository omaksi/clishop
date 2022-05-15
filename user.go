package rdg

import (
	"database/sql"
	"log"

	"ondrejmaksi.com/db2project/db"
)

/* User struct containing id, email and basket item count */
type User struct {
	Id              int
	Email           string
	BasketItemCount int
}

func rowsToUsers(rows *sql.Rows) []User {
	res := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Email, &user.BasketItemCount)
		if err != nil {
			panic(err)
		}
		res = append(res, user)
	}

	return res
}

func GetUsers() []User {
	sql := `SELECT u.id, u.email, count(bi.id) FROM users u LEFT JOIN basket_items bi ON u.id = bi.user_id GROUP BY u.id ORDER BY u.id`
	log.Println("GetUsers")
	rows, err := db.GetDatabase().Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rowsToUsers(rows)
}

func GetUser(id int) []User {
	sql := `SELECT * FROM users WHERE id=$1`
	rows, err := db.GetDatabase().Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rowsToUsers(rows)
}

func CreateUser(user User) {
	sql := `INSERT INTO users (email) VALUES ($1)`
	_, err := db.GetDatabase().Query(sql, user.Email)
	if err != nil {
		panic(err)
	}
}

func UpdateUser(user User) {
	sql := `UPDATE users SET email=$1 WHERE id=$2 `
	_, err := db.GetDatabase().Query(sql, user.Email, user.Id)
	if err != nil {
		panic(err)
	}
}

func DeleteUser(user User) {
	sql := `DELETE FROM users WHERE id=$1`
	_, err := db.GetDatabase().Query(sql, user.Id)
	if err != nil {
		panic(err)
	}
}
