package service

import (
	"context"

	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

type Repository interface {
	GetUserPermissionsByOrg(ctx context.Context, userID, organizationID string) ([]model.Permission, error)
	UpsertRole(ctx context.Context, role model.Role) error
	DeleteRole(ctx context.Context, id string) error
	GetRole(ctx context.Context, id string) (*model.Role, error)
	UpsertRoleBinding(ctx context.Context, binding model.RoleBinding) error
	DeleteRoleBinding(ctx context.Context, userID, roleID string) error
	//ListUserRoles(ctx context.Context, userID string) ([]model.Role, error)
}

type Service struct {
	repo Repository
}
