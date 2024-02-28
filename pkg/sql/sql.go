package sql

type Table[T any] struct {
	Name       string
	Default    T
	col        []string
	primaryKey []string
}

const colStructTag = "db"

func NewTable[T any](name string, initDefault T) *Table[T] {
	t := Table[T]{
		Name:       name,
		Default:    initDefault,
		col:        setCol(initDefault),
		primaryKey: setPrimaryKey(initDefault),
	}

	return &t
}

func setCol(t any) []string {
	return recursiveGetStructTagValue(t, colStructTag)
}

func setPrimaryKey(t any) []string {
	return recursiveSearchTagValueWithTagPair(t, "key", "primary", colStructTag)
}

func (table Table[T]) GetCol() []string {
	return table.col
}

func (table Table[T]) GetPrimaryKey() []string {
	return table.primaryKey
}
