package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fabiofcferreira/groxy"
)

// Proxy performs a request to the airtable API requested by the
// frontend and replies with the response, acting just like a proxy
func Proxy(w http.ResponseWriter, r *http.Request, c *groxy.Config) (int, error) {
	// For now the proxy only works with data retrievals
	if r.Method != "GET" && r.Method != "POST" {
		return http.StatusMethodNotAllowed, nil
	}

	// HTTP client
	var resp *http.Response
	var err error
	client := &http.Client{}

	// Create AirTable API URL
	url := "https://api.airtable.com/v0/" + c.AppID + r.URL.String()

	if r.Method == "GET" {
		// Create request and add required headers
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Bearer "+c.APIKey)

		// Perform request
		resp, err = client.Do(req)
		if err != nil {
			c.Logger.Errorf(err.Error())
		}
		defer resp.Body.Close()
	} else if r.Method == "POST" {
		reqBody, err := json.Marshal(r.Body)
		if err != nil {
			c.Logger.Errorf(err.Error())
		}

		resp, err = http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			c.Logger.Errorf("Error while performing POST request: %s.", err)
		}
		defer resp.Body.Close()
	}

	// Read response contents
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Logger.Errorf("Couldn't read response: %s", err.Error())
	}

	w.Write(body)

	return 0, nil
}
