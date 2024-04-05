package main

import (
	docx "clyeung-hkl/go-test/pkg/docx"
)

func main() {
	// e := domain.Era{
	// 	EraId: 1,
	// 	Name:  "new era",
	// 	Sortable: domain.Sortable{
	// 		Sorting: 2,
	// 	},
	// }

	// fmt.Println("test")
	// fmt.Println(strconv.Itoa(e.Sorting))
	// out, err := json.Marshal(e)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(out))
	// s := sql.Table[domain.Era]{
	// 	Name: "Era",
	// 	Default: domain.Era{
	// 		EraId: 0,
	// 		Name:  "",
	// 		Sortable: domain.Sortable{
	// 			Sorting: 0,
	// 		},
	// 	},
	// }

	// n := sql.NewTable[domain.Era]("Era", domain.Era{
	// 	EraId: 0,
	// 	Name:  "",
	// 	Sortable: domain.Sortable{
	// 		Sorting: 0,
	// 	},
	// })

	// fmt.Println(n.GetCol())
	// fmt.Println(n.GetPrimaryKey())

	// n.Insert(e)

	// a := []domain.Era{}
	// for _ = range 100 {
	// 	a = append(a, e)
	// }

	// fmt.Println(len(a))
	// n.BulkInsert(a)

	// s.Select(sql.SelectParams{
	// 	Field: []string{"test1", "test2"},
	// 	Where: sql.Where{
	// 		"test": 2,
	// 		"foo":  "bar",
	// 	},
	// 	WhereCon: "ORA",
	// 	Offset:   5,
	// 	Limit:    9,
	// })
	// s.Insert(e)

	docx.Test()
}
