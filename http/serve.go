package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fabiofcferreira/groxy"
)

// Serve spawns an http server that acts as a proxy between frotnend
// and the airtable api
func Serve(c *groxy.Config) {
	// r := mux.NewRouter()

	// r.HandleFunc("/proxy", Wrap(Proxy, c)).Methods("GET", "POST", "OPTIONS")
	http.Handle("/", Wrap(Proxy, c))
	if err := http.ListenAndServe(":"+strconv.Itoa(c.Port), nil); err != nil {
		log.Fatal(err)
	}
}
