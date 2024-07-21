package tools

import (
	"time"
)

func DateToLocalString(t time.Time, format ...string) string {
	if t.IsZero() {
		return ""
	}
	f := time.DateTime
	if len(format) > 0 {
		f = format[0]
	}
	return t.Local().Format(f)
}

func DateToTime(t string, format ...string) time.Time {
	f := time.DateTime
	if len(format) > 0 {
		f = format[0]
	}
	tLocal, _ := time.ParseInLocation(f, t, time.Local)
	return tLocal
}

func GetFirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func GetLastDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 23, 59, 59, 999999999, t.Location()).AddDate(0, 1, -1)
}
