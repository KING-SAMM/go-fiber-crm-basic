package main

import (
	"github.com/KING-SAMM/go-fiber-crm-basic/lead"
	"github.com/KING-SAMM/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
	"fmt"
)

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")

	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection to database opened")
	database.DBConn.AutoMigrate(&lead.Lead())
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	setupRoutes(app)

	app.Listen(3000)
	
	defer database.DBConn.Close()
}