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

	/*	db.Create(&Product{
			Name:  "Notebook",
			Price: 1000.0,
		})

		// crate batch
		products := []Product{
			{Name: "PC", Price: 1500.0},
			{Name: "Tablet", Price: 100.0},
			{Name: "XBox", Price: 1500.0},
		}
		db.Create(&products)*/

	// find first
	//var product Product
	//db.Find(&product, 1)
	//fmt.Println(product)
	//db.Find(&product, "name = ?", "XBox")
	//fmt.Println(product)

	//select all
	//var products []Product
	//db.
	//	Limit(2). //limita quantidade
	//	Offset(2). // pagina
	//	Find(&products)
	//
	//for _, p := range products {
	//	fmt.Println(p)
	//}

	//var products []Product
	//db.Where("price > ?", 1000).Find(&products)
	//for _, p := range products {
	//	fmt.Println(p)
	//}
	//
	//db.Where("name LIKE ?", "%book%").Find(&products)
	//for _, p := range products {
	//	fmt.Println(p)
	//}

	//Update
	//var p Product
	//
	//db.First(&p, 1)
	//p.Name = "Celular"
	//db.Save(&p)
	//
	//var p2 Product
	//db.First(&p2, 1)
	//fmt.Println(p2)
	//
	////Delete
	//db.Delete(&p2)
}
