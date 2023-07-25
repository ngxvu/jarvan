package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUrls(urls []string, filename string) ([][]byte, error) {
	method := "GET"
	client := &http.Client{}

	var bodies [][]byte
	for _, url := range urls {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		bodies = append(bodies, body)
	}

	for _, body := range bodies {
		fmt.Print(string(body))
	}
	// Write the response bodies to a file
	for i, body := range bodies {
		err := ioutil.WriteFile(fmt.Sprintf("%s_%d.txt", filename, i+1), body, 0644)
		if err != nil {
			return bodies, nil
		}
	}
	return bodies, nil
}
