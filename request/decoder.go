package request

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type Decoder func([]byte, interface{}) error

func JsonDecoder(data []byte, res interface{}) error {
	return json.Unmarshal(data, res)
}

func GobDecoder(data []byte, res interface{}) error {
	buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)
	return dec.Decode(res)
}
