package sql

import "fmt"

type Table[T any] struct {
	Name       string
	Default    T
	col        []string
	primaryKey []string
}

func NewTable[T any](name string, initDefault T) *Table[T] {
	t := Table[T]{
		Name:    name,
		Default: initDefault,
		col:     recursiveGetStructTag(initDefault, "col"),
		// primaryKey: utils.RecursiveGetStructTag(initDefault, "pk"),
	}

	a := findPrimaryKey(initDefault)
	fmt.Println(a)

	return &t
}

func (table Table[T]) GetCol() []string {
	return table.col
}
