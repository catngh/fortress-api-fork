package model

type Role struct {
	BaseModel

	Name string `json:"name"`
	Code string `json:"code"`
}
