package main

import (
	"fmt"
	"os"
)

func main() {
	//err := godotenv.Load("../.env")
	fmt.Println("Hellow world")
	init, err := NewInitialize()
	if err != nil {
		return
	}
	//ルーターセッテイング
	e := NewRouter(init)
	//起動
	//e.Logger.Fatal(e.Start(":80"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
