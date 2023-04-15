package user

type ProductRes struct {
	Id         int64
	UserId     int64
	Name       string
	Deskripsi  string
	DetailUser *User
}
