package service

import (
	"net/http"

	"lab05/gateway"

	"github.com/emicklei/go-restful/v3"

	log "github.com/sirupsen/logrus"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) StartWebService() {
	ws := new(restful.WebService)
	restful.Add(ws)

	api := gateway.NewAPI()
	api.RegisterRoutes(ws)

	log.Info("Started serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
