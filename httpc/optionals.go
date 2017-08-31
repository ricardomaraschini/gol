package httpc

// Setter is a function that sets a property on HTTPC
type Setter func(*HTTPC)

// WithDecoder sets the decoder on HTTPC
func WithDecoder(d Decoder) Setter {
	return func(h *HTTPC) {
		h.SetDecoder(d)
	}
}
