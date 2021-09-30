package includes

import (
	"github.com/google/uuid"
	"github.com/johannesscr/api-gateway-example/micro/microservice"
)

type User struct {
	UUID      uuid.UUID `json:"uuid"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
}

func (u *User) Map(mu microservice.User) {
	u.UUID = mu.UUID
	u.FirstName = mu.FirstName
	u.LastName = mu.LastName
	u.Email = mu.Email
}
