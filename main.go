package main

import (
	"log"

	client "github.com/angelthump/cache-replication/client"
	server "github.com/angelthump/cache-replication/server"
	utils "github.com/angelthump/cache-replication/utils"
)

func main() {
	cfgPath, err := utils.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	err = utils.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	client.Initalize()
	server.Initalize()
}
