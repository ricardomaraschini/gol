package parser

import (
	"flag"
	"fmt"
)

// Parameter represents a command line parameter. It must implemented a method
// to get a reference to the inner value, a method to get the parameter name
// and a method to get the type.
type Parameter interface {
	GetRef() interface{}
	GetName() string
	GetType() string
}

// CmdLineParser keeps all parameters and allow users to parse them
type CmdLineParser struct {
	StringParams []Parameter
	BoolParams   []Parameter
}

// New returns a new CmdLineParser setting all optionals passed in
func New(opts ...Optional) CmdLineParser {
	c := CmdLineParser{}
	for _, o := range opts {
		o(&c)
	}
	return c
}

// AddParameters adds a parameter to be parsed from the command line
func (c *CmdLineParser) AddParameters(pars ...Parameter) error {

	for _, p := range pars {
		switch p.GetType() {
		case "string":
			c.StringParams = append(c.StringParams, p)

		case "bool":
			c.BoolParams = append(c.BoolParams, p)

		default:
			return fmt.Errorf("unsupported type %s", p.GetType())
		}
	}

	return nil
}

// Parse parses all the command line. Does not return nothing as the values
// are going to be populated through reference inside each parameter
func (c *CmdLineParser) Parse() {

	for _, p := range c.StringParams {
		s, ok := p.GetRef().(*string)
		if ok == false {
			continue
		}
		flag.StringVar(s, p.GetName(), *s, "")
	}

	for _, p := range c.BoolParams {
		b, ok := p.GetRef().(*bool)
		if ok == false {
			continue
		}
		flag.BoolVar(b, p.GetName(), *b, "")
	}

	flag.Parse()
}

// Usage prints flag defaults usage
func (c *CmdLineParser) Usage() {
	flag.Usage()
}
