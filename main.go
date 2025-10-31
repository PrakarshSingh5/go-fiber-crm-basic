package main

import (
		"github/PrakarshSingh5/go-fiber-crm-basic/database"
		"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}
func initDatabase(){
	var err error
	database.DBConn
}
func main(){
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}
