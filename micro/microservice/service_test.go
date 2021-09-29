package microservice

import (
	"github.com/google/uuid"
	"github.com/johannesscr/api-gateway-example/micro/microtest"
	"testing"
)

func TestNewService(t *testing.T) {
	s := NewService()
	if s == nil {
		t.Errorf("expected %v got %v", service{}, nil)
	}
}

func TestService_SetURL(t *testing.T) {
	s := NewService()
	if s.host != "" {
		t.Errorf("expected '%v' got '%v'", "", s.host)
	}
	sc := "http"
	h := "test.domain.com"
	p := "/user"

	s.SetURL(sc, h)
	s.URL.Path = p

	if s.host != h {
		t.Errorf("expected '%v' got '%v'", h, s.host)
	}
	if s.scheme != sc {
		t.Errorf("expected '%v' got '%v'", sc, s.scheme)
	}
	if s.URL.String() != "http://test.domain.com/user" {
		t.Errorf("expected '%v' got '%v'", "http://test.domain.com/user", s.URL.String())
	}
}

func TestService_GetHome(t *testing.T) {
	//s := NewService()
	//srv := microtest.MockServer(s)
}

func TestService_GetUser(t *testing.T) {
	// create a microservice instance
	s := NewService()
	// startup the microservice
	ms := microtest.MockServer(s)
	defer ms.Close()  // defer shut down the microservice

	u, errors := s.GetUser(uuid.New().String())

	if errors != nil {
		t.Errorf("errors on response: %v", errors)
	}
	if u.FirstName != "james" {
		t.Errorf("expected '%v' got '%v'", "james", u.FirstName)
	}
	if u.LastName != "bond" {
		t.Errorf("expected '%v' got '%v'", "bond", u.LastName)
	}
}