package request

import (
	"fmt"
	"github.com/eric2788/common-utils/regex"
	"net/http"
	"regexp"
	"testing"
)

const (
	ChannelId       = "UCIaC5td9nGG6JeKllWLwFLA"
	LiveKeyword     = "<link rel=\"canonical\" href=\"https://www.youtube.com/watch\\?v=(?P<id>\\w+)\">"
	UpcomingKeyword = "\"isUpcoming\":true"
)

var (
	idRegex       = regexp.MustCompile(LiveKeyword)
	upcomingRegex = regexp.MustCompile(UpcomingKeyword)
)

func TestReadForRegex(t *testing.T) {
	status, err := getChannelStatus(ChannelId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(status)
}

type ChannelStatus struct {
	Type string
	Id   string
}

func getChannelStatus(channelId string) (*ChannelStatus, error) {

	res, err := http.Get(fmt.Sprintf("https://youtube.com/channel/%s/live", channelId))

	if err != nil {
		return nil, err
	} else if res.StatusCode == 404 {
		return nil, fmt.Errorf("not found channel %s", channelId)
	}

	defer res.Body.Close()

	contents, err := ReadForRegex(res, idRegex, upcomingRegex)

	if err != nil {
		return nil, err
	}

	if contents[0] == "" {
		return &ChannelStatus{Type: "NONE"}, nil // no streaming or upcoming
	} else {

		find := regex.GetParams(idRegex, contents[0])
		videoId := find["id"]

		status := &ChannelStatus{Id: videoId}

		if contents[1] != "" {
			status.Type = "UpComing"
		} else {
			status.Type = "Live"
		}

		return status, nil
	}
}

