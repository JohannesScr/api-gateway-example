package microtest

import (
	"io"
	"io/ioutil"
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

type Exchange struct {
	Response Response
	Request *http.Request
}

type Mock struct {
	URL    url.URL
	Server *httptest.Server
	Exchanges []Exchange
	transmission int

}

func MockServer(mx mock) *Mock {
	m := &Mock{
		transmission: 0,
	}
	//m.Response.Header = make(map[string]string)
	m.Server = m.mockServer(mx)
	return m
}

// Append adds an Exchange to the queue (Q) of exchanges between the
// api-gateway and the micro-service. Exchanges in the Q are processed a
// First-In-First-Out (FIFO) manner.
func (m *Mock) Append(e Exchange) {
	m.Exchanges = append(m.Exchanges, e)
}

// transmit mocks the action where the micro-service receives the
func (m *Mock) transmit(r *http.Request) Response {
	if m.transmission == len(m.Exchanges) {
		log.Panic("exceeded mock request/response exchange transmissions")
	}

	e := m.Exchanges[m.transmission]
	e.Request = r
	// increment transmission number
	m.transmission++
	return e.Response
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
		res := m.transmit(r)

		w.WriteHeader(res.Status)
		for key, val := range res.Header {
			w.Header().Set(key, val)
		}

		_, err := w.Write([]byte(res.Body))
		if err != nil {
			log.Panic(err)
		}
	}
}

// ReadRecorder reads the recorder to get the response and decodes the body
// to a slice of bytes.
func ReadRecorder(rec *httptest.ResponseRecorder) (*http.Response, []byte) {
	res := rec.Result()
	xb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}
	err = res.Body.Close()
	if err != nil {
		log.Panic(err)
	}
	return res, xb
}

// NewRequest is based on
func NewRequest(method string, target string, query map[string]string, headers map[string][]string, body io.Reader) *http.Request {
	// new request
	r := httptest.NewRequest(method, target, body)
	// set headers
	r.Header = headers
	// set query params
	q := url.Values{}
	for key, val := range query {
		q.Add(key, val)
	}
	r.URL.RawQuery = q.Encode()
	return r
}
