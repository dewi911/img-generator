package main

import (
	"flag"
	"imggenerator/configs"
	"imggenerator/internal/server"
	"log"
)

var confPath = flag.String("conf-path", "./configs/.env", "Path to config env.")

func main() {
	conf, err := configs.New(*confPath)
	if err != nil {
		log.Fatal(err)
	}
	server.Run(conf)

}
