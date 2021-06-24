package crypt

import (
	"strings"
	"testing"
)

func TestRandString(t *testing.T) {
	type args struct {
		length int
		scope  ScopeType
	}
	tests := []struct {
		name        string
		args        args
		compareType int //0 比较范围  1 比较值
		want        string
	}{
		{"LengthInvalid", args{0, OnlyLowerCase}, 1, ""},
		{name: "OnlyLowerCase", args: args{16, OnlyLowerCase}},
		{name: "OnlyUpperCase", args: args{16, OnlyUpperCase}},
		{name: "OnlyNumber", args: args{16, OnlyNumber}},
		{name: "AlphaLetter", args: args{16, AlphaLetter}},
		{name: "LowerCaseAndNumber", args: args{16, LowerCaseAndNumber}},
		{name: "LowerCaseAndSpecial", args: args{16, LowerCaseAndSpecial}},
		{name: "UpperCaseAndNumber", args: args{16, UpperCaseAndNumber}},
		{name: "UpperCaseAndSpecial", args: args{16, UpperCaseAndSpecial}},
		{name: "NumberAndSpecial", args: args{16, NumberAndSpecial}},
		{name: "AlphaAndNumber", args: args{16, AlphaAndNumber}},
		{name: "AlphaAndSpecial", args: args{16, AlphaAndSpecial}},
		{name: "NoUpperCase", args: args{16, NoUpperCase}},
		{name: "NoLowerCase", args: args{16, NoLowerCase}},
		{name: "AllLetter", args: args{16, AllLetter}},
		{name: "Default", args: args{16, 1000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandString(tt.args.length, tt.args.scope)
			if (tt.compareType == 0 && !compareRandString(got, tt.args.scope)) ||
				(tt.compareType == 1 && got != tt.want) {
				t.Errorf("RandString() = %v", got)
			}
		})
	}
}

func compareRandString(got string, scope ScopeType) bool {
	switch scope {
	case OnlyLowerCase:
		return compareOnlyLowerCase(got)
	case OnlyUpperCase:
		return compareOnlyUpperCase(got)
	case OnlyNumber:
		return compareOnlyNumber(got)
	case AlphaLetter:
		return compareAlphaLetter(got)
	case LowerCaseAndNumber:
		return compareLowerCaseAndNumber(got)
	case LowerCaseAndSpecial:
		return compareLowerCaseAndSpecial(got)
	case UpperCaseAndNumber:
		return compareUpperCaseAndNumber(got)
	case UpperCaseAndSpecial:
		return compareUpperCaseAndSpecial(got)
	case NumberAndSpecial:
		return compareNumberAndSpecial(got)
	case AlphaAndNumber:
		return compareAlphaAndNumber(got)
	case AlphaAndSpecial:
		return compareAlphaAndSpecial(got)
	case NoUpperCase:
		return compareNoUpperCase(got)
	case NoLowerCase:
		return compareNoLowerCase(got)
	case AllLetter:
		return compareAllLetter(got)
	default:
		return compareAlphaLetter(got)
	}
}

func compareAllLetter(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareLowerCase(s) && !compareUpperCase(s) && !compareNumber(s) && !compareSpecial(s) {
			return false
		}
	}
	return true
}

func compareNoLowerCase(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareUpperCase(s) && !compareNumber(s) && !compareSpecial(s) {
			return false
		}
	}
	return true
}

func compareNoUpperCase(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareLowerCase(s) && !compareNumber(s) && !compareSpecial(s) {
			return false
		}
	}
	return true
}

func compareAlphaAndSpecial(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareLowerCase(s) && !compareUpperCase(s) && !compareSpecial(s) {
			return false
		}
	}
	return true
}

func compareAlphaAndNumber(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareLowerCase(s) && !compareUpperCase(s) && !compareNumber(s) {
			return false
		}
	}
	return true
}

func compareNumberAndSpecial(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareNumber(s) && !compareSpecial(s) {
			return false
		}
	}
	return true
}

func compareUpperCaseAndSpecial(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareUpperCase(s) && !compareSpecial(s) {
			return false
		}
	}
	return true
}

func compareUpperCaseAndNumber(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareUpperCase(s) && !compareNumber(s) {
			return false
		}
	}
	return true
}

func compareLowerCaseAndSpecial(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareLowerCase(s) && !compareSpecial(s) {
			return false
		}
	}
	return true
}

func compareLowerCaseAndNumber(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareLowerCase(s) && !compareNumber(s) {
			return false
		}
	}
	return true
}

func compareAlphaLetter(got string) bool {
	for i := range got {
		s := string(got[i])
		if !compareLowerCase(s) && !compareUpperCase(s) {
			return false
		}
	}
	return true
}

func compareOnlyNumber(got string) bool {
	for i := range got {
		if !compareNumber(string(got[i])) {
			return false
		}
	}
	return true
}

func compareOnlyLowerCase(got string) bool {
	for i := range got {
		if !compareLowerCase(string(got[i])) {
			return false
		}
	}
	return true
}

func compareOnlyUpperCase(got string) bool {
	for i := range got {
		if !compareUpperCase(string(got[i])) {
			return false
		}
	}
	return true
}

func compareLowerCase(s string) bool {
	return strings.Contains(lowerAlphaCharset, s)
}

func compareUpperCase(s string) bool {
	return strings.Contains(upperAlphaCharset, s)
}

func compareNumber(s string) bool {
	return strings.Contains(numberCharset, s)
}

func compareSpecial(s string) bool {
	return strings.Contains(specialCharset, s)
}

func TestRandBytes(t *testing.T) {
	type args struct {
		length uint32
	}
	tests := []struct {
		name string
		args args
	}{
		{"LengthIsZero", args{0}},
		{"LengthIsZero", args{10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandBytes(tt.args.length)
			valid := true
			for i := range got {
				if got[i] < byte(0x00) || got[i] > byte(0xff) {
					valid = false
				}
			}
			if !valid {
				t.Errorf("RandBytes() = %v", got)
			}
		})
	}
}

func TestRandInt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Normal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandInt(); got < 0 {
				t.Errorf("RandInt() = %v", got)
			}
		})
	}
}

func TestRandIntN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name        string
		args        args
		compareType int //0 比较范围  1 比较值
		want        int
	}{
		{name: "0-100", args: args{100}},
		{"0", args{0}, 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandIntN(tt.args.n)
			if (tt.compareType == 1 && got != tt.want) || (tt.compareType == 0 && (got < 0 || got >= tt.args.n)) {
				t.Errorf("RandIntN() = %v", got)
			}
		})
	}
}

func TestRandIntRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name        string
		args        args
		compareType int //0 比较范围  1 比较值
		wantBegin   int
		wantEnd     int
	}{
		{"5-100", args{5, 100}, 0, 5, 100},
		{"100-5", args{100, 5}, 0, 5, 100},
		{"-1-5", args{-1, 5}, 0, 0, 5},
		{"0-0", args{0, 0}, 1, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandIntRange(tt.args.min, tt.args.max)
			if (tt.compareType == 0 && (got < tt.wantBegin || got >= tt.wantEnd)) ||
				(tt.compareType == 1 && got != tt.wantBegin) {
				t.Errorf("RandIntRange() = %v", got)
			}
		})
	}
}
