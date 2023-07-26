package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/merakilab9/meradia/pkg/handler"
	"net/http"
)

func main() {

	//urls1 := "https://shopee.vn/api/v4/pages/get_category_tree"
	//urls2 := "https://shopee.vn/api/v4/recommend/recommend?bundle=category_landing_page&cat_level=2&catid=11035584&limit=24&no_filter=1&offset=0&section=category_landing_page_sec"
	//handler.GetUrls([]string{urls1, urls2}, "shoppee")

	r := mux.NewRouter()
	r.HandleFunc("/api/parse-to-json", handler.Service2Handler).Methods("POST")

	fmt.Println("Server started at http://localhost:8080")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
