package main

import (
	"flag"

	"github.com/sdttttt/dcd"
)

var conf *string = flag.String("c", "dcd.yml", "Specify a configuration file.")

func main() {
	flag.Parse()
	dcd.Run(conf)
}
