package request

import (
	"io"
)

type (
	Config struct {
		Url    string
		Method string
		// query payload
		Data        *Payload
		DataEncoder Encoder
		RespDecoder Decoder
		// url query
		Params  map[string]interface{}
		Cookies map[string]string
		Headers map[string]string
	}

	Payload map[string]interface{}

	Configurer func(*Config)
)

func (p *Payload) Get(key string) interface{} {
	return (*p)[key]
}

func (p *Payload) Set(key string, value interface{}) {
	(*p)[key] = value
}

func (p *Payload) Del(key string) {
	delete(*p, key)
}

func (p *Payload) ToReader(encoder Encoder) (io.Reader, error) {
	return encoder(*p)
}



// Configurers

func ConfigureEncoder(encoder Encoder) Configurer {
	return func(c *Config) {
		c.DataEncoder = encoder
	}
}

func ConfigureDecoder(decoder Decoder) Configurer {
	return func(c *Config) {
		c.RespDecoder = decoder
	}
}

func ConfigureData(data map[string]interface{}) Configurer {
	return func(c *Config) {
		p := Payload(data)
		c.Data = &p
	}
}

func ConfigureParams(params map[string]interface{}) Configurer {
	return func(c *Config) {
		c.Params = params
	}
}

func ConfigureCookies(cookies map[string]string) Configurer {
	return func(c *Config) {
		c.Cookies = cookies
	}
}

func ConfigureHeaders(headers map[string]string) Configurer {
	return func(c *Config) {
		c.Headers = headers
	}
}