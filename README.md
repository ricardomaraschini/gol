# gol
Some Go libraries for personal usage

## Command line parameters

```
package main

import (
	"fmt"
	"github.com/ricardomaraschini/gol/cmdline/parameters"
	"github.com/ricardomaraschini/gol/cmdline/parser"
)

func main() {

	server := parameters.NewString(
		parameters.WithName("server"),
	)

	unit := parameters.NewString(
		parameters.WithName("unit"),
	)

	cparser := parser.CmdLineParser{}
	err := cparser.AddParameters(
		&server,
		&unit,
	)
	if err != nil {
		log.Fatal(err)
	}
	cparser.Parse()

	// values not provided through command line
	if server.Value == "" || unit.Value == "" {
		usage()
		return
	}

	fmt.Printf("server is %s\n", server.Value)
	fmt.Printf("unit is %s\n", unit.Value)

}
```
