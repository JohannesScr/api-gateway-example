package includes

import (
	"github.com/google/uuid"
	"github.com/johannesscr/api-gateway-example/micro/microservice"
	"testing"
)

func TestUser_Map(t *testing.T) {
	mu := microservice.User{
		UUID: uuid.New(),
		FirstName: "james",
		LastName: "bond",
		Email: "james@bond.com",
	}

	u := User{}
	array16byte := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if u.UUID != array16byte {
		t.Errorf("expected initialisaztion '%v' got '%v'", array16byte, u.UUID)
	}
	if u.FirstName != "" {
		t.Errorf("expected initialisaztion '%v' got ''", u.FirstName)
	}
	if u.LastName != "" {
		t.Errorf("expected initialisaztion '%v' got ''", u.FirstName)
	}
	if u.Email != "" {
		t.Errorf("expected initialisaztion '%v' got ''", u.FirstName)
	}

	u.Map(mu)
	if u.UUID != mu.UUID {
		t.Errorf("expected initialisaztion '%v' got '%v'", mu.UUID, u.UUID)
	}
	if u.FirstName != mu.FirstName {
		t.Errorf("expected initialisaztion '%v' got '%v'", mu.FirstName, u.FirstName)
	}
	if u.LastName != mu.LastName {
		t.Errorf("expected initialisaztion '%v' got '%v'", mu.LastName ,u.LastName)
	}
	if u.Email != mu.Email {
		t.Errorf("expected initialisaztion '%v' got '%v'", mu.Email, u.Email)
	}
}
