package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetDataFromService1() {
	service1URL := "http://localhost:8080/api/crawl"

	// Gửi yêu cầu HTTP tới Service1 để lấy dữ liệu đã crawl
	resp, err := http.Get(service1URL)
	if err != nil {
		fmt.Println("Lỗi khi gửi yêu cầu tới Service1:", err)
		return
	}
	defer resp.Body.Close()

	// Đọc dữ liệu từ phản hồi
	var dataSer1response []string
	err = json.NewDecoder(resp.Body).Decode(&dataSer1response)
	if err != nil {
		fmt.Println("Lỗi khi đọc dữ liệu từ Service1:", err)
		return
	}

	// Tiếp tục xử lý dữ liệu theo nhu cầu
	fmt.Println("Dữ liệu đã nhận được từ Service1:", dataSer1response)
}
