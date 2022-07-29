package service

import (
	"lab05/gateway"
	"net/http"
	"os"

	"github.com/emicklei/go-restful/v3"

	log "github.com/sirupsen/logrus"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) StartWebService() {
	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{
		FullTimestamp: true,
	}}
	formatter.Line = true
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	ws := new(restful.WebService)
	restful.Add(ws)

	api := gateway.NewAPI(gateway.NewFileRepo())
	api.RegisterRoutes(ws)

	log.Info("Started serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
