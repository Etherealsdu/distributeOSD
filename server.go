package main

import (
	"distributeOSD/objects"
	"distributeOSD/util"
	"log"
	"net/http"
)

func main() {
	config, err := util.LoadConfig(".")
	util.Pconfig = config
	if err != nil {
		log.Fatal("can not load config from app.env")
	}

	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(util.Pconfig.ListenAddress, nil))
}
