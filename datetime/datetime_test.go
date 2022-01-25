package datetime

import (
	"github.com/stretchr/testify/assert"
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
	oTime := "2022-01-24T20:44:51Z"
	hk, err := time.LoadLocation("Asia/Hong_Kong")
	if err != nil {
		t.Fatal(err)
	}
	changed, err := ToTimeZone(oTime, hk)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "2022-01-25T04:44:51+08", FormatISO(changed))
}

func TestParseISO(t *testing.T) {
	iso := "2021-09-01T13:24:29Z"
	date, err := ParseISOStr(iso)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "2021-09-01 21:24:29", FormatMillis(date.UnixMilli()))
}
