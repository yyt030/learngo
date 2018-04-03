package persist

import (
	"context"
	"log"

	"learngo/crawler/engine"

	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("itemSaver got item: #%d: %v", itemCount, item)
			itemCount++

			err := saver(client, index, item)
			if err != nil {
				log.Printf("item saver error: saving item %v: %v",
					item, err)
			}
		}
	}()
	return out, nil
}

func saver(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply type")
	}

	indexService := client.Index().
		Index(index).
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
