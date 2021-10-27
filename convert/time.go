package convert

import (
	"fmt"
	"time"
)

// ToTime interface{} 转换成本地时区的时间，部分字符串在转换时会出现+0000 +0000而不是预期的+0000 UTC，此时需要调用下.UTC()，
// 常见于尾部包含类似-0700或-07:00格式的layout，例如RFC822Z、RFC3339
func ToTime(i interface{}) (time.Time, error) {
	i = indirect(i)

	switch v := i.(type) {
	case time.Time:
		return v, nil
	case int:
		return time.Unix(int64(v), 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	case uint:
		return time.Unix(int64(v), 0), nil
	case uint32:
		return time.Unix(int64(v), 0), nil
	case uint64:
		return time.Unix(int64(v), 0), nil
	case string:
		return parseDate(v)
	case fmt.Stringer:
		return parseDate(v.String())
	case []byte:
		return parseDate(string(v))
	default:
		return time.Time{}, typeError(i, strTime)
	}
}

func parseDate(s string) (time.Time, error) {
	timeFormat := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02T15:04:05-0700", // RFC3339 without timezone hh:mm colon
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05Z0700",  // RFC3339 without T or timezone hh:mm colon
		"2006-01-02 15:04:05",
		"2006.01.02",
		"2006.01.02 15:04:05",
		"2006.01.02 15:04:05 -0700",
		"2006.01.02 15:04:05 MST",
		"2006/01/02",
		"2006/01/02 15:04:05",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05 MST",
		"2006年01月02日",
		"2006年01月02日 15:04:05",
		"2006年01月02日 15时04分05秒",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	for _, format := range timeFormat {
		if t, err := time.Parse(format, s); err == nil {
			return t, err
		}
	}
	return time.Time{}, fmt.Errorf("can't parse date: %s", s)
}
