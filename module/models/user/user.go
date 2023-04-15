package user

import (
	"time"
)

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Create_at *time.Time
	Update_at *time.Time
	Delete_at *time.Time
	Product   *Product
	Products  []Product
}

type Product struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	Name       string `json:"name"`
	Deskripsi  string `json:"desc"`
	Create_at  *time.Time
	Update_at  *time.Time
	Delete_at  *time.Time
	DetailUser *User
}
