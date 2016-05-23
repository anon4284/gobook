package main

import (
	"projects/gobook/server/db"
	"projects/gobook/server/entry"
	"projects/gobook/server/routes"
	"projects/gobook/server/util"
)

func main() {

	myDb := db.New(db.NewConfigFromJSON("creds.json"))
	myEntry := entry.New(myDb.SVC)
	myEntry.Init()

	resp, err := myEntry.Add("temp@gmail.com", "temp", "temp")
	util.LogErrOrResp(resp, err)

	resp2, err := myEntry.Scan()
	util.LogErrOrResp(resp2, err)

	router := router.New(5000)
	myEntry.EnableRouteAdd(router.Router, "/api/entry/add")
	router.ServeHTMLifNotFound("./public/web/index.html")
	router.EnableTestRoute()
	router.Start()
}
