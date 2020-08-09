package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	boolOpt   = flag.Bool("bool-option", false, "the bool option message for usage")
	stringOpt = flag.String("s", "", "the string option message for usage")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: %s [OPTIONS] <greeting message>\n", flag.CommandLine.Name())
		fmt.Fprintf(os.Stderr, "\nDescription: This is greeting program.\n\nOptions:\n")
		flag.PrintDefaults()
	}

}

// $ ./flag  -s "aaa" --bool-option  hoge
func main() {
	flag.Parse()
	fmt.Println("--bool-option:", *boolOpt)
	fmt.Println("-s:", *stringOpt)
	fmt.Println("args:", flag.Args())
}
