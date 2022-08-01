package request

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

type HttpError struct {
	Code     int
	Status   string
	Response *http.Response
}

func (e HttpError) Error() string {
	return fmt.Sprintf("%v: %s", e.Code, e.Status)
}

func Get(url string, response interface{}) error {

	res, err := http.Get(url)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return &HttpError{
			Code:     res.StatusCode,
			Status:   res.Status,
			Response: res,
		}
	}

	return Read(res, response)
}

func GetHtml(url string) (string, error) {

	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	return ReadString(res)
}

func GetBytesByUrl(url string) (img []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = res.Body.Close()
	}()
	img, err = ioutil.ReadAll(res.Body)
	return
}

func ReadString(res *http.Response) (string, error) {
	var err error

	defer func() {
		err = res.Body.Close()
	}()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Read(res *http.Response, response interface{}) error {

	var err error

	defer func() {
		err = res.Body.Close()
	}()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, response)
	return err
}

// ReadForRegex read multiple regexes from response without reading all response body
// need to manually use defer res.body.Close()
func ReadForRegex(res *http.Response, regs ...*regexp.Regexp) ([]string, error) {
	bufReader := bufio.NewReader(res.Body)
	defer res.Body.Close()
	finder := make(map[int]string)

	var err error = nil

	for err != io.EOF {
		line, _, err := bufReader.ReadLine()

		if err != nil {
			return nil, err
		}

		content := string(line)
		for i, reg := range regs {
			if _, ok := finder[i]; !ok && reg.MatchString(content) {
				finder[i] = content
			}
		}

		if len(finder) == len(regs) {
			arr := make([]string, len(finder))
			for i, s := range finder {
				arr[i] = s
			}
			return arr, nil
		}
	}

	arr := make([]string, len(regs))
	for i := range arr {
		if content, ok := finder[i]; ok {
			arr[i] = content
		} else {
			arr[i] = ""
		}
	}
	return arr, nil

}
