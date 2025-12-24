package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var products = []string{"Keyboard", "Mouse", "Monitor", "Laptop Stand", "USB Cable", "Webcam"}
var regions = []string{"EU", "US", "Asia", "South America", "Africa"}

func main() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pass, name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Generator started...")

	for {
		product := products[rand.Intn(len(products))]
		price := 10 + rand.Float64()*200
		qty := rand.Intn(4) + 1
		region := regions[rand.Intn(len(regions))]

		_, err := db.Exec(`
			INSERT INTO orders (product_name, price, quantity, user_region)
			VALUES ($1, $2, $3, $4)
		`, product, price, qty, region)

		if err != nil {
			fmt.Println("Insert error:", err)
		} else {
			fmt.Println("Inserted:", product, price, qty, region)
		}

		time.Sleep(1 * time.Second)
	}
}
