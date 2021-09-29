package includes

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/* DOTENV */

type Env struct {
	Vars map[string]string
}

// Load takes a filename s and loads the contents of the file as environmental
// variables
func (e Env) Load(s string) {
	if e.Vars == nil {
		e.Vars = make(map[string]string)
	}

	bs, _ := os.ReadFile(s)
	rows := strings.Split(string(bs), "\n")
	for _, ln := range rows {
		if ln == "" {
			continue
		}

		xln := strings.Split(ln, "=")
		// set struct map
		e.Vars[xln[0]] = xln[1]
		// set env var
		err := os.Setenv(xln[0], xln[1])
		if err != nil {
			log.Panic(err.Error())
		}
	}
}

/* ENCODING and DECODING */

// Resp is a structure for the basis of all responses to clients.
// This enforces a consistent structure for all responses.
type Resp struct {
	Status    int                `json:"-"`
	Message string              `json:"message"`
	Data    interface{}         `json:"data"`
	Errors  map[string][]string `json:"errors"`
}

// Respond encodes to JSON and writes the response structure to the client
func (resp *Resp) Respond(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(resp.Status)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println("Respond encoding error", err)
		http.Error(w, "encoding error", 500)
	}
}

// Decode decodes the http.Request body to the value pointed to by v.
func Decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// QueryParams parses the URL query parameters to a map
func QueryParams(w http.ResponseWriter, r *http.Request) (map[string][]string, error) {
	qp, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println("queryParams", err)
	}
	return qp, err
}
