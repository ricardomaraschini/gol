package parameters

// NewBool returns a new string parameter with all options set
func NewBool(opts ...Setter) Bool {
	s := Bool{}
	for _, o := range opts {
		o(&s)
	}
	return s
}

// Bool represents a boolean command line parameter
type Bool struct {
	Name  string
	Value bool
}

// SetDefault sets the Bool default value
func (b *Bool) SetDefault(def interface{}) {
	_, ok := def.(bool)
	if ok == false {
		return
	}

	b.Value = def.(bool)
}

// SetName set the parameter name
func (b *Bool) SetName(name string) {
	b.Name = name
}

// GetType return this field type
func (b *Bool) GetType() string {
	return "bool"
}

// GetRef return a reference to the internal value
func (b *Bool) GetRef() interface{} {
	return &b.Value
}

// GetName return the parameter name
func (b *Bool) GetName() string {
	return b.Name
}
