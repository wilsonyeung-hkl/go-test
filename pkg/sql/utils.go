package sql

import (
	"fmt"
	"reflect"
	"strings"
)

// return list of tag value with specific tag
func recursiveGetStructTagValue(s any, tagName string) []string {
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
			arr = append(arr, recursiveGetStructTagValue(v.Field(i).Interface(), tagName)...)
		}
	}

	return arr
}

func recursiveSearchTagValueWithTagPair(s any, searchTag string, searchValue string, targetTag string) []string {
	arr := []string{}

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag, exist := field.Tag.Lookup(searchTag)
		if exist && tag == searchValue {
			colTag, exist := field.Tag.Lookup(targetTag)
			if exist {
				arr = append(arr, colTag)
			}
		}

		if v.Field(i).Kind() == reflect.Struct {
			arr = append(arr, recursiveSearchTagValueWithTagPair(v.Field(i).Interface(), searchTag, searchValue, targetTag)...)
		}
	}

	return arr
}

func recursiveTagMap(s any, tagName string) map[string]interface{} {
	m := make(map[string]interface{})
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag, exist := field.Tag.Lookup(tagName)
		if exist {
			m[tag] = v.Field(i).Interface()
		}
		if v.Field(i).Kind() == reflect.Struct {
			reMap := recursiveTagMap(v.Field(i).Interface(), tagName)

			for key, val := range reMap {
				m[key] = val
			}
		}
	}

	return m
}

func transformBlanketedString(array []interface{}) string {
	builder := strings.Builder{}
	builder.WriteString("(")

	arr := []string{}
	for _, v := range array {
		str := fmt.Sprint(v)
		if reflect.TypeOf(v).Kind() == reflect.String {
			sBuilder := strings.Builder{}
			sBuilder.WriteString("\"")
			sBuilder.WriteString(v.(string))
			sBuilder.WriteString("\"")
			str = sBuilder.String()
		}
		// TODO: handle other type

		arr = append(arr, str)
	}
	builder.WriteString(strings.Join(arr, ", "))

	builder.WriteString(")")

	return builder.String()
}
