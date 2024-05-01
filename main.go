package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {

	num, error := sql.Open("postgres", "postgres://postgres:14022014@localhost/demo?sslmode=disable")
	if error != nil {
		log.Fatal("Error connecting to the database:", error)
	}
	error = num.Ping()
	if error != nil {
		log.Fatal("failed to ping ", error)
	}

	return num
}

type product struct {
	product_name  string
	unit          string
	price         float32
	category_name string
	description   string
}

func main() {

	num := Connection()
	defer num.Close()

	sql := `SELECT p.product_name, p.unit, p.price, c.category_name, c.description
	FROM products p
	JOIN categories c ON p.category_id = c.category_id
	WHERE c.category_name = 'Beverages';`

	rows, err := num.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	var products []product
	for rows.Next() {
		var n product
		err = rows.Scan(&n.product_name, &n.unit, &n.price, &n.category_name, &n.description)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, n)
	}

	for _, product := range products {
		fmt.Println(product)
	}
}
