package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/emicklei/go-restful/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var api *API

func newJSONRequest(method, url string, obj interface{}) *http.Request {
	var payloadReader io.Reader

	if obj != nil {
		payload, err := json.Marshal(obj)
		Expect(err).To(BeNil())
		payloadReader = bytes.NewBuffer(payload)
	}
	req, err := http.NewRequest(method, url, payloadReader)
	Expect(err).To(BeNil())
	req.Header.Set("content-type", "application/json")
	return req
}

func readJSONResponse(in io.Reader, v interface{}) {
	respBody, err := ioutil.ReadAll(in)
	Expect(err).To(BeNil())
	fmt.Println(string(respBody))
	Expect(json.Unmarshal(respBody, v)).To(BeNil())
}

var _ = Describe("API Handler", func() {
	var (
		ws        *restful.WebService
		container *restful.Container
		recorder  *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		api = NewAPI()
		ws = new(restful.WebService)
		container = restful.NewContainer()
		container.Add(ws)
		recorder = httptest.NewRecorder()
	})

	Context("When adding a new book", func() {
		It("should fail as no author is mentioned in query parameters", func() {
			request, _ := http.NewRequest("GET", "/book-app/books", nil)

			container.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
		})
	})

	Context("When getting a book", func() {

	})
})