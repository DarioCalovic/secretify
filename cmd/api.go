package main

import (
	"log"

	pkgapi "github.com/DarioCalovic/secretify/pkg/api"
	"github.com/DarioCalovic/secretify/pkg/api/file"
	"github.com/DarioCalovic/secretify/pkg/api/secret"
	"github.com/DarioCalovic/secretify/pkg/api/setting"
	utilconfig "github.com/DarioCalovic/secretify/pkg/util/config"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
	"github.com/DarioCalovic/secretify/pkg/util/mail"
	"github.com/DarioCalovic/secretify/pkg/util/ticker"
)

func main() {
	cfg, err := utilconfig.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Connect to DB
	sqlitedb := utildb.NewSQLiteDB(cfg.DB.ConnectionURL)
	err = sqlitedb.Initialize()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Init mailer
	mailer := mail.NewMailer(cfg.SMTP.MailJet.APIKeyPublic, cfg.SMTP.MailJet.APIKeyPrivate)

	// Run Ticker
	cfgService := setting.Initialize(cfg)
	fileService := file.Initialize(sqlitedb, cfgService)
	secretService := secret.Initialize(sqlitedb, cfgService, fileService, mailer)
	t := ticker.New(cfg, secretService)
	go t.RunTask()

	err = pkgapi.Start(cfg, sqlitedb, mailer)
	if err != nil {
		log.Fatal(err)
	}
}
