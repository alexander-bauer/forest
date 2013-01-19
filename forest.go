package main

import (
	"flag"
	"log"
	"os"
)

var (
	Version    = "0"
	minversion string

	l = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

// Default values for flags
const (
	Bind = "0.0.0.0" // Interface to bind to
	Port = 8866      // Port to listen on
)

// Flag declarations
var (
	fRoot = flag.String("root", "", "project root")

	fBind = flag.String("bind", Bind, "interface to bind to")
	fPort = flag.Int("port", Port, "port to listen on")
)

func main() {
	log.Printf("Forest %s%s\n", Version, minversion)
	flag.Parse()

	err := Serve(*fRoot, *fBind, *fPort)
	if err != nil {
		l.Printf("Forest error: %s", err)
	}
}
