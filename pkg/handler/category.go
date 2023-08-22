package handler

import (
	"fmt"
	"gitlab.com/merakilab9/j4/pkg/service"
	"gitlab.com/merakilab9/meracore/ginext"
	"net/http"
)

type CateHandlers struct {
	service service.CateInterface
}

func NewCateHandlers(service service.CateInterface) *CateHandlers {
	return &CateHandlers{service: service}
}

func (h *CateHandlers) GetUrlCate(c *ginext.Request) (*ginext.Response, error) {

	rs, err := h.service.GetUrlCate()
	if err != nil {
		return nil, err
	}
	fmt.Println(rs)

	//data := model.RequestData{
	//	rs,
	//}

	//payload, err := json.Marshal(data)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}

	//url := "http://localhost:7099/api/v1/cate/parse-to-json"
	//method := "POST"
	//
	//client := &http.Client{}
	//req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//req.Header.Set("Content-Type", "application/json")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//fmt.Println(string(body))

	return ginext.NewResponseData(http.StatusOK, rs), nil

}
