package main

import (
	"errors"
	"fmt"
	"strings"
)

type Product struct {
	ID       int
	Name     string
	Brand    string
	Price    float64
	Stock    int
	Category string
	Active   bool
}

type Catalog struct {
	products []Product
}

func (catalog *Catalog) AddProduct(newProduct Product) error {
	for _, product := range catalog.products {
		if product.ID == newProduct.ID {
			return errors.New("A product with this ID already exists")
		}
	}
	catalog.products = append(catalog.products, newProduct)
	return nil
}

func (catalog *Catalog) FindByID(id int) (*Product, error) {
	for _, product := range catalog.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, errors.New("Product not found")
}

func (catalog Catalog) FindByCategory(category string) []Product {
	var results []Product
	for _, product := range catalog.products {
		if strings.EqualFold(product.Category, category) {
			results = append(results, product)
		}
	}
	return results
}

func (catalog *Catalog) ApplyDiscount(category string, percentage float64) int {
	modified := 0
	for i, product := range catalog.products {
		if strings.EqualFold(product.Category, category) {
			catalog.products[i].Price = product.Price * (1 - percentage/100)
			modified++
		}
	}
	return modified
}

func (catalog *Catalog) Sell(id int, quantity int) error {
	if quantity <= 0 {
		return errors.New("Quantity must be positive")
	}
	p, err := catalog.FindByID(id)
	if err != nil {
		return err
	}
	if p.Stock < quantity {
		return errors.New("Insufficient stock")
	}
	p.Stock -= quantity
	return nil
}

func (catalog Catalog) Report() string {
	totalValue := 0.0
	totalStock := 0
	for _, product := range catalog.products {
		totalStock += product.Stock
		totalValue += float64(product.Stock) * product.Price
	}
	return fmt.Sprintf(
		"Report — %d product(s) with %d articles in stock, total stock value: $%.2f",
		len(catalog.products),
		totalStock,
		totalValue,
	)
}

func displayProduct(p Product) {
	status := "Inactive"
	if p.Active {
		status = "Active"
	}
	fmt.Printf(
		"  #%d — %s (%s) | $%.2f | stock: %d | %s | %s\n",
		p.ID, p.Name, p.Brand, p.Price, p.Stock, p.Category, status,
	)
}

func menuAdd(c *Catalog) {
	var p Product
	var err error

	fmt.Print("ID: ")
	_, err = fmt.Scan(&p.ID)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Print("Name: ")
	_, err = fmt.Scan(&p.Name)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Print("Brand: ")
	_, err = fmt.Scan(&p.Brand)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Print("Price: ")
	_, err = fmt.Scan(&p.Price)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Print("Stock: ")
	_, err = fmt.Scan(&p.Stock)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Print("Category: ")
	_, err = fmt.Scan(&p.Category)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	p.Active = true
	err = c.AddProduct(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Product added successfully")
}

func menuSearch(c *Catalog) {
	var choice int
	var err error

	fmt.Println("[1] By ID")
	fmt.Println("[2] By category")
	fmt.Print("Choice: ")
	_, err = fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	switch choice {
	case 1:
		var id int
		var p *Product

		fmt.Print("ID: ")
		_, err = fmt.Scan(&id)
		if err != nil {
			fmt.Println("Input error:", err)
			return
		}

		p, err = c.FindByID(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		displayProduct(*p)
	case 2:
		var cat string

		fmt.Print("Category: ")
		_, err = fmt.Scan(&cat)
		if err != nil {
			fmt.Println("Input error:", err)
			return
		}
		products := c.FindByCategory(cat)
		if len(products) == 0 {
			fmt.Println("No products in this category")
			return
		}
		for _, p := range products {
			displayProduct(p)
		}
	default:
		fmt.Println("Invalid choice")
	}
}

func menuSales(c *Catalog) {
	var category string
	var percentage float64
	var err error

	fmt.Print("Category: ")
	_, err = fmt.Scan(&category)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Print("Discount (%): ")
	_, err = fmt.Scan(&percentage)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	if percentage < 0 || percentage > 100 {
		fmt.Println("Error: percentage must be between 0 and 100")
		return
	}
	n := c.ApplyDiscount(category, percentage)
	fmt.Printf("%d product(s) discounted\n", n)
}

func menuSell(c *Catalog) {
	var id, qty int
	var err error

	fmt.Print("Product ID: ")
	_, err = fmt.Scan(&id)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Print("Quantity: ")
	_, err = fmt.Scan(&qty)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	err = c.Sell(id, qty)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Sale recorded")
}

func main() {
	catalog := Catalog{}
	// Thanks the IA for the data
	initialProducts := []Product{
		{ID: 1, Name: "iPhone 15 Pro", Brand: "Apple", Price: 1199.00, Stock: 25, Category: "Smartphone", Active: true},
		{ID: 2, Name: "MacBook Air M3", Brand: "Apple", Price: 1299.00, Stock: 12, Category: "Computer", Active: true},
		{ID: 3, Name: "Galaxy S24 Ultra", Brand: "Samsung", Price: 1099.00, Stock: 18, Category: "Smartphone", Active: true},
		{ID: 4, Name: "PlayStation 5", Brand: "Sony", Price: 499.00, Stock: 8, Category: "Console", Active: true},
		{ID: 5, Name: "AirPods Pro 2", Brand: "Apple", Price: 279.00, Stock: 40, Category: "Audio", Active: true},
	}

	for _, product := range initialProducts {
		err := catalog.AddProduct(product)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Catalog initialized")

	for {
		var choice int
		var err error

		fmt.Println("=== Menu ===")
		fmt.Println("[1] Add")
		fmt.Println("[2] Search")
		fmt.Println("[3] Sales")
		fmt.Println("[4] Sell")
		fmt.Println("[5] Report")
		fmt.Println("[0] Quit")
		fmt.Print("Choice: ")
		_, err = fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		switch choice {
		case 1:
			menuAdd(&catalog)
		case 2:
			menuSearch(&catalog)
		case 3:
			menuSales(&catalog)
		case 4:
			menuSell(&catalog)
		case 5:
			fmt.Println(catalog.Report())
		case 0:
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
