package jsondec

import (
	"encoding/json"
	"errors"
)

// New returns a new JSONDecoder
func New(opts ...Setter) JSONDecoder {
	jdec := JSONDecoder{}
	for _, o := range opts {
		o(&jdec)
	}
	return jdec
}

// JSONDecoder is our wrapper around json unmashal
type JSONDecoder struct {
	target interface{}
}

// SetTarget set the destination for the unmarshal process
func (j *JSONDecoder) SetTarget(t interface{}) {
	j.target = t
}

// Decode decodes the []byte into target object
func (j *JSONDecoder) Decode(b []byte) error {
	if j.target == nil {
		return errors.New("invalid target found")
	}

	return json.Unmarshal(b, &j.target)
}

// GetTarget returns the target object
func (j *JSONDecoder) GetTarget() interface{} {
	return j.target
}
