package models

type User struct {
	ID           int64   `gorm:"primaryKey"`
	FirstName    *string `json: "firstName" validate:"required, min=2, max=100"`
	LastName     *string `json: "lastName" validate:"required, min=2, max=100"`
	Password     *string `json: "password" validate:"required, min=2, max=100"`
	Email        *string `json: "email" validate:"email, required"`
	Token        *string `json: "token"`
	UserType     *string `json: "userType" validate:"required, eq=ADMIN|eq=USER"`
	RefreshToken *string `json: "refreshToken"`
}
