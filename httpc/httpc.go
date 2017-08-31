package httpc

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// Decoder is an interface to an entity able to decode the body of the request,
// it must support a Decode() call so we can pass the body content to it
type Decoder interface {
	Decode([]byte) error
}

// New returns a new HTTPC object
func New(opts ...Setter) HTTPC {
	h := HTTPC{}
	for _, o := range opts {
		o(&h)
	}
	return h
}

// HTTPC is our writer around http requests
type HTTPC struct {
	Dec        Decoder
	statusLast int
}

// SetDecoder sets the decoder for the body content
func (h *HTTPC) SetDecoder(d Decoder) {
	h.Dec = d
}

// Success returns if last request http status code was between 200 and 300
func (h *HTTPC) Success() bool {
	return h.statusLast != 0 && h.statusLast >= 200 && h.statusLast < 300
}

// Get does a get into the url and returns the body as an slice of bytes
func (h *HTTPC) Get(url string) ([]byte, error) {

	h.statusLast = 0
	rep, err := http.Get(url)
	if err != nil {
		return make([]byte, 0), err
	}
	defer rep.Body.Close()
	h.statusLast = rep.StatusCode

	return h.processResultBody(rep.Body)
}

// Post posts json to url, using the []byte as the body of the request. No
// regular post fields supported.
func (h *HTTPC) Post(url string, body []byte) ([]byte, error) {

	h.statusLast = 0
	buf := bytes.NewBuffer(body)
	rep, err := http.Post(url, "application/json", buf)
	if err != nil {
		return make([]byte, 0), err
	}
	defer rep.Body.Close()
	h.statusLast = rep.StatusCode

	return h.processResultBody(rep.Body)
}

// processResultBody receives an io reader, converts it into a []byte and if
// Dec is set we call it, otherwise we return the []byte directly
func (h *HTTPC) processResultBody(body io.Reader) ([]byte, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return make([]byte, 0), nil
	}

	if h.Dec == nil {
		return b, nil
	}

	return b, h.Dec.Decode(b)
}
