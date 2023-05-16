package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/ilyakaznacheev/simple-rbac/internal/mocks"
	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

func TestService_CreateRole(t *testing.T) {
	type args struct {
		role model.Role
	}
	type mockArgs struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		mock    mockArgs
		wantErr bool
	}{
		{
			name: "good creation 1",
			args: args{
				role: model.Role{
					ID:          "5210693d-0257-42c9-94b2-356c0470ba46",
					Permissions: []model.Permission{model.PermManageUsers, model.PermManageEnvironments},
				},
			},
			mock: mockArgs{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "good creation 2",
			args: args{
				role: model.Role{
					ID:          "ef8c45ee-cea5-4337-ae58-81e735ebfe5d",
					Permissions: []model.Permission{model.PermCreateProject, model.PermDeleteProject, model.PermDeployProject},
				},
			},
			mock: mockArgs{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "good creation 3",
			args: args{
				role: model.Role{
					ID:          "efa0a5e2-4bf7-46db-a0c4-d4cb50907418",
					Permissions: nil,
				},
			},
			mock: mockArgs{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "bad creation",
			args: args{
				role: model.Role{
					ID:          "9a473258-6ae3-4bd5-964c-56a8562bab29",
					Permissions: []model.Permission{model.PermManageUsers},
				},
			},
			mock: mockArgs{
				err: fmt.Errorf("some error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(mocks.Repository)
			r.On("UpsertRole", mock.Anything, tt.args.role).
				Once().Return(tt.mock.err)

			s := &Service{
				repo: r,
			}
			if err := s.CreateRole(context.Background(), tt.args.role); (err != nil) != tt.wantErr {
				t.Errorf("CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_UpdateRole(t *testing.T) {
	type args struct {
		role model.Role
	}
	type mockArgs struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		mock    mockArgs
		wantErr bool
	}{
		{
			name: "good creation 1",
			args: args{
				role: model.Role{
					ID:          "5210693d-0257-42c9-94b2-356c0470ba46",
					Permissions: []model.Permission{model.PermManageUsers, model.PermManageEnvironments},
				},
			},
			mock: mockArgs{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "good creation 2",
			args: args{
				role: model.Role{
					ID:          "ef8c45ee-cea5-4337-ae58-81e735ebfe5d",
					Permissions: []model.Permission{model.PermCreateProject, model.PermDeleteProject, model.PermDeployProject},
				},
			},
			mock: mockArgs{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "good creation 3",
			args: args{
				role: model.Role{
					ID:          "efa0a5e2-4bf7-46db-a0c4-d4cb50907418",
					Permissions: nil,
				},
			},
			mock: mockArgs{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "bad creation",
			args: args{
				role: model.Role{
					ID:          "9a473258-6ae3-4bd5-964c-56a8562bab29",
					Permissions: []model.Permission{model.PermManageUsers},
				},
			},
			mock: mockArgs{
				err: fmt.Errorf("some error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(mocks.Repository)
			r.On("UpsertRole", mock.Anything, tt.args.role).
				Once().Return(tt.mock.err)

			s := &Service{
				repo: r,
			}
			if err := s.UpdateRole(context.Background(), tt.args.role); (err != nil) != tt.wantErr {
				t.Errorf("CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_DeleteRole(t *testing.T) {
	type args struct {
		roleID string
	}
	type mockArgs struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		mock    mockArgs
		wantErr bool
	}{
		{
			name: "good deletion",
			args: args{
				roleID: "87352a21-32ae-442e-beb4-0080b68e9abd",
			},
			mock: mockArgs{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "bad deletion",
			args: args{
				roleID: "87352a21-32ae-442e-beb4-0080b68e9abd",
			},
			mock: mockArgs{
				err: fmt.Errorf("some error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(mocks.Repository)
			r.On("DeleteRole", mock.Anything, tt.args.roleID).
				Once().Return(tt.mock.err)

			s := &Service{
				repo: r,
			}
			if err := s.DeleteRole(context.Background(), tt.args.roleID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetRole(t *testing.T) {

	type args struct {
		roleID string
	}
	type mockArgs struct {
		role model.Role
		err  error
	}
	tests := []struct {
		name    string
		args    args
		mock    mockArgs
		want    model.Role
		wantErr bool
	}{
		{
			name: "good read 1",
			args: args{
				roleID: "2996aad3-1c4f-4a6e-a8da-ed3951fb6db1",
			},
			mock: mockArgs{
				role: model.Role{
					ID:          "2996aad3-1c4f-4a6e-a8da-ed3951fb6db1",
					Permissions: []model.Permission{model.PermReadLogs, model.PermModifyLogs},
				},
				err: nil,
			},
			want: model.Role{
				ID:          "2996aad3-1c4f-4a6e-a8da-ed3951fb6db1",
				Permissions: []model.Permission{model.PermReadLogs, model.PermModifyLogs},
			},
			wantErr: false,
		},
		{
			name: "good read 2",
			args: args{
				roleID: "382d211c-2a22-47c2-aa7d-7bf9a512b13b",
			},
			mock: mockArgs{
				role: model.Role{
					ID:          "382d211c-2a22-47c2-aa7d-7bf9a512b13b",
					Permissions: []model.Permission{model.PermModifyUserPermissions},
				},
				err: nil,
			},
			want: model.Role{
				ID:          "382d211c-2a22-47c2-aa7d-7bf9a512b13b",
				Permissions: []model.Permission{model.PermModifyUserPermissions},
			},
			wantErr: false,
		},
		{
			name: "bad read",
			args: args{
				roleID: "e297951d-1c2c-420a-851d-d535340b10cb",
			},
			mock: mockArgs{
				err: fmt.Errorf("some error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(mocks.Repository)
			r.On("GetRole", mock.Anything, tt.args.roleID).
				Once().Return(tt.mock.role, tt.mock.err)

			s := &Service{
				repo: r,
			}
			got, err := s.GetRole(context.Background(), tt.args.roleID)
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

func TestService_AssignRole(t *testing.T) {
	type args struct {
		userID         string
		roleID         string
		organizationID string
	}
	type mockArgs struct {
		binding model.RoleBinding
		err     error
	}
	tests := []struct {
		name    string
		args    args
		mock    mockArgs
		wantErr bool
	}{
		{
			name: "normal binding",
			args: args{
				userID:         "2cb6fe9b-09f9-42ea-bb78-bb0c72d30e26",
				roleID:         "649b574b-b808-423e-9d27-004492be3b25",
				organizationID: "8b296ffb-0e8e-48c2-819b-adc80b5e407a",
			},
			mock: mockArgs{
				binding: model.RoleBinding{
					UserID: "2cb6fe9b-09f9-42ea-bb78-bb0c72d30e26",
					RoleID: "649b574b-b808-423e-9d27-004492be3b25",
					Scope: model.Scope{
						OrganizationIDs: []string{"8b296ffb-0e8e-48c2-819b-adc80b5e407a"},
					},
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "bad binding",
			args: args{
				userID:         "2cb6fe9b-09f9-42ea-bb78-bb0c72d30e26",
				roleID:         "649b574b-b808-423e-9d27-004492be3b25",
				organizationID: "8b296ffb-0e8e-48c2-819b-adc80b5e407a",
			},
			mock: mockArgs{
				binding: model.RoleBinding{
					UserID: "2cb6fe9b-09f9-42ea-bb78-bb0c72d30e26",
					RoleID: "649b574b-b808-423e-9d27-004492be3b25",
					Scope: model.Scope{
						OrganizationIDs: []string{"8b296ffb-0e8e-48c2-819b-adc80b5e407a"},
					},
				},
				err: fmt.Errorf("test error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(mocks.Repository)
			r.On("UpsertRoleBinding", mock.Anything, tt.mock.binding).
				Once().Return(tt.mock.err)

			s := &Service{
				repo: r,
			}
			if err := s.AssignRole(context.Background(), tt.args.userID, tt.args.roleID, tt.args.organizationID); (err != nil) != tt.wantErr {
				t.Errorf("AssignRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_RevokeRole(t *testing.T) {
	type args struct {
		userID string
		roleID string
	}
	type mockArgs struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		mock    mockArgs
		wantErr bool
	}{
		{
			name: "good revoke",
			args: args{
				userID: "812edad7-b14e-40b0-8c66-9de17148a7d0",
				roleID: "f7b2bceb-2f82-4758-8703-f4a281419853",
			},
			mock: mockArgs{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "bad revoke",
			args: args{
				userID: "812edad7-b14e-40b0-8c66-9de17148a7d0",
				roleID: "f7b2bceb-2f82-4758-8703-f4a281419853",
			},
			mock: mockArgs{
				err: fmt.Errorf("test error"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(mocks.Repository)
			r.On("DeleteRoleBinding", mock.Anything, tt.args.userID, tt.args.roleID).
				Once().Return(tt.mock.err)

			s := &Service{
				repo: r,
			}
			if err := s.RevokeRole(context.Background(), tt.args.userID, tt.args.roleID); (err != nil) != tt.wantErr {
				t.Errorf("RevokeRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
