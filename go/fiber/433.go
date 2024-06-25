package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	certFile = "/etc/letsencrypt/live/dyf.newclip.cn/fullchain.pem"
	keyFile  = "/etc/letsencrypt/live/dyf.newclip.cn/privkey.pem"
)

func main433TO80() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Asia/Chongqing",
	}))
	app.Static("/", "./resume.html")
	go http80To443()
	log.Fatal(app.ListenTLS(":443", certFile, keyFile))
}

func http80To443() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Asia/Chongqing",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://dyf.newclip.cn", 301)
	})
	log.Fatal(app.Listen(":80"))
}
