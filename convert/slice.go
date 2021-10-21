package convert

import (
	"reflect"
	"strings"
)

// ToSlice interface{} 转换为[]interface{}
func ToSlice(i interface{}) ([]interface{}, error) {
	if i == nil {
		return nil, typeError(i, strSlice)
	}

	rt, rv := indirectTypeValue(i)
	switch rt.Kind() {
	case reflect.Array, reflect.Slice:
		s := make([]interface{}, rv.Len())
		for j := 0; j < rv.Len(); j++ {
			s[j] = rv.Index(j).Interface()
		}
		return s, nil
	case reflect.String:
		s := make([]interface{}, len(rv.String()))
		for j, v := range rv.String() {
			s[j] = v
		}
		return s, nil
	default:
		return nil, typeError(i, strSlice)
	}
}

// ToIntSlice interface{} 转换为[]int{}
func ToIntSlice(i interface{}) ([]int, error) {
	if i == nil {
		return nil, typeError(i, strIntSlice)
	}

	rt, rv := indirectTypeValue(i)
	switch rt.Kind() {
	case reflect.Array, reflect.Slice:
		s := make([]int, rv.Len())
		for j := 0; j < rv.Len(); j++ {
			v, err := ToInt(rv.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			s[j] = v
		}
		return s, nil
	default:
		s, err := ToInt(rv.Interface())
		if err != nil {
			return nil, err
		}
		return []int{s}, nil
	}
}

// ToStringSlice interface{}转换成[]string
func ToStringSlice(i interface{}) ([]string, error) {
	if i == nil {
		return nil, typeError(i, strStringSlice)
	}

	rt, rv := indirectTypeValue(i)
	switch rt.Kind() {
	case reflect.String:
		return strings.Fields(rv.String()), nil
	case reflect.Array, reflect.Slice:
		s := make([]string, rv.Len())
		for j := 0; j < rv.Len(); j++ {
			v, err := ToString(rv.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			s[j] = v
		}
		return s, nil
	default:
		s, err := ToString(rv.Interface())
		if err != nil {
			return nil, err
		}
		return []string{s}, nil
	}
}
