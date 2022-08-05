package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Test struct {
	id               int
	method           string
	path             string
	queryParams      func() qPar
	body             func() string
	expectedStatus   int
	expectedBodyType interface{}
}

type Param struct {
	id    string
	value string
}
type qPar []Param

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

type idBody struct {
	Id string `json:"id"`
}

type idsBody struct {
	Ids    []string      `json:"ids"`
	Errors []interface{} `json:"errors"`
}

var returnedIdBody = &idBody{}
var returnedIdsBody = &idsBody{}
var partialIDsBody = &idsBody{}
var product1 = generateProduct()
var url = "http://localhost:8080"
var products []Product

var nonExistentID = func() qPar {
	s := "invalid"
	return qPar{{"id", s}}
}

var storageType = "memory"
var singleProdPath = "/store/" + storageType + "/product/single"
var batchProdPath = "/store/" + storageType + "/product/batch"

var tests = []Test{
	{1, "POST", singleProdPath, nil, func() string { return product1 }, 200, returnedIdBody},
	{2, "POST", singleProdPath, nil, func() string { return product1 }, 409, nil},
	{3, "GET", singleProdPath, singleQueryParam, nil, 200, nil},
	{4, "GET", singleProdPath, nil, nil, 400, nil},
	{5, "PATCH", singleProdPath, nil, generateProdDiff, 200, nil},
	{6, "DELETE", singleProdPath, singleQueryParam, nil, 200, nil},
	{7, "GET", singleProdPath, singleQueryParam, nil, 404, nil},
	{8, "POST", batchProdPath, nil, func() string { return generateProducts(10) }, 200, returnedIdsBody},
	{9, "POST", batchProdPath, nil, nil, 400, nil},
	{10, "POST", batchProdPath, nil, partialProducts, 200, partialIDsBody},
	{11, "GET", batchProdPath, generateQueryParams, nil, 200, nil},
	{12, "GET", batchProdPath, nil, nil, 400, nil},
	{13, "DELETE", batchProdPath, generateQueryParams, nil, 200, nil},
	{14, "GET", batchProdPath, generateQueryParams, nil, 404, returnedIdsBody},
	{15, "PATCH", batchProdPath, nil, generateProdsDiff, 200, nil},
	{16, "PATCH", batchProdPath, nil, nil, 400, nil},
}

func main() {
	rand.Seed(time.Now().Unix())
	var passed int

	for i, t := range tests {
		if err := runTest(&t); err != nil {
			log.Printf("Error in test %d: %v", t.id, err)
		} else {
			passed++
			log.Printf("Test %d passed", i+1)
		}
	}
	log.Printf("Test results: %d/%d passed", passed, len(tests))
}

func generateWord() string {
	idx := rand.Intn(len(words))
	return words[idx]
}

func runTest(t *Test) error {
	var body io.Reader
	if t.body != nil {
		s := t.body()
		body = strings.NewReader(s)
	}
	req, err := http.NewRequest(t.method, url+t.path, body)
	if err != nil {
		log.Printf("[ERROR] Failed to create request err:%v", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	if t.queryParams != nil {
		q := req.URL.Query()
		for _, v := range t.queryParams() {
			q.Add(v.id, v.value)
		}
		req.URL.RawQuery = q.Encode()
		fmt.Println("request:", q.Encode())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[ERROR] Failed exec request err: %v", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != t.expectedStatus {
		return fmt.Errorf("got %d instead of %d status code", resp.StatusCode, t.expectedStatus)
	}
	if t.expectedBodyType != nil {
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("could not read body err:%w", err)
		}
		err = json.Unmarshal(buf, &t.expectedBodyType)
		if err != nil {
			return typeErr(t.expectedBodyType)
		}
		log.Printf("Test %d Retrieved body: %+v", t.id, t.expectedBodyType)
	}
	return nil
}

func generateProducts(count int) string {
	products = make([]Product, count)
	for i := range products {
		products[i] = Product{
			Name:         generateWord(),
			Manufacturer: generateWord(),
			Price:        rand.Intn(20),
			Stock:        rand.Intn(100),
			Tags:         []string{generateWord(), generateWord()},
		}
	}
	buf, _ := json.Marshal(products)
	return string(buf)
}

func partialProducts() string {
	partProds := products
	for i := range partProds {
		if i < len(partProds)/2 {
			continue
		}
		partProds[i] = Product{
			Name:         generateWord(),
			Manufacturer: generateWord(),
			Price:        rand.Intn(20),
			Stock:        rand.Intn(100),
			Tags:         []string{generateWord(), generateWord()},
		}
	}
	buf, _ := json.Marshal(partProds)
	return string(buf)
}

func generateQueryParams() qPar {
	ids := returnedIdsBody.Ids
	var params = qPar{}
	for _, id := range ids {
		params = append(params, Param{"ids", id})
	}
	return params
}

func singleQueryParam() qPar {
	return qPar{{"id", returnedIdBody.Id}}
}

func generateProduct() string {
	product := Product{
		Name:         generateWord(),
		Manufacturer: generateWord(),
		Price:        rand.Intn(20),
		Stock:        rand.Intn(100),
		Tags:         []string{generateWord(), generateWord()},
	}
	buf, _ := json.Marshal(product)
	return string(buf)
}

func generateProdDiff() string {
	product := ProductDiff{
		ID: returnedIdBody.Id,
		Diff: struct {
			Price int      `json:"price"`
			Stock int      `json:"stock"`
			Tags  []string `json:"tags"`
		}{
			Price: rand.Intn(10),
			Stock: rand.Intn(10),
			Tags:  []string{generateWord(), generateWord()},
		},
	}
	buf, _ := json.Marshal(product)
	return string(buf)
}

func generateProdsDiff() string {
	ids := returnedIdsBody.Ids
	prodDiffs := make([]ProductDiff, len(ids))
	for i := range prodDiffs {
		prodDiffs[i] = ProductDiff{
			ID: ids[i],
			Diff: struct {
				Price int      `json:"price"`
				Stock int      `json:"stock"`
				Tags  []string `json:"tags"`
			}{
				Price: rand.Intn(10),
				Stock: rand.Intn(10),
				Tags:  []string{generateWord(), generateWord()},
			},
		}
	}

	buf, _ := json.Marshal(prodDiffs)
	return string(buf)
}

func typeErr(expected interface{}) error {
	buf, _ := json.Marshal(expected)
	return fmt.Errorf("body does not respect structure %s", buf)
}
