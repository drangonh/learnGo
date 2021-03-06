package persist

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic"
	"gomodtest/crawler/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		//must set false on docker
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			//log.Printf("Item Saver: got item "+
			//	"#%d: %v", itemCount, item)

			err := save(client, item, index)
			if err != nil {
				fmt.Println(err)
			}
			itemCount++
		}
	}()

	return out, err
}

func save(client *elastic.Client, item engine.Item, index string) error {
	if item.Type == "" {
		return errors.New("must supply type")
	}

	if item.Id == "" {
		return errors.New("must supply id")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
