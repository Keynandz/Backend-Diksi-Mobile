package models

type Akun struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Image struct {
    Id     int    `json:"imageid"`
    Name   string `json:"name"`
    Mading []byte `json:"-"`
}
