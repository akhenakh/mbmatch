package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/handlers"
	"github.com/namsral/flag"

	"github.com/akhenakh/mbmatch/mbtiles"
)

var (
	tilesPath       = flag.String("tilesPath", "./hawaii.mbtiles", "mbtiles file path")
	port            = flag.Int("port", 8000, "port to listen for HTTP")
	hostname        = flag.String("hostname", fmt.Sprintf("127.0.0.1:%d", *port), "the hostname to come back at tiles")
	debug           = flag.Bool("debug", false, "enable debug")
	enforceReferrer = flag.Bool("enforceReferrer", false, "enforce referrer check using hostname")

	pathTpl = []string{"osm-liberty-gl.style", "solarized-dark.style", "planet.json"}
)

type server struct {
	box         *packr.Box
	fileHandler http.Handler
}

func addAllowOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		h.ServeHTTP(w, r)
	})
}
func enforceReferrerHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if *hostname != "" && strings.HasPrefix(r.Referer(), "http://"+*hostname) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	// serve file normally
	if !isTpl(path) {
		s.fileHandler.ServeHTTP(w, r)
		return
	}

	sf, err := s.box.FindString(path)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
	}

	p := map[string]interface{}{"Hostname": *hostname}

	tmplt := template.New("tpl")
	tmplt, err = tmplt.Parse(sf)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
	}
	tmplt.Execute(w, p)
}

func main() {
	flag.Parse()

	if *tilesPath == "" {
		flag.Usage()
		return
	}

	db, err := mbtiles.NewDB(*tilesPath)
	if err != nil {
		log.Fatal(err)
	}
	db.Debug = *debug

	box := packr.NewBox("./htdocs")

	s := &server{
		box:         &box,
		fileHandler: http.FileServer(box),
	}
	http.Handle("/tiles/", addAllowOrigin(enforceReferrerHandler(db)))
	http.Handle("/", handlers.CompressHandler(addAllowOrigin(s)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func isTpl(path string) bool {
	for _, p := range pathTpl {
		if p == path {
			return true
		}
	}
	return false
}
