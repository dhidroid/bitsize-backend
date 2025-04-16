package database

import "fmt"

func ConnectDB(dbconnect bool) bool {
	if dbconnect {
		return true
	} else {
		fmt.Println("null")
	}

	return false
}
