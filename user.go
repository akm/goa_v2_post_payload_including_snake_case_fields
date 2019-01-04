package snakecasefields

import (
	"context"
	"log"

	user "github.com/akm/goa_v2_post_payload_including_snake_case_fields/gen/user"
)

// user service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{logger}
}

// Create implements create.
func (s *usersrvc) Create(ctx context.Context, p *user.UserPayload) (res *user.User, err error) {
	res = &user.User{}
	s.logger.Print("user.create")
	res = &user.User{
		FirstName: p.FirstName,
	}
	if p.LastName != nil {
		res.LastName = *p.LastName
	}
	return
}
