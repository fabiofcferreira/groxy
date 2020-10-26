package http

import (
	"io/ioutil"
	"net/http"

	"github.com/fabiofcferreira/groxy"
)

// Proxy performs a request to the airtable API requested by the
// frontend and replies with the response, acting just like a proxy
func Proxy(w http.ResponseWriter, r *http.Request, c *groxy.Config) (int, error) {
	// For now the proxy only works with data retrievals
	if r.Method != "GET" {
		return http.StatusMethodNotAllowed, nil
	}

	// HTTP client
	client := &http.Client{}

	// Create AirTable API URL
	url := "https://api.airtable.com/v0/" + c.AppID + r.URL.String()

	// Create request and add required headers
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	// Perform request
	resp, err := client.Do(req)
	if err != nil {
		c.Logger.Errorf(err.Error())
	}
	defer resp.Body.Close()

	// Read response contents
	body, _ := ioutil.ReadAll(resp.Body)

	w.Write(body)

	return 0, nil
}
