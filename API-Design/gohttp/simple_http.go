package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Write([]byte("This is a POST request"))
			return
		}

		w.Write([]byte("Hello, World!"))
	})
	http.ListenAndServe(":8080", nil)
}
