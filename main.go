package main

import (
	"flag"
	"log"

	spt "github.com/deepset/httphashtool/script"
)

func main() {

	var (
		serverURL []string
		parallel  *int
	)

	//Reading command line arguments at start of script
	parallel = flag.Int("parallel", 10, "Number of concurrent request allowed")
	flag.Parse()
	serverURL = flag.Args()

	// checking if no url is given to parse
	if len(serverURL) == 0 {
		log.Fatalln("Enter urls to parse!")
	}

	// createWorkers will create parallel workers and get hashes of given serverURL
	err := spt.CreateWorkers(serverURL, *parallel)
	if err != nil {
		log.Fatalln(err)
	}
}
