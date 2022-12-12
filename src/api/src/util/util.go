package util

import (
	"os"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

type RandNum struct {
	ID      uint `json:"id" gorm:"primary_key"`
	RandNum int  `json:"randNum" gorm:"randNum"`
}
