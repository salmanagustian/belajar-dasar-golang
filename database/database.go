package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init running..")
	connection = "Postgres"
}

func GetDatabase() string {
	return connection
}
