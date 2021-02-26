package json

import (
	"encoding/json"

	"github.com/Akshit8/url-shortner/pkg/url"
	errs "github.com/pkg/errors"
)

// Serializer binds encode-decode methods
type Serializer struct{}

// Decode decodes byte array to json
func (s *Serializer) Decode(input []byte) (*urls.URL, error) {
	redirect := &urls.URL{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errs.Wrap(err, "serializer.json.Decode")
	}
	return redirect, nil
}

// Encode encodes a go struct to byte array
func (s *Serializer) Encode(input *urls.URL) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.json.Encode")
	}
	return rawMsg, nil
}
