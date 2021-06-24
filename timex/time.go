package timex

import (
	"fmt"
	"github.com/tzdq/go-utils/mathx"
	"strings"
	"time"
)

const (
	SecondsPerDay   = 86400
	HoursPerDay     = 24
	SecondsPerHours = 3600

	// YFormat 年月日格式 有14种格式，这里列举常用的三种
	YFormat          = "2006"
	YMFormat         = "200601"            //年月
	YMDFormat        = "20060102"          //年月日
	YMDFormatWithSep = "2006-01-02"        //年-月-日
	YMDHFormat       = "20060102 03 PM"    //年与日 时
	YMDHMFormat      = "20060102 03:04 PM" //年与日 时:分

	// 时分秒格式
	HMSFormat        = "150405"
	HMFormat         = "1504"
	HMSFormatWithSep = "15:04:05"

	// 时间格式
	DateTimeFormat      = "2006-01-02 15:04:05"
	DateTimeFormatMilli = "2006-01-02 15:04:05.000"
	DateTimeFormatMicro = "2006-01-02 15:04:05.000000"
	DateTimeFormatNano  = "2006-01-02 15:04:05.000000000"
)

const (
	// 比较类型
	IsOneYear = iota
	IsOneMonth
	IsOneWeek
	IsOneDay
	IsOneHour
	IsOneMinute
)

// GetNow 获取当前时间戳(s)
func GetNow() int64 {
	return time.Now().Unix()
}

// GetNowMs 获取当前时间戳(ms)
func GetNowMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetNowUs 获取当前时间戳(us)
func GetNowUs() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// GetNowNs 获取当前时间戳(ns)
func GetNowNs() int64 {
	return time.Now().UnixNano()
}

// GetNowString 获取当前时间
func GetNowString() string {
	return time.Now().Format(DateTimeFormat)
}

// GetNowRFC3339 获取RFC3339格式
func GetNowRFC3339() string {
	return GetNowTime().Format(time.RFC3339)
}

// GetNowTime 获取当前时间(2021-05-06 15:02:24.4718541 +0800 CST)
func GetNowTime() time.Time {
	return time.Now().In(time.Local)
}

// GetNowYMD 获取当前时间的年月日
func GetNowYMD() string {
	return GetNowTime().Format(HMSFormatWithSep)
}

// GetNowHMS 获取当前时间的时分秒
func GetNowHMS() string {
	return GetNowTime().Format(YMDFormatWithSep)
}

// IsLeapYear 判断是否是闰年
func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

// CompareTime 比较2个时间是否属于一定范围
func CompareTime(time1, time2 int64, compareRange int) bool {
	tm1 := time.Unix(time1, 0)
	tm2 := time.Unix(time2, 0)
	switch compareRange {
	case IsOneYear:
		if tm1.Format(YFormat) == tm2.Format(YFormat) {
			return true
		}
	case IsOneMonth:
		if tm1.Format(YMFormat) == tm2.Format(YMFormat) {
			return true
		}
	case IsOneWeek:
		if GetWeekDateStr(tm1, time.Monday, "YYYYMMDD") == GetWeekDateStr(tm2, time.Monday, "YYYYMMDD") {
			return true
		}
	case IsOneDay:
		if tm1.Format(YMDFormat) == tm2.Format(YMDFormat) {
			return true
		}
	case IsOneHour:
		if tm1.Format(YMDHFormat) == tm2.Format(YMDHFormat) {
			return true
		}
	case IsOneMinute:
		if tm1.Format(YMDHMFormat) == tm2.Format(YMDHMFormat) {
			return true
		}
	}
	return false
}

// CalcIntervalDays 计算两个时间戳之间间隔天数，向下取整，begin和end表示时间戳，单位秒
func CalcIntervalDays(begin, end int64) int64 {
	return mathx.AbsInt64(end-begin) / SecondsPerDay
}

// CalcIntervalHours 计算两个时间戳之间的间隔小时数，向下取整
func CalcIntervalHours(begin, end int64) int64 {
	return mathx.AbsInt64(end-begin) / SecondsPerHours
}

// DayBeginTime 获取当天开始时间
func DayBeginTime(t time.Time) time.Time {
	y, m, d := t.Date()
	n := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	return n
}

// DayEndTime 获取当天结束时间
func DayEndTime(t time.Time) time.Time {
	y, m, d := t.Date()
	n := time.Date(y, m, d, 23, 59, 59, 999999999, time.Local)
	return n
}

// DaySecs 获取指定时间是当天的第几秒
func DaySecs(t time.Time) int64 {
	return t.Unix() - DayBeginTime(t).Unix()
}

// Before 当前时间减去多少秒
func Before(sec int64) string {
	return time.Unix(time.Now().Unix()-sec, 0).Format(DateTimeFormat)
}

// After 当前时间加上多少秒
func After(sec int64) string {
	return time.Unix(time.Now().Unix()+sec, 0).Format(DateTimeFormat)
}

// GetWeekDate 获取当周某天,返回t时刻日期
func GetWeekDate(t time.Time, w time.Weekday) time.Time {
	d := w - t.Weekday()
	// 在中国，周日是最后一天，因此针对周日做个修正
	if w == time.Sunday && d != 0 {
		d += 7
	} else if t.Weekday() == time.Sunday && d != 0 {
		d -= 7
	}
	return t.AddDate(0, 0, int(d))
}

// GetWeekDateStr 获取当周某天的年月日字符串格式， 如20180402
func GetWeekDateStr(t time.Time, w time.Weekday, format string) string {
	return ToString(GetWeekDate(t, w), format)
}

// ToTime 2006-01-02 15:04:05字符串格式转换Time格式
func ToTime(layout string) time.Time {
	t, _ := time.ParseInLocation(DateTimeFormat, layout, time.Local)
	return t
}

// ToUnix 2006-01-02 15:04:05字符串格式转换时间戳
func ToUnix(layout string) int64 {
	t, err := time.ParseInLocation(DateTimeFormat, layout, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// ToRFC3339 2006-01-02 15:04:05字符串格式转换RFC3339格式
func ToRFC3339(layout string) string {
	t, err := time.ParseInLocation(DateTimeFormat, layout, time.Local)
	if err != nil {
		return ""
	}
	return t.Format(time.RFC3339)
}

// RFC3339ToUnix 2006-01-02T15:04:05Z07:00字符串格式转换时间戳
func RFC3339ToUnix(layout string) int64 {
	tm, err := time.ParseInLocation(time.RFC3339, layout, time.Local)
	if err != nil {
		return 0
	}
	return tm.Unix()
}

// ToString 格式化time.Time为字符串
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func ToString(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}

// UnixToString 时间戳格式转化为2006-01-02 15:04:05格式字符串
func UnixToString(sec int64) string {
	tm := time.Unix(sec, 0)
	return tm.Format(DateTimeFormat)
}
