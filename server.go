package main

import (
	"htmx-test/routes"
	"htmx-test/db"
	"htmx-test/models"
)

func main() {
    data := db.NewDB()
    item := &models.Item{
        Id: 0,
        Name: "Milk",
        Quantity: 1,
    }
    data.Add(item)
    e := routes.Setup(data)
    e.Logger.Fatal(e.Start(":3000"))
}
