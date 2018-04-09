package persist

import (
	"learngo/crawler/engine"
	"learngo/crawler/persist"

	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Saver(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}

	return err
}
