// main.go

package main

import (
	"os"
)

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))


    /*fmt.Println(os.Getenv("APP_DB_USERNAME"))
	fmt.Println(os.Getenv("APP_DB_NAME"))
	fmt.Println(os.Getenv("APP_DB_PASSWORD"))
    */

	a.Run(":2130")
}
