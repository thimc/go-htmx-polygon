package main

import (
	"github.com/gofiber/fiber/v2"
)

func HandleIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func HandleSearch(c *fiber.Ctx) error {
	ticker := c.Query("ticker")
	results, err := SearchTicker(ticker)
	if err != nil {
		return err
	}

	return c.Render("results", fiber.Map{
		"Ticker":  ticker,
		"Results": results,
	})
}

func HandleValues(c *fiber.Ctx) error {
	ticker := c.Params("ticker")
	values, err := GetDailyValues(ticker)
	if err != nil {
		return err
	}

	return c.Render("values", fiber.Map{
		"Ticker": ticker,
		"Values": values,
	})
}
