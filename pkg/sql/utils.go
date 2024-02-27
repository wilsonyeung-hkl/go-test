package sql

import "reflect"

// return list of tag value with specific tag
func recursiveGetStructTag(s any, tagName string) []string {
	arr := []string{}

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag, exist := field.Tag.Lookup(tagName)
		if exist {
			arr = append(arr, tag)
		}

		if v.Field(i).Kind() == reflect.Struct {
			arr = append(arr, recursiveGetStructTag(v.Field(i).Interface(), tagName)...)
		}
	}

	return arr
}

func findPrimaryKey(s any) []string {
	arr := []string{}

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag, exist := field.Tag.Lookup("pk")
		if exist && tag == true {
			colTag, exist := field.Tag.Lookup("col")
			if exist {
				arr = append(arr, colTag)
			}
		}

		if v.Field(i).Kind() == reflect.Struct {
			arr = append(arr, recursiveGetStructTag(v.Field(i).Interface(), "pk")...)
		}
	}

	return arr
}
