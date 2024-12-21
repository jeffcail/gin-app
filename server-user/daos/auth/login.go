package auth

import (
	"github.com/mazezen/ginframe/server-user/models"
)

func FindUserByUsername(username string) (*models.User, error) {
	m := new(models.User)
	admin, err := m.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return admin, nil
}
