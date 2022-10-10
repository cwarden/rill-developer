// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"
)

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedOn holds the string denoting the created_on field in the database.
	FieldCreatedOn = "created_on"
	// FieldUpdatedOn holds the string denoting the updated_on field in the database.
	FieldUpdatedOn = "updated_on"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the organization in the database.
	Table = "organizations"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "organization_groups"
	// GroupsInverseTable is the table name for the Groups entity.
	// It exists in this package in order to avoid circular dependency with the "groups" package.
	GroupsInverseTable = "groups"
	// ProjectsTable is the table that holds the projects relation/edge.
	ProjectsTable = "projects"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
	// ProjectsColumn is the table column denoting the projects relation/edge.
	ProjectsColumn = "organization_projects"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "organization_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldCreatedOn,
	FieldUpdatedOn,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"organization_id", "groups_id"}
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"organization_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDescription holds the default value on creation for the "Description" field.
	DefaultDescription string
	// DefaultCreatedOn holds the default value on creation for the "created_on" field.
	DefaultCreatedOn func() time.Time
	// DefaultUpdatedOn holds the default value on creation for the "updated_on" field.
	DefaultUpdatedOn func() time.Time
)
