package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikoafianando/be-1.1/database"
	"github.com/ikoafianando/be-1.1/models"
	"github.com/ikoafianando/be-1.1/utils"
)

func Dashboard(c *fiber.Ctx) error {
	tokenStr := c.Get("X-API-KEY")
	_, err := utils.ValidateJWT(tokenStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}

	var totalUsers int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE verified = 1").Scan(&totalUsers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot query total users"})
	}

	var users []models.User
	rows, err := database.DB.Query("SELECT id, name, email, created_at, verified_at FROM users WHERE verified = 1")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot query users"})
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		var createdAt, verifiedAt []uint8

		// Scan the results into the user struct
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &createdAt, &verifiedAt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot scan users"})
		}

		// Assign the scanned time values
		user.CreatedAt, err = utils.ParseTime(createdAt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot parse created_at"})
		}
		user.VerifiedAt, err = utils.ParseTime(verifiedAt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot parse verified_at"})
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "rows error"})
	}

	return c.JSON(fiber.Map{"total_users": totalUsers, "users": users})
}
