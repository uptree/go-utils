package datetime

import (
	"time"
)

const (
	DefaultLayout  = "2006-01-02 15:04:05"
	DayLayout      = "2006-01-02"
	YYYYMMDDLayout = "20060102"
)

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

// Datetime 获取当前时间
func Datetime() string {
	return Date(Time(), DefaultLayout)
}

// Today 获取今天日期
func Today() string {
	return Date(Time(), DayLayout)
}

// ToTime 时间戳转时间
func ToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// StringToTime 字符串转时间
func StringToTime(datetime string, layout string) time.Time {
	tm2, err := time.ParseInLocation(layout, datetime, time.Local)
	if err != nil {
		return time.Time{}
	}
	return tm2
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

// TodayStartTime Today => 00:00:00 获取今天起始时间戳
func TodayStartTime() time.Time {
	return DayStartTime(time.Now())
}

// TodayEndTime Today => 23:59:59 获取今天结束时间戳
func TodayEndTime() time.Time {
	return DayEndTime(time.Now())
}

// DayStartTime Day => 00:00:00 获取当天起始时间戳
func DayStartTime(t time.Time) time.Time {
	tm2 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return tm2
}

// DayEndTime Day => 23:59:59 获取当天结束时间戳
func DayEndTime(t time.Time) time.Time {
	tm2 := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local)
	return tm2
}

// WeekStartTime Monday => 00:00:00 获取周起始时间戳
func WeekStartTime(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	tm2 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).
		AddDate(0, 0, offset)
	return tm2
}

// WeekEndTime Sunday => 23:59:59 获取周结束时间戳
func WeekEndTime(t time.Time) time.Time {
	offset := 0
	if w := t.Weekday(); w != 0 {
		offset = int(time.Saturday + 1 - w)
	}
	tm2 := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local).
		AddDate(0, 0, offset)
	return tm2
}

// MonthStartTime 月初 => 00:00:00 获取月起始时间戳
func MonthStartTime(t time.Time) time.Time {
	tm2 := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	return tm2
}

// MonthEndTime 月末 => 23:59:59 获取月结束时间戳
func MonthEndTime(t time.Time) time.Time {
	e := MonthStartTime(t).AddDate(0, 1, -1)
	tm2 := time.Date(t.Year(), t.Month(), e.Day(), 23, 59, 59, 0, time.Local)
	return tm2
}

// YearStartTime 年初 => 00:00:00 获取年起始时间戳
func YearStartTime(t time.Time) time.Time {
	tm2 := time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
	return tm2
}

// YearEndTime 年末 => 23:59:59 获取年起始时间戳
func YearEndTime(t time.Time) time.Time {
	tm2 := time.Date(t.Year(), time.December, 31, 23, 59, 59, 0, time.Local)
	return tm2
}

// TimeDiff 时间差
func TimeDiff(t1, t2 time.Time) time.Duration {
	if t1.After(t2) {
		t1, t2 = t2, t1
	}
	return t2.Sub(t1)
}

// TimeDiffNaturalDays 计算两个时间之间的自然天数
func TimeDiffNaturalDays(t1, t2 time.Time) int {
	if t1.After(t2) {
		t1, t2 = t2, t1
	}
	d1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
	d2 := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, t2.Location())
	// 计算天数差
	days := int(d2.Sub(d1).Hours()/24) + 1
	return days
}
