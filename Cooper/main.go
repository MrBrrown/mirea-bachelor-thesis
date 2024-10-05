package main

import (
	"fmt"

	db "example.com/coomper/DataBase"
	api "example.com/coomper/api"

	giga "example.com/coomper/assistantCore/detector"
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}

	err = giga.InitDetectot()
	if err != nil {
		panic(err)
	}

	x := giga.GetDetector()
	fmt.Println(x.Process("Какой процесс выполняет оборудование"))

	api.InitServer(":3333")
}
