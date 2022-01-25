package datetime

import (
	"fmt"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	ISOFormat  = "2006-01-02T15:04:05Z07"
)

func ToTimeZone(iso string, location *time.Location) (time.Time, error) {
	t, err := ParseISOStr(iso)
	if err != nil {
		return t, err
	} else if location == nil {
		return t, fmt.Errorf("location is null")
	}
	return t.In(location), nil
}

func FormatISO(time time.Time) string {
	return time.Format(ISOFormat)
}

func FormatISOFromStr(timeStr string) (string, error) {
	t, err := time.Parse(TimeFormat, timeStr)
	if err != nil {
		return "", err
	}
	return FormatISO(t), nil
}

func FormatSeconds(ts int64) string {
	return FromSeconds(ts).Format(TimeFormat)
}

func ParseISOStr(iso string) (time.Time, error) {
	return time.Parse(ISOFormat, iso)
}

func FormatMillis(ts int64) string {
	return time.UnixMilli(ts).Format(TimeFormat)
}

func Duration(before, after int64) time.Duration {
	first, second := FromSeconds(before), FromSeconds(after)
	if first.After(second) {
		return first.Sub(second)
	} else {
		return second.Sub(first)
	}
}

func FromSeconds(seconds int64) time.Time {
	return time.Unix(seconds, 0)
}
