package api

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ilyakaznacheev/simple-rbac/api/auth-go"
	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

var _ auth.AuthorityServiceServer = (*Server)(nil)
var _ auth.RoleServiceServer = (*Server)(nil)

func (s *Server) AuthorizeAction(ctx context.Context, action *auth.Action) (*empty.Empty, error) {
	p, err := mapApiPermission(action.Permission)
	if err != nil {
		return nil, err
	}

	err = s.app.CheckAuthority(ctx, action.UserId, action.OrgId, p)

	return nil, mapError(err)
}

func (s *Server) GetRole(ctx context.Context, request *auth.GetRoleRequest) (*auth.Role, error) {
	r, err := s.app.GetRole(ctx, request.Id)
	if err != nil {
		return nil, mapError(err)
	}

	perm := make([]auth.Permission, len(r.Permissions))
	for i, p := range r.Permissions {
		perm[i] = mapModelPermission(p)
	}

	return &auth.Role{
		RoleId:      r.ID,
		Permissions: perm,
	}, nil
}

func (s *Server) CreateRole(ctx context.Context, role *auth.Role) (*auth.Role, error) {
	r := model.Role{
		ID:          role.RoleId,
		Permissions: make([]model.Permission, 0, len(role.Permissions)),
	}

	for _, p := range role.Permissions {
		pp, err := mapApiPermission(p)
		if err != nil {
			return nil, err
		}
		r.Permissions = append(r.Permissions, pp)
	}

	err := s.app.CreateRole(ctx, r)
	if err != nil {
		return nil, mapError(err)
	}

	return role, nil
}

func (s *Server) UpdateRole(ctx context.Context, role *auth.Role) (*auth.Role, error) {
	r := model.Role{
		ID:          role.RoleId,
		Permissions: make([]model.Permission, 0, len(role.Permissions)),
	}

	for _, p := range role.Permissions {
		pp, err := mapApiPermission(p)
		if err != nil {
			return nil, err
		}
		r.Permissions = append(r.Permissions, pp)
	}

	err := s.app.UpdateRole(ctx, r)
	if err != nil {
		return nil, mapError(err)
	}

	return role, nil
}

func (s *Server) DeleteRole(ctx context.Context, request *auth.DeleteRoleRequest) (*empty.Empty, error) {
	err := s.app.DeleteRole(ctx, request.Id)
	if err != nil {
		return nil, mapError(err)
	}

	return &empty.Empty{}, nil
}

func (s *Server) CreateRoleBinding(ctx context.Context, req *auth.CreateRoleBindingRequest) (*empty.Empty, error) {
	err := s.app.AssignRole(ctx, req.UserId, req.RoleId, req.OrgId)
	if err != nil {
		return nil, mapError(err)
	}

	return nil, nil
}

func (s *Server) DeleteRoleBinding(ctx context.Context, req *auth.DeleteRoleBindingRequest) (*empty.Empty, error) {
	err := s.app.RevokeRole(ctx, req.UserId, req.RoleId)
	if err != nil {
		return nil, mapError(err)
	}

	return nil, nil
}

func mapError(err error) error {
	switch {
	case errors.Is(err, model.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, model.ErrUnauthorized):
		return status.Error(codes.PermissionDenied, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}

func mapApiPermission(p auth.Permission) (r model.Permission, err error) {
	switch p {
	case auth.Permission_PERM_MANAGE_USERS:
		r = model.PermManageUsers
	case auth.Permission_PERM_MODIFY_USER_PERMISSIONS:
		r = model.PermModifyUserPermissions
	case auth.Permission_PERM_CREATE_PROJECT:
		r = model.PermCreateProject
	case auth.Permission_PERM_DELETE_PROJECT:
		r = model.PermDeleteProject
	case auth.Permission_PERM_DEPLOY_PROJECT:
		r = model.PermDeployProject
	case auth.Permission_PERM_MANAGE_ENVIRONMENTS:
		r = model.PermManageEnvironments
	case auth.Permission_PERM_READ_LOGS:
		r = model.PermReadLogs
	case auth.Permission_PERM_MODIFY_LOGS:
		r = model.PermModifyLogs
	case auth.Permission_PERM_AUDIT_LOGS:
		r = model.PermAuditLogs
	default:
		err = status.Error(codes.InvalidArgument, "unknown permission")
	}
	return
}

func mapModelPermission(p model.Permission) (r auth.Permission) {
	switch p {
	case model.PermManageUsers:
		return auth.Permission_PERM_MANAGE_USERS
	case model.PermModifyUserPermissions:
		return auth.Permission_PERM_MODIFY_USER_PERMISSIONS
	case model.PermCreateProject:
		return auth.Permission_PERM_CREATE_PROJECT
	case model.PermDeleteProject:
		return auth.Permission_PERM_DELETE_PROJECT
	case model.PermDeployProject:
		return auth.Permission_PERM_DEPLOY_PROJECT
	case model.PermManageEnvironments:
		return auth.Permission_PERM_MANAGE_ENVIRONMENTS
	case model.PermReadLogs:
		return auth.Permission_PERM_READ_LOGS
	case model.PermModifyLogs:
		return auth.Permission_PERM_MODIFY_LOGS
	case model.PermAuditLogs:
		return auth.Permission_PERM_AUDIT_LOGS
	default:
		return auth.Permission_PERM_NOOP

	}
}
