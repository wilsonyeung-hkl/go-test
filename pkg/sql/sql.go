package sql

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Table[T any] struct {
	t    T
	Name string
}

func (t Table[T]) Print(a any) any {
	fmt.Printf("sql.Print start with Table: %s \n", t.Name)
	out, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))

	v := reflect.TypeOf(a)

	for i := 0; i < v.NumField(); i++ {
		attr := v.Field(i)
		fmt.Println("---")
		fmt.Println(attr)
		fmt.Println(attr.Name)
		fmt.Println(attr.Type)
		fmt.Println(attr.Anonymous)
		fmt.Println("---")
	}

	return nil
}

// TODO: select statement
