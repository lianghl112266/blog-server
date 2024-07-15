package test

import (
	"flag"
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	x := flag.Bool("x", false, "xxxx")
	flag.Parse()
	fmt.Print(x)
}
