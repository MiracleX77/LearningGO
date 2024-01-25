package main

import (
	"log"

	_ "github.com/lib/pq"
)

func createProduct(product *Product) {
	_, err := db.Exec("INSERT INTO public.products (name, price) VALUES ($1, $2);", product.Name, product.Price)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow("SELECT * FROM products WHERE id = $1;", id)
	err := row.Scan(&p.Id, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func getProducts() ([]Product, error) {
	rows, err := db.Query("SELECT * FROM products;")
	if err != nil {
		return nil, err
	}
	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
func updateProduct(id int, product *Product) error {
	_, err := db.Exec(
		"UPDATE products SET name = $1, price = $2 WHERE id = $3;",
		product.Name, product.Price, id)
	if err != nil {
		return err
	}
	return nil
}
