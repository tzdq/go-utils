package convert

import (
	"fmt"
	"strconv"
)

// ToString interface{}转换为string类型
func ToString(i interface{}) (string, error) {
	i = indirectStringOrError(i)
	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case int:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int32:
		return strconv.FormatInt(int64(s), 10), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(s, 10), nil
	case []byte:
		return string(s), nil
	case []rune:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	case complex64:
		return complexToString(float64(real(s)), float64(imag(s))), nil
	case complex128:
		return complexToString(real(s), imag(s)), nil
	default:
		return "", typeError(i, strString)
	}
}

// complexToString 复数转换为string
func complexToString(real, imag float64) string {
	r := strconv.FormatFloat(real, 'f', -1, 64)
	i := strconv.FormatFloat(imag, 'f', -1, 64)

	// a+bi 如果a等于0 纯虚数  b为0实数 特殊情况a b都为0
	if isEqual(real, 0) {
		if !isEqual(imag, 0) { //纯虚数
			return i + "i"
		}
		return "0" //0+0i 实数
	}
	if isEqual(imag, 0) {
		return r
	}
	return r + "+" + i + "i"
}
