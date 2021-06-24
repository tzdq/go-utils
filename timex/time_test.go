package timex

import (
	"reflect"
	"testing"
	"time"
)

func TestDayBeginTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			"2021-04-25", args{time.Unix(1619351526, 0)}, time.Unix(1619280000, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DayBeginTime(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DayBeginTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDayEndTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"2021-04-25", args{time.Unix(1619351526, 0)}, time.Unix(1619366399, 999999999)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DayEndTime(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DayEndTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaySecs(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"2021-04-25", args{time.Unix(1619366399, 0)}, 86399},
		{"2021-04-25", args{time.Unix(1619280000, 0)}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DaySecs(tt.args.t); got != tt.want {
				t.Errorf("DaySecs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcIntervalDays(t *testing.T) {
	type args struct {
		begin int64
		end   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// 2021-04-25 19:52:06~2021-05-01 19:01:20
		{"5Day", args{1619351526, 1619866880}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcIntervalDays(tt.args.begin, tt.args.end); got != tt.want {
				t.Errorf("CalcIntervalDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcIntervalHours(t *testing.T) {
	type args struct {
		begin int64
		end   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// 2021-04-25 19:52:06~2021-04-26 05:51:20
		{"9Hour", args{1619351526, 1619387480}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcIntervalHours(tt.args.begin, tt.args.end); got != tt.want {
				t.Errorf("CalcIntervalHours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWeekDateStr(t *testing.T) {
	type args struct {
		tm     time.Time
		tw     time.Weekday
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"2021-04-25", args{time.Unix(1619351526, 0), time.Friday, "YYYYMMDD"}, "20210423"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeekDateStr(tt.args.tm, tt.args.tw, tt.args.format); got != tt.want {
				t.Errorf("GetWeekDateStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWeekDate(t *testing.T) {
	type args struct {
		t time.Time
		w time.Weekday
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"2021-04-25", args{time.Unix(1619301080, 0), time.Friday}, time.Unix(1619128280, 0)},
		{"2021-03-20", args{time.Unix(1616190680, 0), time.Sunday}, time.Unix(1616277080, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeekDate(tt.args.t, tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWeekDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"LeapYear", args{2000}, true},
		{"OrdinaryYear", args{2001}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeapYear(tt.args.year); got != tt.want {
				t.Errorf("IsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareTime(t *testing.T) {
	type args struct {
		time1        int64
		time2        int64
		compareRange int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"OneDay", args{1619350900, 1619314900, IsOneDay}, true},
		{"DiffDay", args{1650850900, 1619314900, IsOneDay}, false},
		{"OneHour", args{1619314900, 1619312500, IsOneHour}, true},
		{"DiffHour", args{1619314900, 1619348500, IsOneHour}, false},
		{"OneMinute", args{1619348500, 1619348480, IsOneMinute}, true},
		{"DiffMinute", args{1619348500, 1619314900, IsOneMinute}, false},
		{"OneMonth", args{1619348480, 1617274880, IsOneMonth}, true},
		{"DiffMonth", args{1609498880, 1619348480, IsOneMonth}, false},
		{"OneWeek", args{1619351526, 1619262080, IsOneWeek}, true},
		{"DiffWeek", args{1619351526, 1619434880, IsOneWeek}, false},
		{"OneYear", args{1609498880, 1630494080, IsOneYear}, true},
		{"DiffYear", args{1609498880, 1882954880, IsOneYear}, false},
		{"InvalidParam", args{1609498880, 1882954880, 6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareTime(tt.args.time1, tt.args.time2, tt.args.compareRange); got != tt.want {
				t.Errorf("CompareTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToTime(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"2021-04-26", args{"2021-04-26 10:30:25"}, time.Unix(1619404225, 0)},
		{"InvalidParam", args{"2021-04-31 10:30:25"}, time.Time{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTime(tt.args.s); got != tt.want {
				t.Errorf("ToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUnix(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"2021-04-26", args{"2021-04-26 10:30:25"}, 1619404225},
		{"InvalidParam", args{"2021-04-31 10:30:25"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUnix(tt.args.s); got != tt.want {
				t.Errorf("ToUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToRFC3339(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"2021-04-26", args{"2021-04-26 10:30:25"}, "2021-04-26T10:30:25+08:00"},
		{"InvaldParam", args{"2021-04-31 10:30:25"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToRFC3339(tt.args.s); got != tt.want {
				t.Errorf("ToRFC3339() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRFC3339ToUnix(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"2021-04-26", args{"2021-04-26T10:30:25+08:00"}, 1619404225},
		{"InvaldParam", args{"2021-04-31 10:30:25"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RFC3339ToUnix(tt.args.s); got != tt.want {
				t.Errorf("RFC3339ToUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name   string
		want   string
		time   time.Time
		format string
	}{
		{
			name:   "YYYY-MM-DD",
			want:   "2021-05-06",
			format: "YYYY-MM-DD",
			time:   time.Date(2021, 5, 6, 12, 0, 0, 0, time.Local),
		},
		{
			name:   "HH:mm:ss",
			want:   "12:00:00",
			format: "HH:mm:ss",
			time:   time.Date(2021, 5, 6, 12, 0, 0, 0, time.Local),
		},
		{
			name:   "YYYY-MM-DD HH:mm:ss",
			want:   "2021-05-06 12:00:00",
			format: "YYYY-MM-DD HH:mm:ss",
			time:   time.Date(2021, 5, 6, 12, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.time, tt.format); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnixToString(t *testing.T) {
	type args struct {
		sec int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"2021-04-26", args{1619404225}, "2021-04-26 10:30:25"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnixToString(tt.args.sec); got != tt.want {
				t.Errorf("UnixToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
