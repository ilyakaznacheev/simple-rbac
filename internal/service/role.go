package service

import (
	"context"

	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

// CreateRole creates a new role
func (s *Service) CreateRole(ctx context.Context, role model.Role) error {
	return s.repo.UpsertRole(ctx, role)
}

// UpdateRole updates an existing role
func (s *Service) UpdateRole(ctx context.Context, role model.Role) error {
	return s.repo.UpsertRole(ctx, role)
}

// GetRole returns a role by id
func (s *Service) GetRole(ctx context.Context, id string) (model.Role, error) {
	return s.repo.GetRole(ctx, id)
}

// DeleteRole deletes a role by id
func (s *Service) DeleteRole(ctx context.Context, id string) error {
	return s.repo.DeleteRole(ctx, id)
}

// AssignRole assigns a role to a user in an organization scope
func (s *Service) AssignRole(ctx context.Context, userID, roleID, organizationID string) error {
	rb := model.RoleBinding{
		UserID: userID,
		RoleID: roleID,
		Scope: model.Scope{
			OrganizationIDs: []string{organizationID},
		},
	}

	return s.repo.UpsertRoleBinding(ctx, rb)
}

// RevokeRole revokes a role from a user
func (s *Service) RevokeRole(ctx context.Context, userID, roleID string) error {
	return s.repo.DeleteRoleBinding(ctx, userID, roleID)
}
