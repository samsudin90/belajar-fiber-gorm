package models

type LoginRequest struct {
	Email    string `jsin:"email" validate:"required,email"`
	Password string `jsin:"password" validate:"required"`
}
