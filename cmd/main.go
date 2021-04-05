package main

import (
	"flag"

	"github.com/sdttttt/huck"
)

var conf *string = flag.String("c", huck.DEFAULT_CONFIG_FILE_NAME, "Specify a configuration file.")

func main() {
	flag.Parse()
	huck.Run(*conf)
}
