package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID  int     `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name"`
	Lokasi string  `json:"lokasi"`
	Rating float32 `json:"rating"`
}