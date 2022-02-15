package models

type Monster struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	Classification string `json:"classification"`
}
