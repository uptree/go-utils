package datetime

import "time"

const DefaultLayout string = "2006-01-02 15:04:05"
const DayLayout string = "2006-01-02"

// Time 获取当前时间戳
func Time() int64 {
	return time.Now().Unix()
}

// MilliTime 获取当前毫秒时间戳
func MilliTime() int64 {
	return time.Now().UnixNano() / 1e6
}

// MicroTime 获取当前微秒时间戳
func MicroTime() int64 {
	return time.Now().UnixNano() / 1e3
}

// Date 时间戳格式化
func Date(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}

// Timestamp 时间转时间戳
func Timestamp(datetime string, layout string) int64 {
	tm2, err := time.ParseInLocation(layout, datetime, time.Local)
	if err != nil {
		return 0
	}
	return tm2.Unix()
}

// Datetime 获取当前时间
func Datetime() string {
	return Date(Time(), DefaultLayout)
}

// Today 获取今天日期
func Today() string {
	return Date(Time(), DayLayout)
}

// TodayStartTime  Today => 00:00:00 获取当天起始时间戳
func TodayStartTime() int64 {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	return tm1
}

// TodayEndTime Today => 23:59:59 获取当天结束时间戳
func TodayEndTime() int64 {
	tm1 := TodayStartTime() + 86400 - 1
	return tm1
}
