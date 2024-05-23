package handlers

import (
	"crypto/rand"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/ikoafianando/be-1.1/database"
	"github.com/ikoafianando/be-1.1/models"
	"github.com/ikoafianando/be-1.1/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data models.UserRegister
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot hash password"})
	}

	data.Password = string(hashedPassword)

	verificationCode := make([]byte, 16)
	_, err = rand.Read(verificationCode)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot generate verification code"})
	}

	verificationCodeStr := utils.EncodeBase64(verificationCode)

	result, err := database.DB.Exec("INSERT INTO users (name, email, address, phone, password, verification_code, created_at, verified) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		data.Name, data.Email, data.Address, data.Phone, data.Password, verificationCodeStr, time.Now(), false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot insert user into database"})
	}

	if err := utils.SendVerificationEmail(data.Email, verificationCodeStr); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot send verification email"})
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot retrieve last insert id"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": lastInsertID, "message": "registration successful, please verify your email"})
}

func VerifyEmail(c *fiber.Ctx) error {
	var request models.UserVerify
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	var user models.User
	err := database.DB.QueryRow("SELECT id, verified FROM users WHERE verification_code = ? and email = ?", request.Code, request.Email).Scan(&user.ID, &user.Verified)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid verification code"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot query user"})
	}

	if user.Verified {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already verified"})
	}

	_, err = database.DB.Exec("UPDATE users SET verified = ?, verified_at = ? WHERE id = ?", true, time.Now(), user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update user"})
	}

	return c.JSON(fiber.Map{"message": "email verified successfully"})
}

func Login(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	var user models.User
	err := database.DB.QueryRow("SELECT id, password, verified FROM users WHERE email = ?", data.Email).Scan(&user.ID, &user.Password, &user.Verified)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot query user"})
	}

	if !user.Verified {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "email not verified"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "incorrect password"})
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}
