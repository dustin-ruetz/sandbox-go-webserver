package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"log"
)

func Test_API_Handler(t *testing.T) {
	// type Response struct {
	// 	Route   string `json:"route"`
	// 	Request struct {
	// 		Path string `json:"path"`
	// 	} `json:"request"`
	// }

	request := httptest.NewRequest(http.MethodGet, "/api/", nil)
	responseRecorder := httptest.NewRecorder()
	log.Println("request =", request)
	log.Println("responseRecorder =", responseRecorder)

	API(responseRecorder, request)
	response := responseRecorder.Result()
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	// var unmarshalledJSON Response
	// json.Unmarshal(data, &unmarshalledJSON)
	log.Println("response =", response)
	log.Println("data =", data)
	log.Println("err =", err)
	// log.Println("unmarshalledJSON =", unmarshalledJSON)

	if err != nil {
		t.Errorf("Expected error to be `nil`; received %v", err)
	}
	// if string(data) != "ABC" {
	// 	t.Errorf("Expected ABC got %v", string(data))
	// }
}
