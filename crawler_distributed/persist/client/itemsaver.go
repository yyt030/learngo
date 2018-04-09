package client

import (
	"log"

	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	cli, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver: got item #%d: %v", itemCount, item)
			itemCount++

			// Call RPC to save item
			result := ""
			err = cli.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("item saver: error: %v", err)
			}
		}
	}()
	return out, nil
}
