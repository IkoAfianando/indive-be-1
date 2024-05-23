package models

import (
	"time"
)

type User struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Address          string    `json:"address"`
	Phone            string    `json:"phone"`
	Password         string    `json:"-"`
	Verified         bool      `json:"verified"`
	VerificationCode string    `json:"-"`
	CreatedAt        time.Time `json:"created_at"`
	VerifiedAt       time.Time `json:"verified_at"`
}

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserVerify struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}
