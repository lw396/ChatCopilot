package sqlite

type SQLite struct {
	key  string
	path string
}

func NewSQLiteClient(key, path string) *SQLite {
	return &SQLite{
		key:  key,
		path: path,
	}
}
