package sql

import (
	"fmt"
	"strings"
)

type Where map[string]interface{}

type SelectParams struct {
	Field    []string
	Where    Where
	WhereCon string
	All      bool
	Offset   int
	Limit    int
}

func (t Table[T]) Select(params SelectParams) ([]T, error) {
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

		conArray := []string{}

		for p, v := range params.Where {
			whereBuilder := strings.Builder{}
			whereBuilder.WriteString(p)
			whereBuilder.WriteString(" = ")
			whereBuilder.WriteString(fmt.Sprint(v))
			conArray = append(conArray, whereBuilder.String())
		}

		logicalConnector := " AND "

		if strings.ToLower(params.WhereCon) == "or" {
			logicalConnector = " OR "
		}

		builder.WriteString(strings.Join(conArray, logicalConnector))
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
	// TODO: update
	return nil, nil
}
