package datetime

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {

	before := time.Now().Unix() - 86400

	t.Logf("before: %d\n", before)

	after := time.Now().Unix()

	t.Logf("after: %d\n", after)

	assert.Equal(t, time.Duration(24), Duration(before, after)/time.Hour)
}

func TestChangeTimeZone(t *testing.T) {
	oTime := "2022-01-30T14:00:00Z"
	hk, err := time.LoadLocation("Asia/Hong_Kong")
	if err != nil {
		t.Fatal(err)
	}
	changed, err := ToTimeZone(oTime, hk)
	if err != nil {
		t.Fatal(err)
	}

	result := FormatISO(changed)

	assert.Equal(t, "2022-01-30T22:00:00+08", result)

	date, err := ParseISOStr(result)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 22, date.Hour())
	assert.Equal(t, 0, date.Minute())
	assert.Equal(t, 0, date.Second())

}

func TestParseISO(t *testing.T) {
	iso := "2021-09-01T13:24:29Z"
	date, err := ParseISOStr(iso)
	if err != nil {
		t.Fatal(err)
	}

	// assign expected with different timezone result from environment variable
	var expected string
	switch os.Getenv("TZ") {
	case "Asia/Hong_Kong":
		expected = "2021-09-01T22:24:29+08"
	case "Europe/London":
		expected = "2021-09-01T13:24:29+00"
	case "America/New_York":
		expected = "2021-09-01T05:24:29-04"
	default:
		otherTz, err := time.LoadLocation("Asia/Hong_Kong")
		if err != nil {
			t.Fatal(err)
		}
		ti, err := ToTimeZone(iso, otherTz)
		if err != nil {
			t.Fatal(err)
		}
		expected = FormatISO(ti)
	}
	assert.Equal(t, expected, FormatISO(date))
}
