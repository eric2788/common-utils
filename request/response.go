package request

import "net/http"

type Response struct {
	Resp *http.Response
	Config *Config
}

func (r *Response) GetHeader(key string) string {
	return r.Resp.Header.Get(key)
}

func (r *Response) GetHeaders() map[string][]string {
	return r.Resp.Header
}

func (r *Response) GetStatusCode() int {
	return r.Resp.StatusCode
}

func (r *Response) GetStatus() string {
	return r.Resp.Status
}

func (r *Response) GetCookie(key string) *http.Cookie {
	for _, cookie := range r.Resp.Cookies() {
		if cookie.Name == key {
			return cookie
		}
	}
	return nil
}

func (r *Response) GetCookies() []*http.Cookie {
	return r.Resp.Cookies()
}

func (r *Response) GetRequestUrl() string {
	return r.Resp.Request.URL.String()
}