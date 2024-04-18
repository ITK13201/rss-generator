package infrastructure

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewSqlClient(cfg *domain.Config) *ent.Client {
	entOptions := []ent.Option{}
	if cfg.Debug {
		entOptions = append(entOptions, ent.Debug())
	}

	client, err := ent.Open("mysql", (*cfg).Database.DSN(), entOptions...)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
