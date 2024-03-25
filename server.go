package main

import (
	"infoto-web-2/routes"
	"infoto-web-2/db"
	"infoto-web-2/models"
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
