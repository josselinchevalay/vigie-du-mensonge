package authorize_authed_user

import (
	"fmt"
	"vdm/core/locals"
	"vdm/core/models"
)

type Service interface {
	authorizeAuthedUser(authedUser *locals.AuthedUser) error
}

type service struct {
	repo Repository
}

func (s *service) authorizeAuthedUser(authedUser *locals.AuthedUser) error {
	roles, err := s.repo.getUserRoles(authedUser.ID)
	if err != nil {
		return fmt.Errorf("failed to get user roles: %v", err)
	}

	authedUser.Roles = make([]models.RoleName, len(roles))
	for i := range roles {
		authedUser.Roles[i] = roles[i].Name
	}

	return nil
}
