package convert

import (
	"fmt"
	"math"
	"reflect"
)

// indirect 如果i是指针，返回多次解引用后的基础类型，反之直接返回i或nil
func indirect(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	if t := reflect.TypeOf(i); t.Kind() != reflect.Ptr {
		return i
	}
	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// indirectTypeValue 如果i是指针，返回多次解引用后的reflect.Type和reflect.Value，反之直接返回
func indirectTypeValue(i interface{}) (rt reflect.Type, rv reflect.Value) {
	if i == nil {
		return
	}
	rt = reflect.TypeOf(i)
	rv = reflect.ValueOf(i)
	if rt.Kind() != reflect.Ptr {
		return
	}
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return
		}
		rv = rv.Elem()
	}
	return indirectTypeValue(rv.Interface())
}

// indirectStringOrError 如果i是指针，返回多次解引用后的基础类型或fmt.Stringer或错误实现
func indirectStringOrError(i interface{}) interface{} {
	if i == nil {
		return nil
	}

	errType := reflect.TypeOf((*error)(nil)).Elem()
	fmtStringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(i)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errType) && v.Kind() == reflect.Ptr &&
		!v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func typeError(i interface{}, t string) error {
	return fmt.Errorf("%#v of type %T can't convert to %v", i, i, t)
}

// typeError 函数中t参数
const (
	strUint        = "uint"
	strUint8       = "uint8"
	strUint16      = "uint16"
	strUint32      = "uint32"
	strUint64      = "uint64"
	strInt         = "int"
	strInt8        = "int8"
	strInt16       = "int16"
	strInt32       = "int32"
	strInt64       = "int64"
	strFloat32     = "float32"
	strFloat64     = "float64"
	strBool        = "bool"
	strString      = "string"
	strSlice       = "[]interface{}"
	strStringSlice = "[]string"
	strIntSlice    = "[]int"
)

const epsilon = 1e-16

// isEqual 判断浮点数a与b是否相等
func isEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
