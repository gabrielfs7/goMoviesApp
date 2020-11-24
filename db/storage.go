package db

import (
	"github.com/kr/pretty"
)

type Video struct {
	Title    string
	Director string
	Year     int
}

type DB []Video

func Init() DB {
	return DB{}
}

func (d *DB) Add(v Video) {
	*d = append(*d, v)
}

func (d *DB) Table() string {
	return pretty.Sprint(d)
}
