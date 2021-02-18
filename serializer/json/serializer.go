package json

import (
	"encoding/json"

	"github.com/Akshit8/url-shortner/url"
	errs "github.com/pkg/errors"
)

// Redirect struct
type Redirect struct{}

// Decode decodes byte array to json
func (r *Redirect) Decode(input []byte) (*url.Redirect, error) {
	redirect := &url.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errs.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode encodes a go struct to byte array
func (r *Redirect) Encode(input *url.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
