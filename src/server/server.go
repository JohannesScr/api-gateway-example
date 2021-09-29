package server

import (
	"github.com/gorilla/mux"
	"github.com/johannesscr/api-gateway-example/src/handler"
	"net/http"
)

func NewServer() *Server {
	s := &Server{}
	s.Router = mux.NewRouter()

	// register routes
	s.routes()
	return s
}

type Server struct {
	Redis bool
	Router *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

/* ROUTING */

func (s *Server) routes() {
	s.Router.HandleFunc("/", s.prop(handler.HandleHome)).Methods("GET")
	s.Router.HandleFunc("/user/-", s.prop(handler.HandleUserGet)).Methods("GET")
}

func (s *Server) prop(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
	}
}

/* ENCODING and DECODING */

/*
causes circular imports

func (s *Server) Respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println("encoding error", err)
			http.Error(w, "encoding error", 500)
		}
	}
}

func (s *Server) Decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) QueryParams(w http.ResponseWriter, r *http.Request) (map[string][]string, error) {
	qp, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println("queryParams", err)
	}
	return qp, err
}
*/
