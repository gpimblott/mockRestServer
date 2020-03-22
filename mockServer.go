package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

/* Define data location is on www subdirectory */
var dataPath string = `./www/`

/*
Handle a request for a URI and return the json file at that location in the data path.
*/
func returnFileData(w http.ResponseWriter, req *http.Request) {
	logRequest(req)

	filename := createFilenameFromUri(req.RequestURI)
	log.Printf("Reading file %s", filename)

	w.Header().Set("Content-Type", "application/json")

	f, err := os.Open(filename)
	defer f.Close() // Close after function return
	if err != nil {
		message := `{ "error" : "File Not Found"}`
		bytes, _ := json.Marshal(message)
		w.Write(bytes)
	} else {
		io.Copy(w, f)
	}
}

/*
Convert the received URI into a filename to load and return to the client.
*/
func createFilenameFromUri(uri string) string {
	var filename = ""
	if len(uri) <= 1 {
		filename = dataPath + "index.json"
	} else {
		filename = dataPath + uri[1:] + ".json"
	}
	return filename
}

func logRequest(req *http.Request) {
	log.Printf("Processing: %s", req.RequestURI)
}

/*
Return the specified env var or the fallback if not defined.
*/
func getEnvWithFallback(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

/*
Application entry point.
Define the HTTP handler and start the server on the port defined by
environment variable PORT or 8090 if not defined.
*/
func main() {
	port := getEnvWithFallback("PORT", "8090")
	dataPath = getEnvWithFallback( "DATA_DIR" , "./www/")

	log.Printf("Running server on port %s", port)
	http.HandleFunc("/", returnFileData)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
