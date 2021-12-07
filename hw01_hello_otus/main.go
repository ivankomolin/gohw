package main

import (
	"io"
	"os"

	"golang.org/x/example/stringutil"
)

func main() {
	io.WriteString(os.Stdout, stringutil.Reverse("Hello, OTUS!"))
}
