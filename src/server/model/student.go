package model

type Student struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone" "`
	BirthDay string `json:"birth_day"`
}
