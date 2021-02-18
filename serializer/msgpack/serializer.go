package msgpack

import (
	"github.com/Akshit8/url-shortner/url"
	errs "github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

// Redirect struct
type Redirect struct{}

// Decode decodes byte array to json
func (r *Redirect) Decode(input []byte) (*url.Redirect, error) {
	redirect := &url.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errs.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode encodes a go struct to byte array
func (r *Redirect) Encode(input *url.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
