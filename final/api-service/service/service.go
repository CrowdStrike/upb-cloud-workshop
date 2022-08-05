package service

import (
	"exam-api/gateways/api"
	"exam-api/gateways/memory"
	"net/http"
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/emicklei/go-restful/v3"
	log "github.com/sirupsen/logrus"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) StartWebService() {
	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}}
	formatter.Line = true
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	ws := new(restful.WebService)
	restful.Add(ws)

	storage := memory.NewStore()

	apiManager := api.NewAPI(storage)
	apiManager.RegisterRoutes(ws)

	log.Printf("Started api service on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
