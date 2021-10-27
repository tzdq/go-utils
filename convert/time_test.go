package convert

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestToTime(t *testing.T) {
	tm := time.Date(2021, 10, 21, 20, 36, 43, 0, time.Local)
	tmUTC := time.Date(2021, 10, 21, 20, 36, 43, 0, time.UTC)
	tmNoSecond := time.Date(2021, 10, 21, 20, 36, 0, 0, time.Local)
	tmNoSecondUTC := time.Date(2021, 10, 21, 20, 36, 0, 0, time.UTC)
	tmNanoUTC := time.Date(2021, 10, 21, 20, 36, 43, 500000000, time.UTC)
	tmNanoStampUTC := time.Date(0, 10, 21, 20, 36, 43, 500000000, time.UTC)
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		utc     bool
		wantErr bool
	}{
		{name: "int", args: args{1634819803}, want: tm},
		{name: "int32", args: args{int32(1634819803)}, want: tm},
		{name: "int64", args: args{int64(1634819803)}, want: tm},
		{name: "uint", args: args{uint(1634819803)}, want: tm},
		{name: "uint32", args: args{uint32(1634819803)}, want: tm},
		{name: "uint64", args: args{uint64(1634819803)}, want: tm},
		{name: "time.Time", args: args{tm}, want: tm},
		{name: "bytes", args: args{[]byte("2021-10-21 20:36:43 +0800 CST")}, want: tm},
		{name: "fmt.Stringer", args: args{TimeStringer{tm}}, want: tm},
		{name: "time string", args: args{"2021-10-21 20:36:43 +0800 CST"}, want: tm},
		{name: "ansic", args: args{"Thu Oct 21 20:36:43 2021"}, want: tmUTC},
		{name: "UnixDate", args: args{"Thu Oct 21 20:36:43 CST 2021"}, want: tm},
		{name: "RubyDate", args: args{"Thu Oct 21 20:36:43 +0800 2021"}, want: tm},
		{name: "RFC822", args: args{"21 Oct 21 20:36 CST"}, want: tmNoSecond},
		{name: "RFC822Z", args: args{"21 Oct 21 20:36 +0000"}, want: tmNoSecondUTC, utc: true},
		{name: "RFC850", args: args{"Thursday, 21-Oct-21 20:36:43 CST"}, want: tm},
		{name: "RFC1123", args: args{"Thu, 21 Oct 2021 20:36:43 CST"}, want: tm},
		{name: "RFC1123Z", args: args{"Thu, 21 Oct 2021 20:36:43 +0000"}, want: tmUTC, utc: true},
		{name: "Kitchen", args: args{"11:00PM"}, want: time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC)},
		{name: "Stamp", args: args{"Oct 21 20:36:43"}, want: time.Date(0, 10, 21, 20, 36, 43, 0, time.UTC)},
		{name: "StampMilli", args: args{"Oct 21 20:36:43.500"}, want: tmNanoStampUTC},
		{name: "StampMicro", args: args{"Oct 21 20:36:43.500000"}, want: tmNanoStampUTC},
		{name: "StampNano", args: args{"Oct 21 20:36:43.500000000"}, want: tmNanoStampUTC},
		{name: "2021年10月21日", args: args{"2021年10月21日"}, want: time.Date(2021, 10, 21, 0, 0, 0, 0, time.UTC)},
		{name: "2021年10月21日 20:36:43", args: args{"2021年10月21日 20:36:43"}, want: tmUTC},
		{name: "2021年10月21日 20时36分43秒", args: args{"2021年10月21日 20时36分43秒"}, want: tmUTC},
		{name: "2021/10/21", args: args{"2021/10/21"}, want: time.Date(2021, 10, 21, 0, 0, 0, 0, time.UTC)},
		{name: "2021/10/21 20:36:43", args: args{"2021/10/21 20:36:43"}, want: tmUTC},
		{name: "2021/10/21 20:36:43 +0800", args: args{"2021/10/21 20:36:43 +0000"}, want: tmUTC, utc: true},
		{name: "2021/10/21 20:36:43 CST", args: args{"2021/10/21 20:36:43 UTC"}, want: tmUTC},
		{name: "2021.10.21", args: args{"2021.10.21"}, want: time.Date(2021, 10, 21, 0, 0, 0, 0, time.UTC)},
		{name: "2021.10.21 20:36:43", args: args{"2021.10.21 20:36:43"}, want: tmUTC},
		{name: "2021.10.21 20:36:43 +0800", args: args{"2021.10.21 20:36:43 +0800"}, want: tm},
		{name: "2021.10.21 20:36:43 CST", args: args{"2021.10.21 20:36:43 CST"}, want: tm},
		{name: "iso8601", args: args{"2021-10-21T20:36:43"}, want: tmUTC},
		{name: "2021-10-21", args: args{"2021-10-21"}, want: time.Date(2021, 10, 21, 0, 0, 0, 0, time.UTC)},
		{name: "2021-10-21 20:36:43", args: args{"2021-10-21 20:36:43"}, want: tmUTC},
		{name: "21 Oct 2021", args: args{"21 Oct 2021"}, want: time.Date(2021, 10, 21, 0, 0, 0, 0, time.UTC)},
		{name: "RFC3339Nano", args: args{"2021-10-21T20:36:43.5Z"}, want: tmNanoUTC},
		{name: "RFC3339", args: args{"2021-10-21T20:36:43Z"}, want: tmUTC},
		{name: "RFC3339 no timezone", args: args{"2021-10-21T20:36:43+0000"}, want: tmUTC, utc: true},
		{name: "RFC3339 without T or timezone hh:mm colon", args: args{"2021-10-21 20:36:43+0800"}, want: tm},
		{name: "RFC3339 without T", args: args{"2021-10-21 20:36:43+08:00"}, want: tm},
		{name: "RFC3339 with 2021-10-21 20:36:43 -0700", args: args{"2021-10-21 20:36:43Z"}, want: tmUTC},
		{name: "RFC3339 with 2021-10-21 20:36:43 -07:00", args: args{"2021-10-21 20:36:43+08:00"}, want: tm},
		{name: "error", args: args{errors.New("error")}, want: time.Time{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToTime(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.utc {
				got = got.UTC()
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type TimeStringer struct {
	tm time.Time
}

func (ts TimeStringer) String() string {
	return ts.tm.String()
}
