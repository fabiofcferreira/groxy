package http

import (
	"encoding/json"
	"net/http"

	"github.com/fatih/color"

	"github.com/fabiofcferreira/groxy"
)

// GroxyHandler is a type definition for a function
type GroxyHandler func(w http.ResponseWriter, r *http.Request, c *groxy.Config) (int, error)

type response struct {
	Code    int
	Message string
}

// Wrap returns an http handler func that runs an API method itself, with access to the config
func Wrap(h GroxyHandler, c *groxy.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handle pre-flight requests
		w.Header().Set("Access-Control-Allow-Headers", "content-type, authorization, accept")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		// set json as the standard format for responses
		w.Header().Set("Content-Type", "application/json")

		// don't waste computing efforts when the request method is OPTIONS
		if r.Method == "OPTIONS" {
			return
		}

		var (
			code int
			err  error
		)

		defer func() {
			// if the return codes are 0 and no errors
			// it means the important data has been written into
			// the response body and the request is finished
			if code == 0 && err == nil {
				return
			}

			res := &response{
				Code: code,
			}

			if code != 0 {
				w.WriteHeader(code)
			}

			// Write JSON into response body
			data, e := json.MarshalIndent(res, "", "\t")
			if e != nil {
				color.Red(e.Error())
				return
			}
			w.Write(data)

			if err != nil {
				c.Logger.Error(err.Error())
			}

			return
		}()

		code, err = h(w, r, c)
	}
}
