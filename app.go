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

	resp2, err := myEntry.Scan()
	util.LogErrOrResp(resp2, err)

	router := router.New(7016)
	myEntry.EnableRouteAdd(router.Router, "/api/entry/add")
	myEntry.EnableRouteGet(router.Router, "/api/entry/get")
	router.ServeHTMLifNotFound("./public/build/web/index.html")
	router.EnableTestRoute()
	router.Start()
}
