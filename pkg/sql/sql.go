package sql

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Table[T any] struct {
	t    T
	Name string
}

type Where map[string]interface{}

type SelectParams struct {
	Field  []string
	Where  Where
	All    bool
	Offset int
	Limit  int
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

func checkWhereParams(w Where) error {
	if len(w) == 0 {
		return nil
	}

	return nil
}

func (t Table[T]) Select(params SelectParams) ([]T, error) {
	checkWhereParams(params.Where)
	fmt.Println("select start")
	fmt.Println(params.Field)
	fmt.Println(params.Where)
	fmt.Println(params.All)

	builder := strings.Builder{}
	builder.WriteString("SELECT ")
	if params.All || len(params.Field) == 0 {
		builder.WriteString("*")
	} else {
		builder.WriteString(strings.Join(params.Field, ", "))
	}
	builder.WriteString(" FROM `")
	builder.WriteString(t.Name)
	builder.WriteString("`")

	if !params.All && len(params.Where) > 0 {
		builder.WriteString(" WHERE ")

		andArray := []string{}

		for p, v := range params.Where {
			whereBuilder := strings.Builder{}
			whereBuilder.WriteString(p)
			whereBuilder.WriteString(" = ")
			whereBuilder.WriteString(fmt.Sprint(v))
			andArray = append(andArray, whereBuilder.String())
		}
		builder.WriteString(strings.Join(andArray, " AND "))
	}

	if params.Limit > 0 {
		builder.WriteString(" LIMIT ")
		builder.WriteString(fmt.Sprint(params.Limit))
	}

	if params.Offset > 0 {

		builder.WriteString(" OFFSET ")
		builder.WriteString(fmt.Sprint(params.Offset))
	}

	fmt.Println(builder.String())
	return nil, nil
}
