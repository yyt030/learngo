package persist

import (
	"context"
	"log"

	"learngo/crawler/engine"

	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver() chan engine.Item {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("itemSaver got item: #%d: %v", itemCount, item)
			itemCount++

			err := saver(item, client)
			if err != nil {
				log.Printf("item saver error: saving item %v: %v",
					item, err)
			}
		}
	}()
	return out
}

func saver(item engine.Item, client *elastic.Client) error {
	if item.Type == "" {
		return errors.New("must supply type")
	}

	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
