package models

type RandNum struct {
	ID      uint `json:"id" gorm:"primary_key"`
	RandNum int  `json:"randNum" gorm:"randNum"`
}
