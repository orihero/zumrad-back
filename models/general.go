package models

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
