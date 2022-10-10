// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: "Unknown"},
		{Name: "created_on", Type: field.TypeTime},
		{Name: "updated_on", Type: field.TypeTime},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: "Unknown"},
		{Name: "created_on", Type: field.TypeTime},
		{Name: "updated_on", Type: field.TypeTime},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_on", Type: field.TypeTime},
		{Name: "updated_on", Type: field.TypeTime},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Default: "Unknown"},
		{Name: "created_on", Type: field.TypeTime},
		{Name: "updated_on", Type: field.TypeTime},
		{Name: "organization_projects", Type: field.TypeInt, Nullable: true},
		{Name: "user_projects", Type: field.TypeInt, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "projects_organizations_projects",
				Columns:    []*schema.Column{ProjectsColumns[5]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "projects_users_projects",
				Columns:    []*schema.Column{ProjectsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_on", Type: field.TypeTime},
		{Name: "updated_on", Type: field.TypeTime},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Default: "Unknown"},
		{Name: "created_on", Type: field.TypeTime},
		{Name: "updated_on", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// GroupsUsersColumns holds the columns for the "groups_users" table.
	GroupsUsersColumns = []*schema.Column{
		{Name: "groups_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupsUsersTable holds the schema information for the "groups_users" table.
	GroupsUsersTable = &schema.Table{
		Name:       "groups_users",
		Columns:    GroupsUsersColumns,
		PrimaryKey: []*schema.Column{GroupsUsersColumns[0], GroupsUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_users_groups_id",
				Columns:    []*schema.Column{GroupsUsersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "groups_users_user_id",
				Columns:    []*schema.Column{GroupsUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupsProjectsColumns holds the columns for the "groups_projects" table.
	GroupsProjectsColumns = []*schema.Column{
		{Name: "groups_id", Type: field.TypeInt},
		{Name: "project_id", Type: field.TypeInt},
	}
	// GroupsProjectsTable holds the schema information for the "groups_projects" table.
	GroupsProjectsTable = &schema.Table{
		Name:       "groups_projects",
		Columns:    GroupsProjectsColumns,
		PrimaryKey: []*schema.Column{GroupsProjectsColumns[0], GroupsProjectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_projects_groups_id",
				Columns:    []*schema.Column{GroupsProjectsColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "groups_projects_project_id",
				Columns:    []*schema.Column{GroupsProjectsColumns[1]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OrganizationGroupsColumns holds the columns for the "organization_groups" table.
	OrganizationGroupsColumns = []*schema.Column{
		{Name: "organization_id", Type: field.TypeInt},
		{Name: "groups_id", Type: field.TypeInt},
	}
	// OrganizationGroupsTable holds the schema information for the "organization_groups" table.
	OrganizationGroupsTable = &schema.Table{
		Name:       "organization_groups",
		Columns:    OrganizationGroupsColumns,
		PrimaryKey: []*schema.Column{OrganizationGroupsColumns[0], OrganizationGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_groups_organization_id",
				Columns:    []*schema.Column{OrganizationGroupsColumns[0]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "organization_groups_groups_id",
				Columns:    []*schema.Column{OrganizationGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OrganizationUsersColumns holds the columns for the "organization_users" table.
	OrganizationUsersColumns = []*schema.Column{
		{Name: "organization_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// OrganizationUsersTable holds the schema information for the "organization_users" table.
	OrganizationUsersTable = &schema.Table{
		Name:       "organization_users",
		Columns:    OrganizationUsersColumns,
		PrimaryKey: []*schema.Column{OrganizationUsersColumns[0], OrganizationUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_users_organization_id",
				Columns:    []*schema.Column{OrganizationUsersColumns[0]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "organization_users_user_id",
				Columns:    []*schema.Column{OrganizationUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RolePermissionColumns holds the columns for the "role_permission" table.
	RolePermissionColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt},
		{Name: "permission_id", Type: field.TypeInt},
	}
	// RolePermissionTable holds the schema information for the "role_permission" table.
	RolePermissionTable = &schema.Table{
		Name:       "role_permission",
		Columns:    RolePermissionColumns,
		PrimaryKey: []*schema.Column{RolePermissionColumns[0], RolePermissionColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_permission_role_id",
				Columns:    []*schema.Column{RolePermissionColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_permission_permission_id",
				Columns:    []*schema.Column{RolePermissionColumns[1]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserRoleColumns holds the columns for the "user_role" table.
	UserRoleColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// UserRoleTable holds the schema information for the "user_role" table.
	UserRoleTable = &schema.Table{
		Name:       "user_role",
		Columns:    UserRoleColumns,
		PrimaryKey: []*schema.Column{UserRoleColumns[0], UserRoleColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_role_user_id",
				Columns:    []*schema.Column{UserRoleColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_role_role_id",
				Columns:    []*schema.Column{UserRoleColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupsTable,
		OrganizationsTable,
		PermissionsTable,
		ProjectsTable,
		RolesTable,
		UsersTable,
		GroupsUsersTable,
		GroupsProjectsTable,
		OrganizationGroupsTable,
		OrganizationUsersTable,
		RolePermissionTable,
		UserRoleTable,
	}
)

func init() {
	ProjectsTable.ForeignKeys[0].RefTable = OrganizationsTable
	ProjectsTable.ForeignKeys[1].RefTable = UsersTable
	GroupsUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupsUsersTable.ForeignKeys[1].RefTable = UsersTable
	GroupsProjectsTable.ForeignKeys[0].RefTable = GroupsTable
	GroupsProjectsTable.ForeignKeys[1].RefTable = ProjectsTable
	OrganizationGroupsTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationGroupsTable.ForeignKeys[1].RefTable = GroupsTable
	OrganizationUsersTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationUsersTable.ForeignKeys[1].RefTable = UsersTable
	RolePermissionTable.ForeignKeys[0].RefTable = RolesTable
	RolePermissionTable.ForeignKeys[1].RefTable = PermissionsTable
	UserRoleTable.ForeignKeys[0].RefTable = UsersTable
	UserRoleTable.ForeignKeys[1].RefTable = RolesTable
}
