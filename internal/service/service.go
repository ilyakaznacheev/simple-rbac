package service

import (
	"context"

	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

// Repository is a storage layer for the RBAC system.
type Repository interface {
	GetUserPermissionsByOrg(ctx context.Context, userID, organizationID string) ([]model.Permission, error)
	UpsertRole(ctx context.Context, role model.Role) error
	DeleteRole(ctx context.Context, id string) error
	GetRole(ctx context.Context, id string) (*model.Role, error)
	UpsertRoleBinding(ctx context.Context, binding model.RoleBinding) error
	DeleteRoleBinding(ctx context.Context, userID, roleID string) error
	//ListUserRoles(ctx context.Context, userID string) ([]model.Role, error)
}

// Service is a business logic layer for the RBAC system.
type Service struct {
	repo Repository
}

// New creates a new instance of the Service.
func New(repo Repository) *Service {
	return &Service{repo: repo}
}
