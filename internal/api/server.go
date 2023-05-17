package api

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/ilyakaznacheev/simple-rbac/api/auth-go"
	"github.com/ilyakaznacheev/simple-rbac/internal/model"
)

type Application interface {
	CheckAuthority(ctx context.Context, userID, organizationID string, permission model.Permission) error
	CreateRole(ctx context.Context, role model.Role) error
	UpdateRole(ctx context.Context, role model.Role) error
	GetRole(ctx context.Context, id string) (*model.Role, error)
	DeleteRole(ctx context.Context, id string) error
	AssignRole(ctx context.Context, userID, roleID, organizationID string) error
	RevokeRole(ctx context.Context, userID, roleID string) error
}

type Server struct {
	server *grpc.Server
	app    Application
}

func New(app Application) *Server {
	s := grpc.NewServer()

	srv := &Server{
		server: s,
		app:    app,
	}

	auth.RegisterAuthorityServiceServer(s, srv)
	auth.RegisterRoleServiceServer(s, srv)

	return srv
}

func (s *Server) Serve(address string) error {
	lst, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("listen port for GRPC server: %w", err)
	}

	return s.server.Serve(lst)
}

func (s *Server) Close() error {
	s.server.Stop()
	return nil
}
