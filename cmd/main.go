package main

import (
	"r_keeper/db"
	"r_keeper/logger"
)

func main() {
	if err := logger.Init(); err != nil {
		panic(err)
	}

	if err := db.ConnectToDB(); err != nil {
		panic(err)
	}

	if err := db.Migrate(); err != nil {
		panic(err)
	}
}
