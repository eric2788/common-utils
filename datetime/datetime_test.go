package datetime

import (
	"github.com/stretchr/testify/assert"
	"log"
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

	tz := os.Getenv("TZ")

	if tz == "" {
		tz = "Asia/Hong_Kong"
	}

	log.Println("Test with timezone: ", tz)
	// assign expected with different timezone result from environment variable
	var expected string
	switch tz {
	case "Asia/Hong_Kong":
		expected = "2021-09-01 21:24:29"
	case "Europe/London":
		expected = "2021-09-01 13:24:29"
	case "America/New_York":
		expected = "2021-09-01 05:24:29"
	default:
		log.Println("The test with timezone: ", tz, " is not supported")
		log.Println("pass the test")
		return
	}

	assert.Equal(t, expected, FormatMillis(date.UnixMilli()))
}
