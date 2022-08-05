package domain

import (
	"crypto/sha1"
	"encoding/hex"
)

// Product is the API representation of a product object
type Product struct {
	Name         string   `json:"name"`
	Manufacturer string   `json:"manufacturer"`
	Price        int      `json:"price"`
	Stock        int      `json:"stock"`
	Tags         []string `json:"tags"`
}

type ProductDiff struct {
	ID   string `json:"id"`
	Diff struct {
		Price int      `json:"price"`
		Stock int      `json:"stock"`
		Tags  []string `json:"tags"`
	} `json:"diff"`
}

// GetHash returns a sha1 value over the name and manufacturer fields of a Product
func (p *Product) GetHash() string {
	h := sha1.New()
	h.Write([]byte(p.Name + p.Manufacturer))
	return hex.EncodeToString(h.Sum(nil))
}
