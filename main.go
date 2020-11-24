package main

import (
	"db"
	"log"
)

func main() {
	v1 := Video{
		Title:    "Inception",
		Director: "Christofer Nolan",
		Year:     2010,
	}

	myDb := db.Init()
	myDb.Add(v1)

	log.Print(myDb)
}
