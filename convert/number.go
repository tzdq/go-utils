package convert

import (
	"fmt"
	"strconv"
)

// ToBool interface{}转换为bool类型
func ToBool(i interface{}) (bool, error) {
	i = indirect(i)

	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		return i.(int) != 0, nil
	case int8:
		return i.(int8) != 0, nil
	case int16:
		return i.(int16) != 0, nil
	case int32:
		return i.(int32) != 0, nil
	case int64:
		return i.(int64) != 0, nil
	case uint:
		return i.(uint) != 0, nil
	case uint8:
		return i.(uint8) != 0, nil
	case uint16:
		return i.(uint16) != 0, nil
	case uint32:
		return i.(uint32) != 0, nil
	case uint64:
		return i.(uint64) != 0, nil
	case string:
		return strconv.ParseBool(i.(string))
	case []byte:
		return strconv.ParseBool(string(i.([]byte)))
	default:
		return false, typeError(i, strBool)
	}
}

// ToInt interface{}转换为int类型
func ToInt(i interface{}) (int, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int64:
		return int(t), nil
	case int32:
		return int(t), nil
	case int16:
		return int(t), nil
	case int8:
		return int(t), nil
	case int:
		return t, nil
	case uint64:
		return int(t), nil
	case uint32:
		return int(t), nil
	case uint16:
		return int(t), nil
	case uint8:
		return int(t), nil
	case uint:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	case string:
		v, err := strconv.ParseInt(t, 0, 0)
		if err != nil {
			return 0, err
		}
		return int(v), nil
	case []byte:
		v, err := strconv.ParseInt(string(t), 0, 0)
		if err != nil {
			return 0, err
		}
		return int(v), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, typeError(i, strInt)
	}
}

// ToInt8 interface{}转换为int8类型
func ToInt8(i interface{}) (int8, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int64:
		return int8(t), nil
	case int32:
		return int8(t), nil
	case int16:
		return int8(t), nil
	case int8:
		return t, nil
	case int:
		return int8(t), nil
	case uint64:
		return int8(t), nil
	case uint32:
		return int8(t), nil
	case uint16:
		return int8(t), nil
	case uint8:
		return int8(t), nil
	case uint:
		return int8(t), nil
	case float32:
		return int8(t), nil
	case float64:
		return int8(t), nil
	case string:
		v, err := strconv.ParseInt(t, 0, 0)
		if err != nil {
			return 0, err
		}
		return int8(v), nil
	case []byte:
		v, err := strconv.ParseInt(string(t), 0, 0)
		if err != nil {
			return 0, err
		}
		return int8(v), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, typeError(i, strInt8)
	}
}

// ToInt16 interface{}转换为int16类型
func ToInt16(i interface{}) (int16, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int64:
		return int16(t), nil
	case int32:
		return int16(t), nil
	case int16:
		return t, nil
	case int8:
		return int16(t), nil
	case int:
		return int16(t), nil
	case uint64:
		return int16(t), nil
	case uint32:
		return int16(t), nil
	case uint16:
		return int16(t), nil
	case uint8:
		return int16(t), nil
	case uint:
		return int16(t), nil
	case float32:
		return int16(t), nil
	case float64:
		return int16(t), nil
	case string:
		v, err := strconv.ParseInt(t, 0, 0)
		if err != nil {
			return 0, err
		}
		return int16(v), nil
	case []byte:
		v, err := strconv.ParseInt(string(t), 0, 0)
		if err != nil {
			return 0, err
		}
		return int16(v), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, typeError(i, strInt16)
	}
}

// ToInt32 interface{}转换为int32类型
func ToInt32(i interface{}) (int32, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int64:
		return int32(t), nil
	case int32:
		return t, nil
	case int16:
		return int32(t), nil
	case int8:
		return int32(t), nil
	case int:
		return int32(t), nil
	case uint64:
		return int32(t), nil
	case uint32:
		return int32(t), nil
	case uint16:
		return int32(t), nil
	case uint8:
		return int32(t), nil
	case uint:
		return int32(t), nil
	case float32:
		return int32(t), nil
	case float64:
		return int32(t), nil
	case string:
		v, err := strconv.ParseInt(t, 0, 0)
		if err != nil {
			return 0, err
		}
		return int32(v), nil
	case []byte:
		v, err := strconv.ParseInt(string(t), 0, 0)
		if err != nil {
			return 0, err
		}
		return int32(v), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, typeError(i, strInt32)
	}
}

// ToInt64 interface{}转换为int64类型
func ToInt64(i interface{}) (int64, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int64:
		return t, nil
	case int32:
		return int64(t), nil
	case int16:
		return int64(t), nil
	case int8:
		return int64(t), nil
	case int:
		return int64(t), nil
	case uint64:
		return int64(t), nil
	case uint32:
		return int64(t), nil
	case uint16:
		return int64(t), nil
	case uint8:
		return int64(t), nil
	case uint:
		return int64(t), nil
	case float32:
		return int64(t), nil
	case float64:
		return int64(t), nil
	case string:
		v, err := strconv.ParseInt(t, 0, 0)
		if err != nil {
			return 0, err
		}
		return v, nil
	case []byte:
		v, err := strconv.ParseInt(string(t), 0, 0)
		if err != nil {
			return 0, err
		}
		return v, nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, typeError(i, strInt64)
	}
}

