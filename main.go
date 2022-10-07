package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
)

func main() {

	app := fiber.New()

	app.Get("/api/v1/starships/available", func(c *fiber.Ctx) error {

		passengers := c.Query("passengers")

		// invoke
		baseUrl := "https://swapi.dev/api/starships"
		resp, err := http.Get(baseUrl)
		if err != nil {
			fmt.Println("Error read")
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error read")
		}

		// Convert
		var mapData map[string]interface{}
		if err := json.Unmarshal([]byte(body), &mapData); err != nil {
			fmt.Println("Error read")
		}

		nave := 0
		for _, v := range mapData["results"].([]interface{}) {

			p := v.(map[string]interface{})["passengers"]

			if p == passengers {

				nave = nave + 1
			}
		}

		// Return results.
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"number_of_passengers":      passengers,
			"number_of_ships_available": nave,
		})
	})

	app.Listen(":3000")
}
