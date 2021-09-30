package handler

import (
	"encoding/json"
	"github.com/johannesscr/api-gateway-example/micro/microservice"
	"github.com/johannesscr/api-gateway-example/micro/microtest"
	"github.com/johannesscr/api-gateway-example/src/includes"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	Testing the handler package for unit testing the handlers as well as
	the micro-service integrations.
*/

func TestHandleHome(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	HandleHome(rec, req)

	res := rec.Result()
	rawBody := struct {
		Message string
		Data    map[string]bool
		Errors  map[string][]string
	}{}
	err := json.NewDecoder(res.Body).Decode(&rawBody)
	if err != nil {
		t.Errorf("unable to decode res body")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %v got %v", http.StatusOK, res.StatusCode)
	}
	if !rawBody.Data["alive"] {
		t.Errorf("expected %v got %v", true, rawBody.Data["alive"])
	}
}

func TestHandleUserGet(t *testing.T) {
	s := microservice.NewService()
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

	e := microtest.Exchange{
		Response: microtest.Response{
			Status: 200,
			Header: map[string]string{
				"x-token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.",
			},
			Body: `{
				"message": "user found successfully", 
				"data": {
					"user": {
						"uuid": "6a67f46e-d9de-4d63-8283-bf5a5aa1e582", 
						"first_name": "james", 
						"last_name": "bond", 
						"email": "007@mi6.co.uk"
					}
				}, 
				"errors": {}
			}`,
		},
	}
	ms.Append(e)

	err := s.SetEnv()
	if err != nil {
		t.Errorf("unable to set microservice env vars")
	}

	rec := httptest.NewRecorder()

	q := map[string]string{
		"userUuid": "6a67f46e-d9de-4d63-8283-bf5a5aa1e582",
	}
	req := microtest.NewRequest("GET", "/user/-", q, nil, nil)

	HandleUserGet(rec, req)
	res, xb := microtest.ReadRecorder(rec)

	type data struct {
		User includes.User `json:"user"`
	}
	type resp struct {
		Message string              `json:"message"`
		Data    data                `json:"data"`
		Errors  map[string][]string `json:"errors"`
	}

	r := resp{}
	err = json.Unmarshal(xb, &r)
	if err != nil {
		t.Errorf("unable to unmarshal response body")
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %v got %v", http.StatusOK, res.StatusCode)
	}
	if r.Data.User.FirstName != "james" {
		t.Errorf("expected '%v' got '%v'", "james", r.Data.User.FirstName)
	}
	if r.Data.User.LastName != "bond" {
		t.Errorf("expected '%v' got '%v'", "bond", r.Data.User.LastName)
	}
}