// ToUint interface{}转换为uint类型
func ToUint(i interface{}) (uint, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int64:
		if t < 0 {
			return 0, typeError(i, strUint)
		}
		return uint(t), nil
	case int32:
		if t < 0 {
			return 0, typeError(i, strUint)
		}
		return uint(t), nil
	case int16:
		if t < 0 {
			return 0, typeError(i, strUint)
		}
		return uint(t), nil
	case int8:
		if t < 0 {
			return 0, typeError(i, strUint)
		}
		return uint(t), nil
	case int:
		if t < 0 {
			return 0, typeError(i, strUint)
		}
		return uint(t), nil
	case uint64:
		return uint(t), nil
	case uint32:
		return uint(t), nil
	case uint16:
		return uint(t), nil
	case uint8:
		return uint(t), nil
	case uint:
		return t, nil
	case float32:
		if t < 0 {
			return 0, typeError(i, strUint)
		}
		return uint(t), nil
	case float64:
		if t < 0 {
			return 0, typeError(i, strUint)
		}
		return uint(t), nil
	case string:
		v, err := strconv.ParseUint(t, 10, 0)
		if err != nil {
			return 0, err
		}
		return uint(v), nil
	case []byte:
		v, err := strconv.ParseUint(string(t), 10, 0)
		if err != nil {
			return 0, err
		}
		return uint(v), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, typeError(i, strUint)
	}
}

// ToUint8 interface{}转换为uint8类型
func ToUint8(i interface{}) (uint8, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int64:
		if t < 0 {
			return 0, typeError(i, strUint8)
		}
		return uint8(t), nil
	case int32:
		if t < 0 {
			return 0, typeError(i, strUint8)
		}
		return uint8(t), nil
	case int16:
		if t < 0 {
			return 0, typeError(i, strUint8)
		}
		return uint8(t), nil
	case int8:
		if t < 0 {
			return 0, typeError(i, strUint8)
		}
		return uint8(t), nil
	case int:
		if t < 0 {
			return 0, typeError(i, strUint8)
		}
		return uint8(t), nil
	case uint64:
		return uint8(t), nil
	case uint32:
		return uint8(t), nil
	case uint16:
		return uint8(t), nil
	case uint8:
		return t, nil
	case uint:
		return uint8(t), nil
	case float32:
		if t < 0 {
			return 0, typeError(i, strUint8)
		}
		return uint8(t), nil
	case float64:
		if t < 0 {
			return 0, typeError(i, strUint8)
		}
		return uint8(t), nil
	case string:
		v, err := strconv.ParseUint(t, 10, 8)
		if err != nil {
			return 0, err
		}
		return uint8(v), nil
	case []byte:
		v, err := strconv.ParseUint(string(t), 10, 8)
		if err != nil {
			return 0, err
		}
		return uint8(v), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, typeError(i, strUint8)
	}
}

// ToUint16 interface{}转换为uint16类型
func ToUint16(i interface{}) (uint16, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int64:
		if t < 0 {
			return 0, typeError(i, strUint16)
		}
		return uint16(t), nil
	case int32:
		if t < 0 {
			return 0, typeError(i, strUint16)
		}
		return uint16(t), nil
	case int16:
		if t < 0 {
			return 0, typeError(i, strUint16)
		}
		return uint16(t), nil
	case int8:
		if t < 0 {
			return 0, typeError(i, strUint16)
		}
		return uint16(t), nil
	case int:
		if t < 0 {
			return 0, typeError(i, strUint16)
		}
		return uint16(t), nil
	case uint64:
		return uint16(t), nil
	case uint32:
		return uint16(t), nil
	case uint16:
		return t, nil
	case uint8:
		return uint16(t), nil
	case uint:
		return uint16(t), nil
	case float32:
		if t < 0 {
			return 0, typeError(i, strUint16)
		}
		return uint16(t), nil
	case float64:
		if t < 0 {
			return 0, typeError(i, strUint16)
		}
		return uint16(t), nil
	case string:
		v, err := strconv.ParseUint(t, 10, 16)
		if err != nil {
			return 0, err
		}
		return uint16(v), nil
	case []byte:
		v, err := strconv.ParseUint(string(t), 10, 16)
		if err != nil {
			return 0, err
		}
		return uint16(v), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, typeError(i, strUint16)
	}
}

