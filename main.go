package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikoafianando/be-1.1/database"
	"github.com/ikoafianando/be-1.1/handlers"
)

func main() {
	app := fiber.New()

	database.Init()

	app.Post("/register", handlers.Register)
	app.Post("/verify/email", handlers.VerifyEmail)
	app.Post("/login", handlers.Login)
	app.Get("/dashboard", handlers.Dashboard)

	app.Get("/auth/google/login", handlers.OauthGoogleLogin)
	app.Get("/auth/google/callback", handlers.OauthGoogleCallback)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
