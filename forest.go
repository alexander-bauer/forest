package main

import (
	"log"
	"os"
)

var (
	Version    = "0"
	minversion string

	l = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

func main() {
	logInfo(l)
}

func logInfo(l *log.Logger) {
	l.Printf(`Forest v%s`, Version+minversion)
}
