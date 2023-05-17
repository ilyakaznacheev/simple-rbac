package inmemory

import (
	"context"
	"fmt"

	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

// userPerm is a deep tree of organization-scoped user permissions
// (user -> organization -> permissions)
type userPerm map[string]map[string]map[model.Permission]struct{}

// userBindings is a deep tree of organization-scoped user roles
// (user -> organization -> role id)
type userBindings map[string]map[string]map[string]struct{}

// userRoles is a map of user roles
type userRoles map[string]model.Role

// Storage is a simple in-memory implementation of the RBAC metadata storage.
//
// Please note, that this implementation is not fault-tolerant and not scalable.
type Storage struct {
	// perms is basically a projection of the role bindings for fast lookup
	perms userPerm
	// bindings contains the actual role bindings
	bindings userBindings
	// roles contains the roles existing in the systems
	roles userRoles
}

// New creates a new instance of the Storage.
func New() *Storage {
	return &Storage{
		perms:    make(userPerm),
		bindings: make(userBindings),
		roles:    make(userRoles),
	}
}

func (s *Storage) GetUserPermissionsByOrg(_ context.Context, userID, organizationID string) ([]model.Permission, error) {
	scope, ok := s.perms[userID]
	if !ok {
		return nil, model.ErrNotFound
	}

	permissions, ok := scope[organizationID]
	if !ok {
		return nil, model.ErrNotFound
	}

	return toSlice(permissions), nil
}

func (s *Storage) UpsertRole(_ context.Context, role model.Role) error {
	s.roles[role.ID] = role
	s.refresh()
	return nil
}

func (s *Storage) DeleteRole(_ context.Context, id string) error {
	// ensure role is not assigned to any user
	for userId, scope := range s.bindings {
		for org, roles := range scope {
			for roleID := range roles {
				if roleID == id {
					return fmt.Errorf("%w: is assigned to user %s in organization %s",
						model.ErrRoleInUse, userId, org)
				}
			}
		}
	}
	// delete the role
	delete(s.roles, id)
	return nil
}

func (s *Storage) GetRole(_ context.Context, id string) (*model.Role, error) {
	r, ok := s.roles[id]
	if !ok {
		return nil, model.ErrNotFound
	}
	return &r, nil
}

func (s *Storage) UpsertRoleBinding(_ context.Context, binding model.RoleBinding) error {

	// assign the role to the user
	if _, ok := s.bindings[binding.UserID]; !ok {
		s.bindings[binding.UserID] = make(map[string]map[string]struct{})
	}

	for _, org := range binding.Scope.OrganizationIDs {
		if _, ok := s.bindings[binding.UserID][org]; !ok {
			s.bindings[binding.UserID][org] = make(map[string]struct{})
		}
		s.bindings[binding.UserID][org][binding.RoleID] = struct{}{}
	}

	s.refresh()
	return nil
}

func (s *Storage) DeleteRoleBinding(_ context.Context, userID, roleID string) error {
	if _, ok := s.bindings[userID]; !ok {
		return model.ErrNotFound
	}
	for _, roles := range s.bindings[userID] {
		delete(roles, roleID)
	}
	s.refresh()
	return nil
}

// refresh rebuilds the permissions projection
func (s *Storage) refresh() {
	s.perms = make(userPerm)
	for userID, scope := range s.bindings {
		for org, roles := range scope {
			for roleID := range roles {
				role, ok := s.roles[roleID]
				if !ok {
					continue
				}
				permissions, ok := s.perms[userID]
				if !ok {
					permissions = make(map[string]map[model.Permission]struct{})
				}
				permissions[org] = toMap(role.Permissions)
				s.perms[userID] = permissions
			}
		}
	}
}

func toSlice(p map[model.Permission]struct{}) []model.Permission {
	res := make([]model.Permission, 0, len(p))
	for perm := range p {
		res = append(res, perm)
	}
	return res
}

func toMap(p []model.Permission) map[model.Permission]struct{} {
	res := make(map[model.Permission]struct{}, len(p))
	for _, perm := range p {
		res[perm] = struct{}{}
	}
	return res
}
