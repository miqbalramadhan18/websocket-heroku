package main

import (
	"log"

	"github.com/aiteung/musik"
	"github.com/miqbalramadhan18/websocket-herokubal/module"
	"github.com/miqbalramadhan18/websocket-herokubal/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go module.RunHub()

	site := fiber.New()
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
