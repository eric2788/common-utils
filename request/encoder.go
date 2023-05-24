package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
)

type Encoder func(data map[string]interface{}) (io.Reader, error)


func JsonEncoder(data map[string]interface{}) (io.Reader, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonData), nil
}

func FormUrlEncodedEncoder(data map[string]interface{}) (io.Reader, error) {
	values := url.Values{}
	for key, value := range data {
		values.Set(key, fmt.Sprintf("%v", value))
	}
	return strings.NewReader(values.Encode()), nil
}

func MultiPartEncoder(data map[string]interface{}) (io.Reader, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	defer w.Close()

	for key, value := range data {
		switch v := value.(type) {
		case string:
			if err := w.WriteField(key, v); err != nil {
				return nil, err
			}
		case []byte:
			if err := w.WriteField(key, string(v)); err != nil {
				return nil, err
			}
		case io.Reader:
			reader := v

			if r, ok := reader.(io.Closer); ok {
				defer r.Close()
			}

			var fw io.Writer
			var err error
			if f, ok := reader.(*os.File); ok {
				if fw, err = w.CreateFormFile(key, f.Name()); err != nil {
					return nil, err
				}
			} else {
				if fw, err = w.CreateFormField(key); err != nil {
					return nil, err
				}
			}

			_, err = io.Copy(fw, reader)
			if err != nil {
				return nil, err
			}
		// 默認一律當 string 處理
		default:
			if err := w.WriteField(key, fmt.Sprint(value)); err != nil {
				return nil, err
			}
		}
	}

	return &b, nil
}
