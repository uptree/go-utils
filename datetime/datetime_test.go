package datetime

import (
	"testing"
	"time"
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

func TestToTime(t *testing.T) {
	t.Logf("ToTime: %v", ToTime(Time()))
}

func TestStringToTime(t *testing.T) {
	datetime := "wrong datetime"
	t.Logf("StringToTime: %v", StringToTime(datetime, DefaultLayout))
}

func TestTimestamp(t *testing.T) {
	datetime := "2021-06-06 11:11:11"
	t.Logf("Date: %d", Timestamp(datetime, DefaultLayout))
}

func TestToday(t *testing.T) {
	t.Logf("Today: %s", Today())
}

func TestTodayStartTime(t *testing.T) {
	t.Logf("TodayStartTime: %v", TodayStartTime())
}

func TestTodayEndTime(t *testing.T) {
	t.Logf("TodayEndTime: %v", TodayEndTime())
}

func TestDayStartTime(t *testing.T) {
	t.Logf("DayStartTime: %v", DayStartTime(time.Now()))
}

func TestDayEndTime(t *testing.T) {
	t.Logf("DayEndTime: %v", DayEndTime(time.Now()))
}

func TestWeekStartTime(t *testing.T) {
	t.Logf("WeekStartTime: %v", WeekStartTime(time.Now()))
}

func TestWeekEndTime(t *testing.T) {
	t.Logf("WeekEndTime: %v", WeekEndTime(time.Now()))
}

func TestMonthStartTime(t *testing.T) {
	t.Logf("MonthStartTime: %v", MonthStartTime(time.Now()))
}

func TestMonthEndTime(t *testing.T) {
	t.Logf("MonthEndTime: %v", MonthEndTime(time.Now()))
}
func TestYearStartTime(t *testing.T) {
	t.Logf("YearStartTime: %v", YearStartTime(time.Now()))
}

func TestYearEndTime(t *testing.T) {
	t.Logf("YearEndTime: %v", YearEndTime(time.Now()))
}

func TestTimeDiff(t *testing.T) {
	diff := TimeDiff(YearStartTime(time.Now()), YearStartTime(time.Now()))
	t.Logf("TimeDiff: %v", diff)
	diff = TimeDiff(time.Now(), time.Now())
	t.Logf("TimeDiff: %v", diff)
}

func TestTimeDiffNaturalDays(t *testing.T) {
	diff := TimeDiffNaturalDays(YearStartTime(time.Now()), YearEndTime(time.Now()))
	t.Logf("TimeDiffNaturalDays: %v", diff)
	diff = TimeDiffNaturalDays(MonthEndTime(time.Now()), MonthStartTime(time.Now()))
	t.Logf("TimeDiffNaturalDays: %v", diff)
	diff = TimeDiffNaturalDays(WeekStartTime(time.Now()), time.Now())
	t.Logf("TimeDiffNaturalDays: %v", diff)
	diff = TimeDiffNaturalDays(time.Now(), time.Now())
	t.Logf("TimeDiffNaturalDays: %v", diff)
}
