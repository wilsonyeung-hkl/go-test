package sql

import (
	"fmt"
	"strings"
)

func (t Table[T]) Insert(params T) error {
	fmt.Println("insert start")
	builder := strings.Builder{}
	builder.WriteString("INSERT INTO ")
	builder.WriteString(t.Name)
	fmt.Println(builder.String())

	return nil
}
