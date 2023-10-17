package fpp

import "gorm.io/gorm"

type FppService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *FppService {
	return &FppService{db: db}
}
