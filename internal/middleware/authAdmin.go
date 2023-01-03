package middleware

import (
	"log"

	"Hotelin-BE/internal/database"
	"Hotelin-BE/internal/models"
	"github.com/gofiber/fiber/v2"
)


func AuthAdmin(c Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		header := ctx.GetReqHeaders()

		if _, ok := header["Authorization"]; !ok {
			return c.Unauthorized(ctx)
		}

		userToken := models.UserToken{}
		err := database.DB.Where("token = ?", header["Authorization"]).First(&userToken).Error
		if err != nil {
			return c.Unauthorized(ctx)
		}

		if userToken.Type != "admin" {
			return c.Unauthorized(ctx)
		}

		ctx.Locals("user", userToken)
		log.Println("User Authenticated")
		return ctx.Next()
	}

}