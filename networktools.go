package main

import (
	"fmt"
	"net/http"
)

func get(url string) (response *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "christianjanev/mcpkg/0.0.1 (christianjanev7@gmail.com)")

	resp, error := client.Do(req)

	if resp.StatusCode != 200 {
		fmt.Println(error)
	}

	return resp, error
}
