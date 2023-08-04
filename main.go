package main

import (
	"net/http"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	server := &http.Server{Addr: ":8080"}
	server.SetKeepAlivesEnabled(false)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<body style='background-color:" + os.Getenv("BG_COLOR") + ";'>" +
			"<h1 style='color:" + os.Getenv("HEADING_COLOR") + ";'>" + os.Getenv("HEADING_TEXT") + "</h1>" +
			"<h3>served from " + hostname + "</h3>" +
			"</body>"))
	})

	server.ListenAndServe()
}
