package sqlite

import "gorm.io/gorm"

type SQLite struct {
	key  string
	path string
	db   map[string]*DB
}

type DB struct {
	tx      *gorm.DB
	msgName []string
}

func NewSQLiteClient(key, path string) *SQLite {
	return &SQLite{
		key:  key,
		path: path,
	}
}
