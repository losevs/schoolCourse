package models

type Kid struct {
	ID      uint64 `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint64 `json:"age"`
	Grade   uint64 `json:"grade"`
}

type Teacher struct {
	ID      uint64 `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Grade   uint64 `json:"grade"`
	Subject string `json:"subject"`
}

type Director struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Subject struct {
	Name string `json:"subj"`
}
