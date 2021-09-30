package microtest

import (
	"github.com/johannesscr/api-gateway-example/micro/microservice"
	"testing"
)

func TestMockServer(t *testing.T) {
	s := microservice.NewService()
	ms := MockServer(s)
	defer ms.Server.Close()

	e := Exchange{
		Response: Response{
			Status: 200,
			Body: `{
				"data": {},
				"errors": {},
				"message": "Welcome to the POS api"
			}`,
		},
	}
	ms.Append(e)

	b := s.GetHome()

	if !b {
		t.Errorf("failed to create mock server")
	}
}

