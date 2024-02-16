package main

import (
	"fmt"
	"os"
)

func main() {
	dbHostname := os.Getenv("DB_HOSTNAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	port := os.Getenv("PORT")

	fmt.Println("hello, world!")
	fmt.Println("hostname: " + dbHostname)
	fmt.Println("username: " + dbUsername)
	fmt.Println("password: " + dbPassword)
	fmt.Println("port: " + port)

}
