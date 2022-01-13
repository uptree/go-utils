package datetime

import (
	"testing"
)

func TestTime(t *testing.T) {
	t.Logf("Time: %d", Time())
}

func TestMilliTime(t *testing.T) {
	t.Logf("MilliTime: %d", MilliTime())
}

func TestMicroTime(t *testing.T) {
	t.Logf("MicroTime: %d", MicroTime())
}

func TestDateTime(t *testing.T) {
	t.Logf("Datetime: %s", Datetime())
}

func TestTimestamp(t *testing.T) {
	datetime := "2021-06-06 11:11:11"
	t.Logf("Date: %d", Timestamp(datetime, DefaultLayout))
}

func TestToday(t *testing.T) {
	t.Logf("Today: %s", Today())
}

func TestTodayEndTime(t *testing.T) {
	t.Logf("TodayEndTime: %d", TodayEndTime())
}

func TestWeekStartTime(t *testing.T) {
	t.Logf("WeekStartTime: %d", WeekStartTime())
}

func TestMonthStartTime(t *testing.T) {
	t.Logf("MonthStartTime: %d", MonthStartTime())
}
