package service

import (
	"context"

	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

// CheckAuthority checks if user has a permission to execute certain action within the organization.
func (s *Service) CheckAuthority(
	ctx context.Context, userID, organizationID string, permission model.Permission) (bool, error) {

	permissions, err := s.repo.GetUserPermissionsByOrg(ctx, userID, organizationID)
	if err != nil {
		return false, err
	}

	// in real-world application, we should use a more efficient data structure, probably a bite mask of permissions,
	// or even better push permission check down to database.
	//
	// there is also a room for cashing, we can easily heat it up while starting the replica and
	// invalidate/update it when some role bindings are changed.
	//
	// but this is out of scope of this example.
	for _, p := range permissions {
		if p == permission {
			return true, nil
		}
	}

	return false, nil
}
