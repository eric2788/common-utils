package request

import (
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Resp *http.Response
	Config *Config
	req *Requester
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

// Scan - not available when Requester already has a response struct
func (r *Response) Scan(res interface{}) error {
	defer r.Resp.Body.Close()
	b, err := io.ReadAll(r.Resp.Body)
	if err != nil {
		return err
	}
	var decoder Decoder = r.req.DefaultDecoder
	if r.Config.RespDecoder != nil {
		decoder = r.Config.RespDecoder
	}
	if decoder == nil {
		return fmt.Errorf("ResponseDecoder is nil")
	}
	return decoder(b, res)
}