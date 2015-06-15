package main

import (
	"flag"
	"fmt"
	"github.com/open-falcon/forwarder/g"
	"github.com/open-falcon/forwarder/http"
	"os"
)

func main() {

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	g.InitRpcClients()

	go http.Start()

	select {}
}
