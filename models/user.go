package models

type User struct {
	FirsName string `json:"first_name"`
	LastName string `json:"last_name"`
	Number   uint   `json:"number"`
}
