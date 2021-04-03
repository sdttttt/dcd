package main

import (
	"flag"

	"github.com/sdttttt/huck"
)

var conf *string = flag.String("c", "dcd.yml", "Specify a configuration file.")

func main() {
	flag.Parse()
	huck.Run(conf)
}
