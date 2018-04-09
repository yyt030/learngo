package main

import (
	"fmt"

	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/persist"
	"learngo/crawler_distributed/rpcsupport"

	"github.com/gpmgo/gopm/modules/log"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	log.Fatal("%v", serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort),
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
