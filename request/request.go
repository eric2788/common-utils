package request

import (
	"net/http"
	"time"
)


func New(factories ...RequesterFactory) *Requester {
	req := &Requester{
		Headers: make(map[string]string),
		Cookies: make(map[string]string),
		Client: http.DefaultClient,
		Timeout: time.Second * 30,
		DefaultDecoder: JsonDecoder,
		DefaultEncoder: JsonEncoder,
	}
	for _, factory := range factories {
		factory(req)
	}
	return req
}

func WithDefaultDecoder(decoder Decoder) RequesterFactory {
	return func(req *Requester) {
		req.DefaultDecoder = decoder
	}
}

func WithDefaultEncoder(encoder Encoder) RequesterFactory {
	return func(req *Requester) {
		req.DefaultEncoder = encoder
	}
}

func WithBaseUrl(baseUrl string) RequesterFactory {
	return func(req *Requester) {
		req.BaseUrl = baseUrl
	}
}

func WithHeaders(headers map[string]string) RequesterFactory {
	return func(req *Requester) {
		req.Headers = headers
	}
}

func WithCookies(cookies map[string]string) RequesterFactory {
	return func(req *Requester) {
		req.Cookies = cookies
	}
}

func WithClient(client *http.Client) RequesterFactory {
	return func(req *Requester) {
		client.Timeout = req.Timeout
		req.Client = client
	}
}

func WithTimeout(timeout time.Duration) RequesterFactory {
	return func(req *Requester) {
		req.Timeout = timeout
		req.Client.Timeout = timeout
	}
}

func WithRequestIntercepters(intercepters ...func(*http.Request) error) RequesterFactory {
	return func(req *Requester) {
		req.RequestIntercepters = intercepters
	}
}

func WithResponseIntercepters(intercepters ...func(*http.Response) error) RequesterFactory {
	return func(req *Requester) {
		req.ResponseIntercepters = intercepters
	}
}

func AddHeader(key, value string) RequesterFactory {
	return func(req *Requester) {
		req.Headers[key] = value
	}
}

func AddCookie(key, value string) RequesterFactory {
	return func(req *Requester) {
		req.Cookies[key] = value
	}
}

func AddRequestIntercepter(intercepter func(*http.Request) error) RequesterFactory {
	return func(req *Requester) {
		req.RequestIntercepters = append(req.RequestIntercepters, intercepter)
	}
}

func AddResponseIntercepter(intercepter func(*http.Response) error) RequesterFactory {
	return func(req *Requester) {
		req.ResponseIntercepters = append(req.ResponseIntercepters, intercepter)
	}
}
