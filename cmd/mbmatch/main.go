package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/akhenakh/mbmatch/mbtiles"
)

var (
	path  = flag.String("path", "", "mbtiles file path")
	port  = flag.Int("port", 6000, "port to listen for HTTP")
	debug = flag.Bool("debug", false, "enable debug")
)

func main() {
	flag.Parse()

	if *path == "" {
		flag.Usage()
	}

	db, err := mbtiles.NewDB(*path)
	if err != nil {
		log.Fatal(err)
	}
	db.Debug = *debug

	http.HandleFunc("/tiles/", db.ServeHTTP)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
