# gol - my GO Library
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

	// app --verbose --server=serveraddr

}
```

## Reading a file and decoding it using json

FileReader reads a file content and returns it as a slice of bytes. If a
Decoder is present(as for example a json decoder), the content is also
sent to the decoder. See example below

```go

allTasks := make([]task.Task, 0)
jdec := jsondec.New(
	jsondec.WithTarget(&allTasks),
)
freader := filereader.New(
	filereader.WithFilePath("/path/to/json/file/with/tasks.json"),
	filereader.WithDecoder(&jdec),
)

rawContent, err := freader.Parse()
// allTasks now contains all content of the file parsed as task.Task struct
```

## Requesting http and decoding result

```go
package main

import (
	"fmt"
	"log"

	"github.com/ricardomaraschini/gol/httpc"
	"github.com/ricardomaraschini/gol/jsondec"
)

type Body struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {

	body := new(Body)
	bodyDecoder := jsondec.New(
		jsondec.WithTarget(body),
	)

	c := httpc.New(
		httpc.WithDecoder(&bodyDecoder),
	)

	raw, err := c.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(raw))
	fmt.Println(body.URL)
	fmt.Println(body.Origin)
}

```
