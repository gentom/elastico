package app

import (
	"context"
	"encoding/json"

	elastic "github.com/olivere/elastic/v7"
)

type App struct {
	url      string
	username string
	password string
	client   *elastic.Client
}

func Start(url, username, password string) *App {
	es, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetMaxRetries(1),
		elastic.SetHealthcheckTimeoutStartup(0),
		elastic.SetBasicAuth(username, password),
	)
	if err != nil {
		panic(err)
	}
	defer es.Stop()
	return &App{
		url:      url,
		username: username,
		password: password,
		client:   es,
	}
}

func (app *App) DefineIndex(ctx context.Context, indexName string, mapping *json.RawMessage) error {
	bytes, err := json.Marshal(mapping)
	if err != nil {
		return err
	}
	_, err = app.client.CreateIndex(indexName).BodyString(string(bytes)).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
