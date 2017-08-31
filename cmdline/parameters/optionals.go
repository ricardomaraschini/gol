package parameters

// Customizable is something that we can set the name and the default
// value on
type Customizable interface {
	SetName(string)
	SetDefault(interface{})
}

// Setter is a function that sets a property on a Customizable through
// a method call
type Setter func(Customizable)

// WithName sets the name on a Customizable interface
func WithName(name string) Setter {
	return func(c Customizable) {
		c.SetName(name)
	}
}

// WithDefault sets the default value on a Customizable interface
func WithDefault(def interface{}) Setter {
	return func(c Customizable) {
		c.SetDefault(def)
	}
}
