package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Request struct {
	Data []string `json:"data"`
}

type Response struct {
	IsSuccess                bool     `json:"issuccess"`
	UserID                   string   `json:"userid"`
	Email                    string   `json:"email_id"`
	Roll_No                  string   `json:"rollno"`
	Numbers                  []string `json:"numbers"`
	Alphabets                []string `json:"alphabets"`
	HighestLowercaseAlphabet string   `json:"highest_lowercase_alphabet"`
}

func processHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.JSON(fiber.Map{
			"operation_code": 1,
		})
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	response := Response{
		IsSuccess:                true,
		UserID:                   "john_doe_17091999",
		Email:                    "john@xyz.com",
		Roll_No:                  "ABCD123",
		Numbers:                  []string{},
		Alphabets:                []string{},
		HighestLowercaseAlphabet: "",
	}

	for _, data := range req.Data {
		if _, err := strconv.Atoi(data); err == nil {
			response.Numbers = append(response.Numbers, data)
		} else if len(data) == 1 && unicode.IsLetter(rune(data[0])) {
			response.Alphabets = append(response.Alphabets, data)
		}
		if len(data) == 1 && unicode.IsLetter(rune(data[0])) && strings.ToLower(data) == data && response.HighestLowercaseAlphabet == "" || strings.Compare(response.HighestLowercaseAlphabet, data) < 0 {
			response.HighestLowercaseAlphabet = data
		}
	}

	return c.JSON(response)
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,OPTIONS",
	}))

	app.Post("/bfhl", processHandler)
	app.Get("/bfhl", processHandler)

	fmt.Println("Server is listening on port 8082")
	app.Listen(":8082")
}
