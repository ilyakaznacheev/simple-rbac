package model

// Permission is an action that can be performed by a user
type Permission string

const (
	PermManageUsers           Permission = "org.user.manage"            // Add or remove other users to an organization
	PermModifyUserPermissions Permission = "org.user.permission.manage" // Modify a user’s permission in an organization
	PermCreateProject         Permission = "org.project.create"         // Create a new project in an organization
	PermDeleteProject         Permission = "org.project.delete"         // Remove a project from an organization
	PermDeployProject         Permission = "org.project.deploy"         // Deploy a project to a specific environment
	PermManageEnvironments    Permission = "org.project.env.manage"     // Modify environments definitions for a project
	PermReadLogs              Permission = "org.project.log.read"       // Read logs for a project
	PermModifyLogs            Permission = "org.project.log.manage"     // Edit logs for a project
	PermAuditLogs             Permission = "org.project.log.audit"      // Audit logs and sensitive data for a project
)

// Role is a collection of permissions
type Role struct {
	// ID is a unique identifier for the role
	ID string
	// Permissions is a list of permissions granted to the role
	Permissions []Permission
}

// RoleBinding is a mapping between a user and a role
type RoleBinding struct {
	// UserID is the ID of the user
	UserID string
	// RoleID is the ID of the role
	RoleID string
	// Scope is the scope of the role binding
	Scope Scope
}

// Scope is the scope in which the role binding applies
type Scope struct {
	// OrganizationIDs is a list of organizations that the role binding applies to
	OrganizationIDs []string
}
