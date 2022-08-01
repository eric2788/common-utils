package request

import (
	"github.com/magiconair/properties/assert"
	"net/http"
	"regexp"
	"testing"
)

const (
	Url           = "https://blive.ericlamm.xyz/"
	StatusPattern = `.*working.*`
)

func TestReadForRegex(t *testing.T) {
	res, err := http.Get(Url)
	if err != nil {
		t.Fatal(err)
	}
	// ReadForRegex has close method
	// defer res.Body.Close()
	// parse json string
	params, err := ReadForRegex(res, regexp.MustCompile(StatusPattern))

	assert.Equal(t, len(params), 1)
}
