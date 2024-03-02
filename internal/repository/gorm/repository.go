package gorm

import (
	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *gormRepository {
	return &gormRepository{
		db: db,
	}
}
