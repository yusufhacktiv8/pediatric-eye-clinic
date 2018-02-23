package main

import "fmt"

func main() {
	a := App{}
	fmt.Printf("Running PEC Server...\n")
	a.Initialize(
		// os.Getenv("APP_DB_USERNAME"),
		// os.Getenv("APP_DB_PASSWORD"),
		// os.Getenv("APP_DB_NAME"))
		"pecadmin",
		"pecadmin123",
		"pecadmin")

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
