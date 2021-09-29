package handler

import (
	"encoding/json"
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
	rawBody := struct{
		Message string
		Data map[string]bool
		Errors map[string][]string
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

