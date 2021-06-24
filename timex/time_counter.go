package timex

import "time"

// TimeCounter 时间统计类，用于统计时间间隔
type TimeCounter struct {
	int64
}

// NewTimeCounter 初始化TimeCount对象
func NewTimeCounter() (t *TimeCounter) {
	t = new(TimeCounter)
	t.Set()
	return t
}

// Set 开始计时
func (tc *TimeCounter) Set() {
	tc.int64 = time.Now().UnixNano()
}

// Get 返回从开始到现在的时间间隔(s)
func (tc *TimeCounter) Get() int64 {
	return (time.Now().UnixNano() - tc.int64) / int64(time.Second)
}

// GetMs 返回从开始到现在的时间间隔(ms)
func (tc *TimeCounter) GetMs() int64 {
	return (time.Now().UnixNano() - tc.int64) / int64(time.Millisecond)
}

// GetUs 返回从开始到现在的时间间隔(us)
func (tc *TimeCounter) GetUs() int64 {
	return (time.Now().UnixNano() - tc.int64) / int64(time.Microsecond)
}

// GetNs 返回从开始到现在的时间间隔(ns)
func (tc *TimeCounter) GetNs() int64 {
	return time.Now().UnixNano() - tc.int64
}
