package main

import (
	"flag"

	"github.com/sdttttt/huck"
)

var conf *string = flag.String("c", huck.DefaultConfigFileName, "Specify a configuration file.")

func main() {
	flag.Parse()
	huck.Run(*conf)
}
