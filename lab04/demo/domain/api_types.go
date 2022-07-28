package domain

import (
	"hash/fnv"
)

type User struct {
	Name      string   `json:"name"`
	Mail      string   `json:"mail"`
	Age       int      `json:"age"`
	Interests []string `json:"interests"`
}

func (u *User) GetHash() int {
	h := fnv.New32a()
	h.Write([]byte(u.Name + u.Mail))
	return int(h.Sum32())
}
