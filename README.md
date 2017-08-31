# gol
Some Go libraries for personal usage

## Command line parameters

```go
package main

import (
	"fmt"
	"github.com/ricardomaraschini/gol/cmdline/parameters"
	"github.com/ricardomaraschini/gol/cmdline/parser"
	"log"
)

func usage() {
	// print usage
	fmt.Println("usage")
}

func main() {

	server := parameters.NewString(
		parameters.WithName("server"),
	)

	flag := parameters.NewBool(
		parameters.WithName("verbose"),
	)

	cparser := parser.CmdLineParser{}
	err := cparser.AddParameters(
		&server,
		&flag,
	)
	if err != nil {
		log.Fatal(err)
	}
	cparser.Parse()

	// values not provided through command line
	if server.Value == "" || flag.Value == false {
		usage()
		return
	}

	fmt.Printf("server is %s\n", server.Value)
	fmt.Printf("verbose is %v\n", flag.Value)

}
```
