package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hellow world")
	init, err := NewInitialize()
	if err != nil {
		return
	}
	//ルーターセッテイング
	e := NewRouter(init)

	env := os.Getenv("ENV")
	if env == "Develop" {
		e.Logger.Fatal(e.Start(":80"))
	} else {
		e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	}
	//起動
}
