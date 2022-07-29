package gateway

import (
	"fmt"
	"net/http"

	"lab05/domain"

	"github.com/emicklei/go-restful/v3"

	log "github.com/sirupsen/logrus"
)

const booksPath = "/books"

type API struct {
	books map[int]domain.Book
}

func NewAPI() *API {
	return &API{books: make(map[int]domain.Book)}
}

func (api *API) RegisterRoutes(ws *restful.WebService) {
	ws.Path("/book-app")
	ws.Route(ws.POST(booksPath).To(api.addBook).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON).Doc("Adding a new book in the database"))
	ws.Route(ws.GET(booksPath).To(api.getBook).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON).Doc("Getting a book from database"))
}

func (api *API) addBook(req *restful.Request, resp *restful.Response) {
	book := &domain.Book{}
	err := req.ReadEntity(book)

	if err != nil {
		log.WithError(err).Error("Failed to parse book json")
		_ = resp.WriteError(http.StatusBadRequest, err)
		return
	}

	_, exists := api.books[book.GetBookHash()]
	if exists {
		log.Error("Book already exists in the database")
		_ = resp.WriteError(http.StatusConflict, fmt.Errorf("book already exists"))
		return
	}

	api.books[book.GetBookHash()] = *book
	log.Info("Book added successfully in database")
}

func (api *API) getBook(req *restful.Request, resp *restful.Response) {
	author := req.QueryParameter("author")
	title := req.QueryParameter("title")

	if author == "" {
		log.Error("Failed to read author")
		_ = resp.WriteError(http.StatusBadRequest, fmt.Errorf("book author must be provided"))
		return
	}

	if title == "" {
		log.Error("Failed to read title")
		_ = resp.WriteError(http.StatusBadRequest, fmt.Errorf("book title must be provided"))
		return
	}

	book := &domain.Book{
		Title: title,
		Author: author,
	}

	b, ok := api.books[book.GetBookHash()]
	if !ok {
		log.Error("Book not found")
		_ = resp.WriteError(http.StatusNotFound, fmt.Errorf("book not found"))
		return
	}

	err := resp.WriteAsJson(b)
	if err != nil {
		log.WithError(err).Error("Failed to write response")
		_ = resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}
