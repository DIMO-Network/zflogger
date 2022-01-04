package main

import (
	"os"

	"github.com/DIMO-Network/zflogger"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// this filter apply logger middleware when the context path is "/error",
// else the zflogger.New are skipped and the flow continue
func filter(c *fiber.Ctx) bool {
	return c.Path() != "/error"
}

func main() {
	log := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.DebugLevel)

	app := fiber.New()

	app.Use(zflogger.New(log, filter))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		a := 0
		return c.JSON(1 / a)
	})

	log.Fatal().Err(app.Listen(":3000")).Str("tag", "server").Send()
}
