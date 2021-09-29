package microtest

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

// mock is the interface that connects all micro-services
type mock interface {
	SetURL(scheme string, host string)
}

type Response struct {
	Status int
	Header map[string]string
	Body string
}

type Mock struct {
	URL    url.URL
	Server *httptest.Server
	Response Response
}

func MockServer(mx mock) *Mock {
	m := &Mock{}
	m.Response.Header = make(map[string]string)
	m.Server = m.mockServer(mx)
	return m
}

// SetURL makes the Mock also of type mock interface
func (m *Mock) SetURL(s string, h string) {
	m.URL.Scheme = s
	m.URL.Host = h
}

// MockServer takes a type mock interface, the type mock interface is the
// interface for any micro-service. Due to go routing any request to the mock
// handler the type mock interface which points (via the URL) to the
// MockServer can return any response provided for any request make to the
// type mock interface
func (m *Mock) mockServer(mx mock) *httptest.Server {
	mockServer := httptest.NewServer(m.mockHandler())

	xs := strings.Split(mockServer.URL, "/")
	scheme := strings.Replace(xs[0], ":", "", 1)
	host := strings.Join(xs[2:], "")
	mx.SetURL(scheme, host)
	return mockServer
}

// mockHandler takes the request properties defined on the Mock and writes
// it to the response of the mockServer which is a mock of the micro-service
// being tested
func (m *Mock) mockHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//log.Println(m.Response.Status)
		//log.Println(m.Response.Header)
		//log.Println(m.Response.Body)

		w.WriteHeader(m.Response.Status)
		for key, val := range m.Response.Header {
			w.Header().Set(key, val)
		}

		_, err := w.Write([]byte(m.Response.Body))
		if err != nil {
			log.Panic(err)
		}
	}
}

