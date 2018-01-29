package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8080")
}

// package main
//
// import (
// 	"fmt"
// 	"net/http"
// )
//
// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8080", nil)
// }
//
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "PEC")
// }
