package main

import (
	"log"
	"projects/gobook/server/db"
	"projects/gobook/server/entry"
	"projects/gobook/server/routes"
)

func main() {
	router := router.New(5000)

	myDb := db.New(db.NewConfigFromJSON("creds.json"))
	myEntry := entry.New(*router.Router, myDb.SVC)
	myEntry.Init()
	resp, err := myEntry.Scan()
	if err != nil {
		log.Print(err)
	} else {
		log.Print(resp)
	}

	router.ServeHTMLifNotFound("./public/web/index.html")
	router.TestRoute()
	router.Start()
}
