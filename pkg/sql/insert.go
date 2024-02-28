package sql

import (
	"fmt"
	"slices"
	"strings"
)

func (table Table[T]) Insert(t T) error {
	structMap := recursiveTagMap(t, "db")
	s := table.buildInsertStatement(structMap)
	fmt.Println(s)

	return nil
}

func (table Table[T]) BulkInsert(t []T) error {
	if len(t) == 1 {
		return table.Insert(t[0])
	}

	mArr := make([]map[string]interface{}, len(t)-1)
	for i, e := range t {
		if i > 0 {
			mArr[i-1] = recursiveTagMap(e, "db")
		}
	}
	s := table.buildInsertStatement(recursiveTagMap(t[0], "db"), mArr...)
	fmt.Println(s)

	return nil
}

func (table Table[T]) buildInsertStatement(params map[string]interface{}, bulk ...map[string]interface{}) string {
	nArr := []string{}
	pArr := make([][]interface{}, len(bulk)+1)

	tArr := []interface{}{}
	for k, _ := range params {

		if !slices.Contains(table.primaryKey, k) {
			nArr = append(nArr, k)
			tArr = append(tArr, params[k])
		}
	}

	pArr[0] = tArr

	for i, entity := range bulk {
		bArr := []interface{}{}
		for _, n := range nArr {
			bArr = append(bArr, entity[n])
		}
		pArr[i+1] = bArr
	}

	builder := strings.Builder{}
	builder.WriteString("INSERT INTO ")
	builder.WriteString(table.Name)
	builder.WriteString("(")
	builder.WriteString(strings.Join(nArr, ", "))
	builder.WriteString(")")
	builder.WriteString(" VALUES ")

	vArr := make([]string, len(bulk)+1)
	for i, a := range pArr {
		vArr[i] = transformBlanketedString(a)
	}
	builder.WriteString(strings.Join(vArr, ", "))

	return builder.String()
}
