package request

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/corpix/uarand"
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

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", uarand.GetRandom())
	res, err := http.DefaultClient.Do(req)

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

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", uarand.GetRandom())
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	return ReadString(res)
}

func GetBytesByUrl(url string) (data []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", uarand.GetRandom())
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}else if res.StatusCode != 200 {
		return nil, &HttpError{
			Code:     res.StatusCode,
			Status:   res.Status,
			Response: res,
		}
	}

	defer func() {
		err = res.Body.Close()
	}()
	data, err = ioutil.ReadAll(res.Body)
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

	finder := make(map[int]string)

	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)

	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			return nil, err
		}

		content := string(scanner.Bytes())
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
