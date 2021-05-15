package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // new
	pgconnect "github.com/revell29/gofiber-rest-api/lib"
	"github.com/revell29/gofiber-rest-api/router"
)

func main() {
	app := fiber.New()

	// check db connection
	con, err := pgconnect.Connection()
	if err != nil {
	fmt.Println(err)
	}

	cc, err := con.Acquire()
	if err != nil {
		fmt.Println(err)
	}

	err = cc.Listen(os.Getenv("DB_NAME"))
	if err != nil {
		fmt.Println(err)
	}

	app.Use(logger.New())
  router.InitRoute(app)

	
	app.Listen(":8000")

}