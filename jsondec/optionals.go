package jsondec

// Setter is a function that sets a property on a JSONDecoder
type Setter func(*JSONDecoder)

// WithTarget returns a Setter that sets the target on a JSONDecoder
func WithTarget(i interface{}) Setter {
	return func(j *JSONDecoder) {
		j.target = i
	}
}
