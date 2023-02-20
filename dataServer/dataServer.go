package main

import (
	"distributeOSD/apiServer/objects"
	"distributeOSD/dataServer/heartbeat"
	"distributeOSD/dataServer/locate"
	"distributeOSD/dataServer/temp"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
