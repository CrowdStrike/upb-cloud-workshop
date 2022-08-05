package service

import (
	"exam-store/api"
	"exam-store/gateways/sql"
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

	db, err := sql.CreatePostgresConnection(
		"0.0.0.0:5432", "",
		"upb",
		"upb",
		"disable")
	if err != nil {
		log.Errorf("Failed creating connection=%+v", err)
	}

	storage := sql.NewProductRepository(db)

	apiManager := api.NewAPI(storage)
	apiManager.RegisterRoutes(ws)

	log.Printf("Started store service on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
