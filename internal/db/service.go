package db

import "gorm.io/gorm"

type Service interface {
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
