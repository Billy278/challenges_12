package user

type UserRes struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Product  *Product
	Products []Product
}
