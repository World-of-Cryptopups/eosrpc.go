package eosrpc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func request(client *http.Client, url string, data interface{}, v interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	r, err := client.Post(url, "application/json", bytes.NewBuffer(d))
	if err != nil {
		return err
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	return nil
}
