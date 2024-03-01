package repository

type SQLiteSequence struct {
	Name string
	Seq  uint64
}

func (SQLiteSequence) TableName() string {
	return "sqlite_sequence"
}
