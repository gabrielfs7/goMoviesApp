package main

import (
	"log"

	"github.com/gabrielfs7/goMoviesApp/db"
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
