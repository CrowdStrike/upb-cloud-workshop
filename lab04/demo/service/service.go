package service

import (
	"log"
	"net/http"

	"example.com/rest-demo/gateways"
	"github.com/emicklei/go-restful/v3"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) StartWebService() {
	ws := new(restful.WebService)
	restful.Add(ws)

	api := gateways.NewAPI()
	api.RegisterRoutes(ws)

	log.Printf("Started serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
