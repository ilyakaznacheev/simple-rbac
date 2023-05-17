package inmemory

import (
	"context"
	"reflect"
	"testing"

	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

const (
	user1 = "user-1-JVZqoYp3"
	user2 = "user-2-d2C1Z5k4"
	user3 = "user-3-40Yt4aTEj"
	role1 = "role-1-W0c8MNSiW"
	role2 = "role-2-muu6NWO6Fa"
	role3 = "role-3-QVbKsoZGs6"
	role4 = "role-4-FlAUHROR"
	role5 = "role-5-UuQ85NaVJ"
	org1  = "org-1-AiOFiFzt4C"
	org2  = "org-2-3m7x79Xhf5"
	org3  = "org-3-x67WaBMDf"
)

type fields struct {
	perms    userPerm
	bindings userBindings
	roles    userRoles
}

func initFields() fields {
	return fields{
		perms: map[string]map[string]map[model.Permission]struct{}{
			user1: {
				org1: {
					model.PermCreateProject: {},
					model.PermDeleteProject: {},
				},
				org2: {
					model.PermDeployProject: {},
				},
			},
			user2: {
				org1: {
					model.PermReadLogs:   {},
					model.PermModifyLogs: {},
				},
				org2: {
					model.PermDeployProject: {},
				},
			},
		},
		bindings: map[string]map[string]map[string]struct{}{
			user1: {
				org1: {
					role1: {},
				},
				org2: {
					role3: {},
				},
			},
			user2: {
				org1: {
					role2: {},
				},
				org2: {
					role3: {},
				},
			},
		},
		roles: map[string]model.Role{
			role1: {
				ID:          role1,
				Permissions: []model.Permission{model.PermCreateProject, model.PermDeleteProject},
			},
			role2: {
				ID:          role2,
				Permissions: []model.Permission{model.PermDeployProject},
			},
			role3: {
				ID:          role3,
				Permissions: []model.Permission{model.PermReadLogs, model.PermModifyLogs},
			},
			role4: {
				ID:          role4,
				Permissions: []model.Permission{model.PermAuditLogs},
			},
		},
	}
}

func TestStorage_DeleteRole(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "unused role",
			fields: initFields(),
			args: args{
				id: role4,
			},
			wantErr: false,
		},
		{
			name:   "used role",
			fields: initFields(),
			args: args{
				id: role1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				perms:    tt.fields.perms,
				bindings: tt.fields.bindings,
				roles:    tt.fields.roles,
			}
			if err := s.DeleteRole(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if _, ok := s.roles[tt.args.id]; ok && !tt.wantErr {
				t.Errorf("DeleteRole() role %s is not deleted", tt.args.id)
			}

		})
	}
}

func TestStorage_DeleteRoleBinding(t *testing.T) {
	type args struct {
		userID string
		roleID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "delete role 1",
			fields: initFields(),
			args: args{
				userID: user1,
				roleID: role1,
			},
			wantErr: false,
		},
		{
			name:   "delete role 2",
			fields: initFields(),
			args: args{
				userID: user2,
				roleID: role2,
			},
			wantErr: false,
		},
		{
			name:   "delete role of nonexistent user",
			fields: initFields(),
			args: args{
				userID: user3,
				roleID: role3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				perms:    tt.fields.perms,
				bindings: tt.fields.bindings,
				roles:    tt.fields.roles,
			}
			if err := s.DeleteRoleBinding(context.Background(), tt.args.userID, tt.args.roleID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRoleBinding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if _, ok := s.bindings[tt.args.userID][tt.args.roleID]; ok {
				t.Errorf("DeleteRoleBinding() binding %s %s is not deleted", tt.args.userID, tt.args.roleID)
			}
		})
	}
}

