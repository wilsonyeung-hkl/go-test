package sql

type Table[T any] struct {
	Name       string
	Default    T
	col        []string
	primaryKey []string
}

func NewTable[T any](name string, initDefault T) *Table[T] {
	t := Table[T]{
		Name:       name,
		Default:    initDefault,
		col:        setCol[T](initDefault),
		primaryKey: setPrimaryKey[T](initDefault),
	}

	return &t
}

func setCol[T any](t T) []string {
	return recursiveGetStructTagValue(t, "db")
}

func setPrimaryKey[T any](t T) []string {
	return recursiveSearchTagValueWithTagPair(t, "key", "primary", "db")
}

func (table Table[T]) GetCol() []string {
	return table.col
}

func (table Table[T]) GetPrimaryKey() []string {
	return table.primaryKey
}
