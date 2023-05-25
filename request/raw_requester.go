package request

type RawRequester struct {
	req *Requester
}

func (r *RawRequester) RequestWithConfig(config *Config) (*Response, error) {
	return r.req.RequestWithConfig(config, nil)
}

func (r *RawRequester) Request(url, method string, configurers ...Configurer) (*Response, error) {
	return r.req.Request(url, method, nil, configurers...)
}

func (r *RawRequester) Get(url string, configurers ...Configurer) (*Response, error) {
	return r.req.Get(url, nil, configurers...)
}

func (r *RawRequester) Post(url string, configurers ...Configurer) (*Response, error) {
	return r.req.Post(url, nil, configurers...)
}

func (r *RawRequester) Put(url string, configurers ...Configurer) (*Response, error) {
	return r.req.Put(url, nil, configurers...)
}

func (r *RawRequester) Delete(url string, configurers ...Configurer) (*Response, error) {
	return r.req.Delete(url, nil, configurers...)
}

func (r *RawRequester) Head(url string, configurers ...Configurer) (*Response, error) {
	return r.req.Head(url, nil, configurers...)
}

func (r *RawRequester) Options(url string, configurers ...Configurer) (*Response, error) {
	return r.req.Options(url, nil, configurers...)
}

func (r *RawRequester) Patch(url string, configurers ...Configurer) (*Response, error) {
	return r.req.Patch(url, nil, configurers...)
}
