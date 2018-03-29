package persist

import (
	"fmt"
	"log"

	"github.com/bean-du/crawler/engine"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 1
		for {
			item := <-out
			log.Printf("ItemSaver : got item #%d, %v", itemCount, item)
			itemCount++
			err := save(client, item, index)
			if err != nil {
				log.Fatalf("Item Saver: save error saving  item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	resp, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}
