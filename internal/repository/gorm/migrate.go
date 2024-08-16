package gorm

import (
	migrate "github.com/rubenv/sql-migrate"
)

func (r *gormRepository) Migrate(dir string, direct migrate.MigrationDirection, step int) (int, error) {
	source := migrate.FileMigrationSource{
		Dir: dir,
	}

	rawDB, err := r.db.DB()
	if err != nil {
		return 0, err
	}

	return migrate.ExecMax(rawDB, "mysql", source, direct, step)
}
