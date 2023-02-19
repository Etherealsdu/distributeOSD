package main

import (
	"distributeOSD/apiServer/objects"
	"distributeOSD/dataServer/heartbeat"
	"distributeOSD/dataServer/locate"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
