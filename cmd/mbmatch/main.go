package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/akhenakh/mbmatch/mbtiles"
	"github.com/gobuffalo/packr"
)

var (
	path  = flag.String("path", "", "mbtiles file path")
	port  = flag.Int("port", 7000, "port to listen for HTTP")
	debug = flag.Bool("debug", false, "enable debug")
)

func main() {
	flag.Parse()

	if *path == "" {
		flag.Usage()
		return
	}

	db, err := mbtiles.NewDB(*path)
	if err != nil {
		log.Fatal(err)
	}
	db.Debug = *debug

	box := packr.NewBox("./htdocs")

	http.HandleFunc("/tiles/", db.ServeHTTP)
	http.Handle("/", http.FileServer(box))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
