package webserver

import (
	"encoding/json"
	"fmt"

	// Excerpts from `log` package description:
	// > That logger writes to standard error and prints the date and time of each logged message.
	// > Every log message is output on a separate line: if the message being printed does not end in a newline, the logger will add one.
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	const port = ":4444"
	const url = "http://localhost" + port

	log.Println("Starting webserver...")
	log.Println("Routes:")
	log.Printf("  %s/api/", url)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("http.ListenAndServe: Failed to start webserver.", "\n err: ", err)
	}
}
