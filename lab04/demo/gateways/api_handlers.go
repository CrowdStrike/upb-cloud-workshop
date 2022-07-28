package gateways

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/rest-demo/domain"
	"github.com/emicklei/go-restful/v3"
)

const echoPath = "/echo"
const usersPatch = "/users"

type API struct {
	users map[int]domain.User
}

func NewAPI() *API {
	return &API{
		users: make(map[int]domain.User),
	}
}

func (api *API) RegisterRoutes(ws *restful.WebService) {
	ws.Path("/my-app")
	ws.Route(ws.POST(echoPath).To(api.echoPOSTHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON).Doc("Writes back a json with what you gave it"))
	ws.Route(ws.GET(echoPath).To(api.echoGETHandler).Writes(restful.MIME_JSON).Doc("Writes back a json with what you gave it"))

	ws.Route(ws.GET(usersPatch).To(api.getUser).Writes(restful.MIME_JSON).Doc("Writes back a json with what you gave it"))
	ws.Route(ws.POST(usersPatch).To(api.addUser).Writes(restful.MIME_JSON).Doc("Writes back a json with what you gave it"))
	// ws.Route(ws.PATCH(usersPatch).To(api.updateUser).Writes(restful.MIME_JSON).Doc("Writes back a json with what you gave it"))
	// ws.Route(ws.DELETE(usersPatch).To(api.deleteUser).Writes(restful.MIME_JSON).Doc("Writes back a json with what you gave it"))

}

func (api *API) echoPOSTHandler(req *restful.Request, resp *restful.Response) {
	body := req.Request.Body
	if body == nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, "nil body"))
		return

	}
	defer body.Close()
	var err error
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	resp.WriteAsJson(map[string]string{"res": string(data)})
}

func (api *API) echoGETHandler(req *restful.Request, resp *restful.Response) {
	param := req.QueryParameter("echo-param")
	resp.WriteAsJson(map[string]string{
		"res": param,
	})
}

func (api *API) addUser(req *restful.Request, resp *restful.Response) {
	usr := &domain.User{}
	err := req.ReadEntity(usr)
	if err != nil {
		log.Printf("[ERROR] Failed to read user, err=%v", err)
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	_, ok := api.users[usr.GetHash()]
	if ok {
		log.Printf("[ERROR] User already exists")
		resp.WriteError(http.StatusConflict, fmt.Errorf("user already exists"))
		return
	}
	api.users[usr.GetHash()] = *usr
}

func (api *API) updateUser(req *restful.Request, resp *restful.Response) {

}

func (api *API) deleteUser(req *restful.Request, resp *restful.Response) {

}

func (api *API) getUser(req *restful.Request, resp *restful.Response) {
	name := req.QueryParameter("name")
	mail := req.QueryParameter("mail")
	if name == "" {
		log.Printf("[ERROR] Failed to read name")
		resp.WriteError(http.StatusBadRequest, fmt.Errorf("user name must be provided"))
		return
	}
	if mail == "" {
		log.Printf("[ERROR] Failed to read name")
		resp.WriteError(http.StatusBadRequest, fmt.Errorf("user mail must be provided"))
		return
	}
	user := domain.User{Name: name, Mail: mail}
	hash := user.GetHash()
	u, ok := api.users[hash]
	if !ok {
		log.Printf("[ERROR] User not found")
		resp.WriteError(http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}
	resp.WriteAsJson(u)
}
