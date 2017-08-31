package parameters

// NewString returns a new string parameter with all options set
func NewString(opts ...Setter) String {
	s := String{}
	for _, o := range opts {
		o(&s)
	}
	return s
}

// String represents a string command line parameter
type String struct {
	Name  string
	Value string
}

// SetDefault sets the String default value
func (s *String) SetDefault(def interface{}) {
	_, ok := def.(string)
	if ok == false {
		return
	}

	s.Value = def.(string)
}

// SetName set the parameter name
func (s *String) SetName(name string) {
	s.Name = name
}

// GetType return this field type
func (s *String) GetType() string {
	return "string"
}

// GetRef return a reference to the internal value
func (s *String) GetRef() interface{} {
	return &s.Value
}

// GetName return the parameter name
func (s *String) GetName() string {
	return s.Name
}
