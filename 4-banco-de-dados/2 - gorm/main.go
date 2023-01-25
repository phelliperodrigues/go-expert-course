package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	//create

	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.0,
	})

	// crate batch
	products := []Product{
		{Name: "PC", Price: 1500.0},
		{Name: "Tablet", Price: 100.0},
		{Name: "XBox", Price: 1500.0},
	}
	db.Create(&products)

}
