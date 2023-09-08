package models

type Kid struct {
	ID      uint64 `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint64 `json:"age"`
	Grade   uint64 `json:"grade"`
}
