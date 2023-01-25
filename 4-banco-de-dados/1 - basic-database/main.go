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
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return err
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ? , price = ? where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return err
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ? ")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var p Product

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectProducts(db *sql.DB) ([]Product, error) {

	rows, err := db.Query("select id, name, price  from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, err
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("============================================================")
	fmt.Println("INSERINDO NOVO PRODUTO")
	fmt.Println("============================================================")
	product := NewProduct("Notebook", 1899.0)
	err = insertProduct(db, product)

	if err != nil {
		panic(err)
	}

	fmt.Println("============================================================")
	fmt.Println("PRODUTO INSERIDO COM SUCESSO | ID:", product.ID)
	fmt.Println("============================================================")

	product.Name = "Tablet"
	product.Price = 1000.0

	fmt.Println("============================================================")
	fmt.Println("ATUALIZANDO PRODUTO")
	fmt.Println("============================================================")
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Println("============================================================")
	fmt.Println("PRODUTO ATUALIZADO COM SUCESSO")
	fmt.Println("============================================================")

	fmt.Println("============================================================")
	fmt.Println("SELECIONANDO PRODUTO", product.ID)
	fmt.Println("============================================================")
	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("============================================================")
	fmt.Println("PRODUTO SELECIONANDO COM SUCESSO")
	fmt.Printf("%#v\n", p)
	fmt.Println("============================================================")

	fmt.Println("============================================================")
	fmt.Println("SELECIONANDO TODOS PRODUTOS")
	fmt.Println("============================================================")
	products, err := selectProducts(db)

	if err != nil {
		panic(err)
	}
	fmt.Println("============================================================")
	fmt.Println("PRODUTO SELECIONANDO COM SUCESSO")
	for _, p := range products {
		fmt.Printf("%#v\n", p)
	}
	fmt.Println("============================================================")

	fmt.Println("============================================================")
	fmt.Println("DELETANDO PRODUTO", product.ID)
	fmt.Println("============================================================")
	err = deleteProduct(db, product.ID)

	if err != nil {
		panic(err)
	}
	fmt.Println("============================================================")
	fmt.Println("PRODUTO DELETADO COM SUCESSO")
	fmt.Println("============================================================")

	fmt.Println("============================================================")
	fmt.Println("VERIFICANDO PRODUTO DELETADO")
	fmt.Println("============================================================")
	p, err = selectProduct(db, product.ID)
	if err != nil {
		fmt.Println("============================================================")
		fmt.Println("PRODUTO NÃO EXISTE")
		fmt.Println("============================================================")

		panic(err)
	}

}
