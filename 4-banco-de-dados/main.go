package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		panic(err)
	}
	return err
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ? , price = ? where id = ?")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		panic(err)
	}
	return err
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ? ")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	var p Product

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		panic(err)
	}
	return &p, nil
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	product := NewProduct("Notebook", 1899.0)
	err = insertProduct(db, product)

	if err != nil {
		panic(err)
	}

	product.Name = "Tablet"
	product.Price = 1000.0

	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", p)
}
