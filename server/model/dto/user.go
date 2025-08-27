package dto

import "time"

type User struct {
	Common   `json:",inline"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Age      *uint8     `json:"age"`
	Birthday *time.Time `json:"birthday"`
}

type CreateUser struct {
	Name     string     `json:"name" binding:"required"`
	Email    string     `json:"email" binding:"required"`
	Age      *uint8     `json:"age"`
	Birthday *time.Time `json:"birthday"`
}

type UpdateUser struct {
	Id       uint       `json:"id" binding:"required"`
	Name     *string    `json:"name"`
	Email    *string    `json:"email"`
	Age      *uint8     `json:"age"`
	Birthday *time.Time `json:"birthday"`
}