func TestStorage_GetRole(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Role
		wantErr bool
	}{
		{
			name:   "existing role",
			fields: initFields(),
			args: args{
				id: role1,
			},
			want: &model.Role{
				ID:          role1,
				Permissions: []model.Permission{model.PermCreateProject, model.PermDeleteProject},
			},
			wantErr: false,
		},
		{
			name:   "non-existing role",
			fields: initFields(),
			args: args{
				id: role5,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				perms:    tt.fields.perms,
				bindings: tt.fields.bindings,
				roles:    tt.fields.roles,
			}
			got, err := s.GetRole(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetUserPermissionsByOrg(t *testing.T) {

	type args struct {
		userID         string
		organizationID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Permission
		wantErr bool
	}{
		{
			name:   "user 1 org 1",
			fields: initFields(),
			args: args{
				userID:         user1,
				organizationID: org1,
			},
			want:    []model.Permission{model.PermCreateProject, model.PermDeleteProject},
			wantErr: false,
		},
		{
			name:   "user 1 org 2",
			fields: initFields(),
			args: args{
				userID:         user1,
				organizationID: org2,
			},
			want:    []model.Permission{model.PermDeployProject},
			wantErr: false,
		},
		{
			name:   "user 1 org 3",
			fields: initFields(),
			args: args{
				userID:         user1,
				organizationID: org3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				perms:    tt.fields.perms,
				bindings: tt.fields.bindings,
				roles:    tt.fields.roles,
			}
			got, err := s.GetUserPermissionsByOrg(context.Background(), tt.args.userID, tt.args.organizationID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserPermissionsByOrg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserPermissionsByOrg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_UpsertRole(t *testing.T) {
	type args struct {
		role model.Role
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "existing role",
			fields: initFields(),
			args: args{
				role: model.Role{
					ID:          role1,
					Permissions: []model.Permission{model.PermCreateProject, model.PermDeleteProject, model.PermDeployProject},
				},
			},
			wantErr: false,
		},
		{
			name:   "new role",
			fields: initFields(),
			args: args{
				role: model.Role{
					ID:          role5,
					Permissions: []model.Permission{model.PermManageEnvironments},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				perms:    tt.fields.perms,
				bindings: tt.fields.bindings,
				roles:    tt.fields.roles,
			}
			if err := s.UpsertRole(context.Background(), tt.args.role); (err != nil) != tt.wantErr {
				t.Errorf("UpsertRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(s.roles[tt.args.role.ID], tt.args.role) {
				t.Errorf("UpsertRole() got = %v, want %v", s.roles[tt.args.role.ID], tt.args.role)
			}
		})
	}
}

func TestStorage_UpsertRoleBinding(t *testing.T) {
	type args struct {
		binding model.RoleBinding
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "new role new org",
			fields: initFields(),
			args: args{
				binding: model.RoleBinding{
					UserID: user1,
					RoleID: role4,
					Scope: model.Scope{
						OrganizationIDs: []string{org3},
					},
				},
			},
			wantErr: false,
		},
		{
			name:   "new role existing org",
			fields: initFields(),
			args: args{
				binding: model.RoleBinding{
					UserID: user1,
					RoleID: role4,
					Scope: model.Scope{
						OrganizationIDs: []string{org1},
					},
				},
			},
			wantErr: false,
		},
		{
			name:   "new role several orgs",
			fields: initFields(),
			args: args{
				binding: model.RoleBinding{
					UserID: user1,
					RoleID: role4,
					Scope: model.Scope{
						OrganizationIDs: []string{org1, org2, org3},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				perms:    tt.fields.perms,
				bindings: tt.fields.bindings,
				roles:    tt.fields.roles,
			}
			if err := s.UpsertRoleBinding(context.Background(), tt.args.binding); (err != nil) != tt.wantErr {
				t.Errorf("UpsertRoleBinding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// check if binding for organisation has role
			for _, org := range tt.args.binding.Scope.OrganizationIDs {
				if _, ok := s.bindings[tt.args.binding.UserID][org][tt.args.binding.RoleID]; !ok {
					t.Errorf("UpsertRoleBinding() got = %v, want %v", s.bindings[org][tt.args.binding.UserID][tt.args.binding.RoleID], tt.args.binding.RoleID)
				}
			}

		})
	}
}
