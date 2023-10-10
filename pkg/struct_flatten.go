package pkg

import (
	"errors"
	"reflect"
	"strings"
)

var (
	ErrNotStruct = errors.New("not a struct")
)

type FlattenData struct {
	Key   string
	Value reflect.Value
}

func StructFlatten(obj any, tagName, sperator string) ([]*FlattenData, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	return structFlattenRecursive(v, tagName, sperator, "")
}

func appendKey(prefix, tag, sperator string) string {
	if prefix == "" {
		return tag
	}
	return strings.Join([]string{prefix, tag}, sperator)
}

// TODO: support slice and map
func structFlattenRecursive(v reflect.Value, tagName, sperator, prefix string) ([]*FlattenData, error) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	// iter numFields
	numFields := v.NumField()
	flattenData := make([]*FlattenData, 0, numFields)
	for i := 0; i < numFields; i++ {
		field := v.Field(i)
		// get tag by tagName
		tag := v.Type().Field(i).Tag.Get(tagName)
		if tag == "" {
			// do nothing
			continue
		}
		nextPrefix := appendKey(prefix, tag, sperator)

		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}
		if field.Kind() == reflect.Struct {
			// recursive
			subFlattenData, err := structFlattenRecursive(field, tagName, sperator, nextPrefix)
			if err != nil {
				return nil, err
			}
			flattenData = append(flattenData, subFlattenData...)
		} else {
			flattenData = append(flattenData, &FlattenData{
				Key:   nextPrefix,
				Value: field,
			})
		}
	}

	return flattenData, nil
}

// TODO: implement StructUnflatten
