package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

// CreateRole creates a new role
func (s *Service) CreateRole(ctx context.Context, role model.Role) error {
	if err := s.repo.UpsertRole(ctx, role); err != nil {
		zap.L().Error("failed to create role", zap.Error(err))
		return err
	}
	return nil
}

// UpdateRole updates an existing role
func (s *Service) UpdateRole(ctx context.Context, role model.Role) error {
	if err := s.repo.UpsertRole(ctx, role); err != nil {
		zap.L().Error("failed to update role", zap.Error(err))
		return err
	}
	return nil
}

// GetRole returns a role by id
func (s *Service) GetRole(ctx context.Context, id string) (*model.Role, error) {
	r, err := s.repo.GetRole(ctx, id)
	if err != nil {
		zap.L().Error("failed to get role", zap.Error(err))
		return nil, err
	}
	return r, nil
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
