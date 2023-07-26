package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//func GetUrls(urls []string) ([][]byte, error) {
//	method := "GET"
//	client := &http.Client{}
//
//	var bodies [][]byte
//	for _, url := range urls {
//		req, err := http.NewRequest(method, url, nil)
//		if err != nil {
//			return nil, err
//		}
//		res, err := client.Do(req)
//		if err != nil {
//			return nil, err
//		}
//		defer res.Body.Close()
//
//		body, err := ioutil.ReadAll(res.Body)
//		if err != nil {
//			return nil, err
//		}
//
//		bodies = append(bodies, body)
//	}
//
//	for _, body := range bodies {
//		fmt.Print(string(body))
//	}
//	return bodies, nil
//}

func Service2Handler(w http.ResponseWriter, r *http.Request) {

	//// Retrieve the URL parameter from the request
	//url := r.URL.Query().Get("api_url")
	//if url == "" {
	//	http.Error(w, "No API URL provided", http.StatusBadRequest)
	//	return
	//}
	//
	//// Call the other API to get the `urls` slice
	//resp, err := http.Get(url)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	http.Error(w, "Failed to fetch URLs from the other API", resp.StatusCode)
	//	return
	//}

	var urls []string
	err := json.NewDecoder(r.Body).Decode(&urls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(urls) == 0 {
		http.Error(w, "No URLs provided", http.StatusBadRequest)
		return
	}

	method := "GET"
	client := &http.Client{}

	var bodies [][]byte
	for _, url := range urls {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bodies = append(bodies, body)
	}

	var result string
	for _, body := range bodies {
		result += string(body)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, result)
}
