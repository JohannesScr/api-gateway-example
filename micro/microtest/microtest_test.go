package microtest

import (
	"github.com/johannesscr/api-gateway-example/micro/microservice"
	"testing"
)

func TestMockServer(t *testing.T) {
	s := microservice.NewService()
	ms := MockServer(s)
	defer ms.Server.Close()

	ms.Response.Status = 200
	ms.Response.Body = `{"data": {},"errors": {},"message": "Welcome to the POS api"}`

	b := s.GetHome()

	if !b {
		t.Errorf("failed to create mock server")
	}
}

