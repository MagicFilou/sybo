package user

import (
	"sybo/models/user"

	"github.com/google/uuid"
)

func New(u *user.User) error {

	u.ID = uuid.New().String()

	return u.New()
}
