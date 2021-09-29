package microtest

import (
	"github.com/google/uuid"
	"github.com/johannesscr/api-gateway-example/src/includes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

type mock interface {
	SetURL(s string, h string)
}

//type MockServer struct {
//	URL url.URL
//}

//func mockHandler(w http.ResponseWriter, r *http.Request) {
//	type user struct {
//		UUID uuid.UUID `json:"uuid"`
//		FirstName string `json:"first_name"`
//		LastName string `json:"last_name"`
//		Email string `json:"email"`
//	}
//
//	u := user{
//		UUID: uuid.New(),
//		FirstName: "james",
//		LastName: "bond",
//		Email: "007@gov.uk",
//	}
//
//	resp := includes.Resp{
//		Status: 200,
//		Message: "user found",
//		Data: map[string]user{
//			"user": u,
//		},
//	}
//	resp.Respond(w, r)
//}

func MockHandle(chReq chan<- string, chRes <-chan string, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bs := []byte(<-chRes)



		resp := includes.Resp{
			Status: 200,
			Message: "user found",
			Data: map[string]user{
				"user": u,
			},
		}
		resp.Respond(w, r)
	}
}

func MockServer(m mock) *httptest.Server {
	mockServer := httptest.NewServer(http.HandlerFunc(mockHandler))

	xs := strings.Split(mockServer.URL, "/")
	scheme := strings.Replace(xs[0], ":", "", 1)
	host := strings.Join(xs[2:], "")
	m.SetURL(scheme, host)
	return mockServer
}

//func (m MockServer) MockHandler() {
//
//}

