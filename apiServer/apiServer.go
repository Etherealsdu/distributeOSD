package main

import (
	"distributeOSD/apiServer/heartbeat"
	"distributeOSD/apiServer/locate"
	"distributeOSD/apiServer/objects"
	"distributeOSD/apiServer/versions"
	"distributeOSD/dataServer/temp"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
