package gorm

import (
	"github.com/lw396/WeComCopilot/internal/repository"

	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) repository.Repository {
	return &gormRepository{
		db: db,
	}
}
