package request

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type (
	Requester struct {
		BaseUrl string
		Headers map[string]string
		Cookies map[string]string
		
		Client *http.Client
		Timeout time.Duration
	
		RequestIntercepters []func(*http.Request) error
		ResponseIntercepters []func(*http.Response) error

		DefaultDecoder Decoder
		DefaultEncoder Encoder
	}

	RequesterFactory func(*Requester)

)

func (r *Requester) do(req *http.Request) (*http.Response, error) {

	// request intercepts
	for _, intercepter := range r.RequestIntercepters {
		err := intercepter(req)
		if err != nil {
			return nil, err
		}
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	// response intercepts
	for _, intercepter := range r.ResponseIntercepters {
		err := intercepter(resp)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (r *Requester) Raw() *RawRequester {
	return &RawRequester{r}
}

func (r *Requester) RequestWithConfig(config *Config, res interface{}) (*Response, error) {
	

	var requestUrl string

	match, err := regexp.Match(`^https?://`, []byte(config.Url))
	if err != nil {
		// regex pattern has problem
		panic(err)
	}

	// set url
	if match {
		requestUrl = config.Url
	} else {
		requestUrl = r.BaseUrl + config.Url
	}

	// set method
	method := config.Method
	if method == "" {
		method = "GET"
	}

	// set payload
	var payload io.Reader
	if config.Data != nil {
		var err error
		
		var encoder Encoder = r.DefaultEncoder

		if config.DataEncoder != nil {
			encoder = config.DataEncoder
		}
		
		if encoder == nil {
			return nil, fmt.Errorf("PayloadEncoder is nil")
		}

		payload, err = config.Data.ToReader(encoder)

		if err != nil {
			return nil, err
		}
	}

	// set params
	if len(config.Params) > 0 {
		v := url.Values{}
		for key, value := range config.Params {
			v.Set(key, fmt.Sprint(value))
		}
		if strings.Contains(requestUrl, "?") {
			requestUrl += "&" + v.Encode()
		} else {
			requestUrl += "?" + v.Encode()
		}
	}

	// create request
	req, err := http.NewRequest(method, requestUrl, payload)
	if err != nil {
		return nil, err
	}

	// set headers
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	// override headers with config
	for key, value := range config.Headers {
		req.Header.Set(key, value)
	}

	// set cookies
	for key, value := range r.Cookies {
		req.AddCookie(&http.Cookie{
			Name: key,
			Value: value,
		})
	}

	// override cookies with config
	for key, value := range config.Cookies {
		req.AddCookie(&http.Cookie{
			Name: key,
			Value: value,
		})
	}

	// do request
	resp, err := r.do(req)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Resp: resp,
		Config: config,
		req: r,
	}

	// no response struct passed
	if res == nil {
		return response, nil
	} else {
		return response, response.Scan(res)
	}
}

func (r *Requester) Request(url, method string, res interface{}, configurers ...Configurer) (*Response, error) {
	config := &Config{
		Url: url,
		Method: method,
	}
	for _, configurer := range configurers {
		configurer(config)
	}
	return r.RequestWithConfig(config, res)
}

func (r *Requester) Get(url string, res interface{}, configurers ...Configurer) (*Response, error) {
	return r.Request(url, http.MethodGet, res, configurers...)
}

func (r *Requester) Post(url string, res interface{}, configurers ...Configurer) (*Response, error) {
	return r.Request(url, http.MethodPost, res, configurers...)
}

func (r *Requester) Put(url string, res interface{}, configurers ...Configurer) (*Response, error) {
	return r.Request(url, http.MethodPut, res, configurers...)
}

func (r *Requester) Patch(url string, res interface{}, configurers ...Configurer) (*Response, error) {
	return r.Request(url, http.MethodPatch, res, configurers...)
}

func (r *Requester) Delete(url string, res interface{}, configurers ...Configurer) (*Response, error) {
	return r.Request(url, http.MethodDelete, res, configurers...)
}

func (r *Requester) Head(url string, res interface{}, configurers ...Configurer) (*Response, error) {
	return r.Request(url, http.MethodHead, res, configurers...)
}

func (r *Requester) Options(url string, res interface{}, configurers ...Configurer) (*Response, error) {
	return r.Request(url, http.MethodOptions, res, configurers...)
}



