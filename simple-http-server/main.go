package main

import "net/http"

func main() {
	httpRootDir := http.Dir("files")
	handler := http.FileServer(httpRootDir)
	http.ListenAndServe(":8888", handler)
}
