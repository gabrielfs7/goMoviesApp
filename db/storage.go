package db

type Video struct {
	Title    string
	Director string
	Year     int
}

type DB []Video

func Init() DB {
	return DB{}
}

func (d DB) Add(v Video) {
	d = append(d, v)
}
