package main

import (
	"flag"
	"fmt"

	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/persist"
	"learngo/crawler_distributed/rpcsupport"

	"github.com/gpmgo/gopm/modules/log"
	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal("%v", serveRpc(fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