// ToUint32 interface{}转换为uint32类型
func ToUint32(i interface{}) (uint32, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int64:
		if t < 0 {
			return 0, typeError(i, strUint32)
		}
		return uint32(t), nil
	case int32:
		if t < 0 {
			return 0, typeError(i, strUint32)
		}
		return uint32(t), nil
	case int16:
		if t < 0 {
			return 0, typeError(i, strUint32)
		}
		return uint32(t), nil
	case int8:
		if t < 0 {
			return 0, typeError(i, strUint32)
		}
		return uint32(t), nil
	case int:
		if t < 0 {
			return 0, typeError(i, strUint32)
		}
		return uint32(t), nil
	case uint64:
		return uint32(t), nil
	case uint32:
		return t, nil
	case uint16:
		return uint32(t), nil
	case uint8:
		return uint32(t), nil
	case uint:
		return uint32(t), nil
	case float32:
		if t < 0 {
			return 0, typeError(i, strUint32)
		}
		return uint32(t), nil
	case float64:
		if t < 0 {
			return 0, typeError(i, strUint32)
		}
		return uint32(t), nil
	case string:
		v, err := strconv.ParseUint(t, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(v), nil
	case []byte:
		v, err := strconv.ParseUint(string(t), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(v), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, typeError(i, strUint32)
	}
}

// ToUint64 interface{}转换为uint64类型
func ToUint64(i interface{}) (uint64, error) {
	i = indirect(i)

	switch t := i.(type) {
	case int64:
		if t < 0 {
			return 0, typeError(i, strUint64)
		}
		return uint64(t), nil
	case int32:
		if t < 0 {
			return 0, typeError(i, strUint64)
		}
		return uint64(t), nil
	case int16:
		if t < 0 {
			return 0, typeError(i, strUint64)
		}
		return uint64(t), nil
	case int8:
		if t < 0 {
			return 0, typeError(i, strUint64)
		}
		return uint64(t), nil
	case int:
		if t < 0 {
			return 0, typeError(i, strUint64)
		}
		return uint64(t), nil
	case uint64:
		return t, nil
	case uint32:
		return uint64(t), nil
	case uint16:
		return uint64(t), nil
	case uint8:
		return uint64(t), nil
	case uint:
		return uint64(t), nil
	case float32:
		if t < 0 {
			return 0, typeError(i, strUint64)
		}
		return uint64(t), nil
	case float64:
		if t < 0 {
			return 0, typeError(i, strUint64)
		}
		return uint64(t), nil
	case string:
		v, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			return 0, err
		}
		return v, nil
	case []byte:
		v, err := strconv.ParseUint(string(t), 10, 64)
		if err != nil {
			return 0, err
		}
		return v, nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, typeError(i, strUint64)
	}
}

// ToFloat64 interface{}转换为float64类型
func ToFloat64(i interface{}) (float64, error) {
	i = indirect(i)

	switch f := i.(type) {
	case float64:
		return f, nil
	case float32:
		//return float64(f), nil //存在精度丢失的可能
		v, err := strconv.ParseFloat(fmt.Sprintf("%f", f), 64)
		if err != nil {
			return 0, err
		}
		return v, nil
	case int:
		return float64(f), nil
	case int8:
		return float64(f), nil
	case int16:
		return float64(f), nil
	case int32:
		return float64(f), nil
	case int64:
		return float64(f), nil
	case uint:
		return float64(f), nil
	case uint8:
		return float64(f), nil
	case uint16:
		return float64(f), nil
	case uint32:
		return float64(f), nil
	case uint64:
		return float64(f), nil
	case string:
		v, err := strconv.ParseFloat(f, 64)
		if err != nil {
			return 0, err
		}
		return v, nil
	case []byte:
		v, err := strconv.ParseFloat(string(f), 64)
		if err != nil {
			return 0, err
		}
		return v, nil
	case bool:
		if f {
			return 1.0, nil
		}
		return 0, nil
	default:
		return 0, typeError(i, strFloat64)
	}
}

// ToFloat32 interface{}转换为float32类型
func ToFloat32(i interface{}) (float32, error) {
	i = indirect(i)

	switch f := i.(type) {
	case float64:
		return float32(f), nil
	case float32:
		return f, nil
	case int:
		return float32(f), nil
	case int8:
		return float32(f), nil
	case int16:
		return float32(f), nil
	case int32:
		return float32(f), nil
	case int64:
		return float32(f), nil
	case uint:
		return float32(f), nil
	case uint8:
		return float32(f), nil
	case uint16:
		return float32(f), nil
	case uint32:
		return float32(f), nil
	case uint64:
		return float32(f), nil
	case string:
		v, err := strconv.ParseFloat(f, 32)
		if err != nil {
			return 0, err
		}
		return float32(v), nil
	case []byte:
		v, err := strconv.ParseFloat(string(f), 32)
		if err != nil {
			return 0, err
		}
		return float32(v), nil
	case bool:
		if f {
			return 1.0, nil
		}
		return 0, nil
	default:
		return 0, typeError(i, strFloat32)
	}
}
