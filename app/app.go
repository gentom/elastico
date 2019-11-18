package app

import (
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
