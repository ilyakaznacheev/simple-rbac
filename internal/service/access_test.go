package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/ilyakaznacheev/simple-rbac/internal/mocks"
	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

func TestService_CheckAuthority(t *testing.T) {

	type args struct {
		userID         string
		organizationID string
		permission     model.Permission
	}
	type mockArgs struct {
		permissions []model.Permission
		err         error
	}
	tests := []struct {
		name    string
		args    args
		mock    mockArgs
		wantErr bool
	}{
		{
			name: "allowed",
			args: args{
				userID:         "68b2a739-1a38-4b98-a202-93851e62b312",
				organizationID: "a37c6138-348e-4012-b8f1-0bd70ecbd229",
				permission:     model.PermManageUsers,
			},
			mock: mockArgs{
				permissions: []model.Permission{model.PermManageUsers, model.PermManageEnvironments},
				err:         nil,
			},
			wantErr: false,
		},
		{
			name: "not allowed",
			args: args{
				userID:         "282e9afb-af0b-4276-9dc7-6dd5a7872899",
				organizationID: "c17cd335-655a-4e9f-b273-c0b573e6af66",
				permission:     model.PermManageUsers,
			},
			mock: mockArgs{
				permissions: []model.Permission{model.PermModifyLogs, model.PermManageEnvironments},
				err:         nil,
			},
			wantErr: true,
		},
		{
			name: "error",
			args: args{
				userID:         "fa333ccb-b29d-4fd1-8a45-26d9d26c474b",
				organizationID: "64bc95ed-6278-45df-8209-00b7ec821f31",
				permission:     model.PermManageUsers,
			},
			mock: mockArgs{
				permissions: nil,
				err:         fmt.Errorf("some error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(mocks.Repository)
			r.On("GetUserPermissionsByOrg", mock.Anything, tt.args.userID, tt.args.organizationID).
				Once().Return(tt.mock.permissions, tt.mock.err)

			s := &Service{
				repo: r,
			}
			err := s.CheckAuthority(context.Background(), tt.args.userID, tt.args.organizationID, tt.args.permission)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckAuthority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
