package core

import "time"

type User struct {
	Id        int       `json:"-" db:"id"`
	Name      string    `json:"name" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
