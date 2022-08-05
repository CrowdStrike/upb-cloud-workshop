package api

import (
	"exam-api/domain"

	"github.com/emicklei/go-restful/v3"
)

const (
	memoryRootPath = "/memory"
	httpRootPath   = "/http"
	redisRootPath  = "/redis"

	productPath = "/product"

	versionSingle = "/single"
	versionBatch  = "/batch"
)

type API struct {
	storage domain.Storage
}

func NewAPI(store domain.Storage) *API {
	return &API{
		storage: store,
	}
}

func (api *API) RegisterRoutes(ws *restful.WebService) {
	ws.Path("/store")
	ws.Route(ws.POST(memoryRootPath + productPath + versionSingle).To(api.createProductMemorySingle))
	ws.Route(ws.GET(memoryRootPath + productPath + versionSingle).To(api.getProductMemorySingle))
	ws.Route(ws.PATCH(memoryRootPath + productPath + versionSingle).To(api.updateProductMemorySingle))
	ws.Route(ws.DELETE(memoryRootPath + productPath + versionSingle).To(api.deleteProductMemorySingle))

	ws.Route(ws.POST(memoryRootPath + productPath + versionBatch).To(api.createProductMemoryBatch))
	ws.Route(ws.GET(memoryRootPath + productPath + versionBatch).To(api.getProductMemoryBatch))
	ws.Route(ws.PATCH(memoryRootPath + productPath + versionBatch).To(api.updateProductMemoryBatch))
	ws.Route(ws.DELETE(memoryRootPath + productPath + versionBatch).To(api.deleteProductMemoryBatch))

	// TODO create similar routes that use the store service. For this you will need to create
	// a http client that implemets the domain.Storage interface and add it to the api.
	// The handlers should be similar to those using memory storage

}
