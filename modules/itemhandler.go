package itemhandler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	pgconnect "github.com/revell29/gofiber-rest-api/lib"
)

type repoItem struct {
	ItemId int
	ItemName string
	ItemCode string
	Barcode string
	AvailableStock int
}

type itemRespository struct {
	Items []repoItem
}


func GetAllItems(c *fiber.Ctx) error {
	conn, err := pgconnect.Connection()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := conn.Query("select item_id, item_name, item_code, barcode, available_stock from items")
	repos := itemRespository{}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		item := repoItem{}

		err := rows.Scan(&item.ItemId, &item.ItemName, &item.ItemCode, &item.Barcode, &item.AvailableStock)

		if err != nil {
			return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
				"status": 500,
				"message": "Api call success",
			})
		}
		repos.Items = append(repos.Items, item)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": 200,
		"message": "ok",
		"data": repos,
	})

}

func DetailItems(c *fiber.Ctx) error {
	conn, err := pgconnect.Connection()

	if err != nil {
		log.Println(err)
	}

	id := c.Params("id")
	rows, err := conn.Query("select item_id, item_name, item_code, barcode, available_stock from items where item_id = $1", id)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"status": 500,
			"message": err,
		})
	}

	defer rows.Close()
	item := repoItem{}
	for  rows.Next() {
		err := rows.Scan(&item.ItemId, &item.ItemName, &item.ItemCode, &item.Barcode, &item.AvailableStock)

		log.Println(err)
		if err == nil {
			log.Println("No rows were returned!")
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": 200,
		"message": "ok",
		"data": item,
	})
}