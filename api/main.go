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

	defer a.DB.Close()

	a.Router.Run(":8080")
}
