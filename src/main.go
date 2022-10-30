package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("go app running...")
	init, err := NewInitialize()
	if err != nil {
		return
	}
	//ルーターセッテイング
	e := NewRouter(init)

	env := os.Getenv("ENV")
	if env == "Develop" || env == "AWS-Prod" {
		e.Logger.Fatal(e.Start(":80"))
	} else {
		e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	}
}
