package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `grom:"type:varchar(20);not null " json:"username"`
	Password string `grom:"type:varchar(20);not null " json:"password"`
	Role     int    `grom:"type:int;not null " json:"role"`
}
