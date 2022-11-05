package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	init, err := NewInitialize()
	if err != nil {
		log.Fatal("can not initialize app")
	}
	defer func(Db *sqlx.DB) {
		err := Db.Close()
		if err != nil {
		}
	}(init.database.DB)
	//ルーターセッテイング
	e := NewRouter(init)

	env := os.Getenv("ENV")
	if env == "Develop" || env == "AWS-Prod" {
		e.Logger.Fatal(e.Start(":80"))
	} else {
		e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	}
}
