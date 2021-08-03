package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
)

// main input: Takes in user input where the Osquery socket is located
// This function registers this Osquery extension using the user provided socket path
// main output: None
func main() {
	// Extract command line arguments
	flSocketPath := flag.String("socket", "", "path to osqueryd extensions socket")
	flTimeout := flag.Int("timeout", 0, "")
	flag.Int("interval", 0, "")
	flag.Parse()

	// allow for osqueryd to create the socket path
	timeout := time.Duration(*flTimeout) * time.Second
	time.Sleep(2 * time.Second)

	// initializing server objecet
	server, err := osquery.NewExtensionManagerServer("foobar", *flSocketPath, osquery.ServerTimeout(timeout))

	// If initializing server fails exit
	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}

	// Create and register a new table plugin with the server.
	// table.NewPlugin requires the table plugin name,
	// a slice of Columns and a generate function.
	// If server.Run() fails exit
	server.RegisterPlugin(table.NewPlugin("foobar", FoobarColumns(), FoobarGenerate))
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}

// FoobarColumns returns the columns that our table will return.
func FoobarColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("foo"),
		table.TextColumn("baz"),
	}
}

// FoobarGenerate will be called whenever the table is queried. It should return
// a full table scan.
func FoobarGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	return []map[string]string{
		{
			"foo": "bar",
			"baz": "baz",
		},
		{
			"foo": "bar",
			"baz": "baz",
		},
	}, nil
}