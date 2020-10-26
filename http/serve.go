package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fatih/color"

	"github.com/fabiofcferreira/groxy"
)

// Serve spawns an http server that acts as a proxy between frotnend
// and the airtable api
func Serve(cfg *groxy.Config) {
	http.Handle("/", Wrap(Proxy, cfg))

	// Spawn http server
	fmt.Println()
	color.HiGreen("Starting HTTP server...")
	if err := http.ListenAndServe(cfg.Host+":"+strconv.Itoa(cfg.Port), nil); err != nil {
		color.HiRed(err.Error())
		log.Fatal(err)
	}
}
