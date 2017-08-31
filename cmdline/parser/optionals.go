package parser

// Optional is a function that act on a CmdLineParser
type Optional func(*CmdLineParser)

// Returns a function that adds a parameter on a CmdLineParser
func WithParam(p Parameter) Optional {
	return func(c *CmdLineParser) {
		c.AddParameters(p)
	}
}
