package nibbler_elasticsearch

import (
	"errors"
	"github.com/markdicksonjr/nibbler"
	"github.com/olivere/elastic/v7"
)

type Extension struct {
	nibbler.NoOpExtension

	Client *elastic.Client
	Url    string

	username string
	password string
}

func (s *Extension) Init(app *nibbler.Application) error {
	var err error

	if app.Config == nil {
		return errors.New("app configuration not provided")
	}

	// if the Url attribute isn't set, find the config in environment variables
	if len(s.Url) == 0 {
		s.Url = app.Config.Raw.Get("elastic", "url").String("")

		if len(s.Url) == 0 {
			s.Url = app.Config.Raw.Get("database", "url").String("http://localhost:9200")
		}
	}

	s.username = app.Config.Raw.Get("elastic", "user").String("")
	s.password = app.Config.Raw.Get("elastic", "password").String("")

	options := []elastic.ClientOptionFunc{
		elastic.SetSniff(false),
		elastic.SetURL(s.Url),
	}

	if s.username != "" && s.password != "" {
		options = append(options, elastic.SetBasicAuth(s.username, s.password))
	}

	s.Client, err = elastic.NewClient(options...)

	return err
}

func (s *Extension) GetName() string {
	return "elasticsearch"
}

// TODO: these are pretty silly

func (s *Extension) NewMatchQuery(name string, text interface{}) *elastic.MatchQuery {
	return elastic.NewMatchQuery(name, text)
}

func (s *Extension) NewMatchAllQuery() *elastic.MatchAllQuery {
	return elastic.NewMatchAllQuery()
}

func (s *Extension) NewMatchNoneQuery() *elastic.MatchNoneQuery {
	return elastic.NewMatchNoneQuery()
}

func (s *Extension) NewMatchPhraseQuery(name string, value interface{}) *elastic.MatchPhraseQuery {
	return elastic.NewMatchPhraseQuery(name, value)
}

func (s *Extension) NewBoolQuery() *elastic.BoolQuery {
	return elastic.NewBoolQuery()
}

func (s *Extension) NewBulkDeleteRequest() *elastic.BulkDeleteRequest {
	return elastic.NewBulkDeleteRequest()
}

func (s *Extension) NewBulkIndexRequest() *elastic.BulkIndexRequest {
	return elastic.NewBulkIndexRequest()
}

func (s *Extension) NewBulkUpdateRequest() *elastic.BulkUpdateRequest {
	return elastic.NewBulkUpdateRequest()
}

func (s *Extension) NewIdsQuery(types ...string) *elastic.IdsQuery {
	return elastic.NewIdsQuery()
}
