package domain

import "hash/fnv"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}

func (b *Book) GetBookHash() int {
	hash := fnv.New32a()
	_, err := hash.Write([]byte(b.Title + b.Author))
	if err != nil {
		return 0
	}
	return int(hash.Sum32())
}
