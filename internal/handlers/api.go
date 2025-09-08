package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func API(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Route   string `json:"route"`
		Request struct {
			Path string `json:"path"`
		} `json:"request"`
	}

	response := fmt.Sprintf(`
{
	"route": "/api/",
	"request": {
		"path": "%s"
	}
}
`,
		r.URL.Path)

	byteResponse := []byte(response)

	// Parse the JSON data; error out if the object's data types don't match the Response struct.
	var unmarshalledJSON Response
	if err := json.Unmarshal(byteResponse, &unmarshalledJSON); err != nil {
		log.Println("json.Unmarshal: Failed to unmarshal JSON.", "\n response =", response, "\n err =", err)
	}

	// Respond with the JSON data; error out if `w.Write` fails for any reason.
	if _, err := w.Write(byteResponse); err != nil {
		log.Fatal("w.Write: Failed to write response.", "\n response =", response, "\n err =", err)
	}
}
