// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/rilldata/rill/server-cloud/ent/groups"
	"github.com/rilldata/rill/server-cloud/ent/organization"
	"github.com/rilldata/rill/server-cloud/ent/permission"
	"github.com/rilldata/rill/server-cloud/ent/project"
	"github.com/rilldata/rill/server-cloud/ent/role"
	"github.com/rilldata/rill/server-cloud/ent/schema"
	"github.com/rilldata/rill/server-cloud/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	groupsFields := schema.Groups{}.Fields()
	_ = groupsFields
	// groupsDescName is the schema descriptor for Name field.
	groupsDescName := groupsFields[0].Descriptor()
	// groups.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	groups.NameValidator = groupsDescName.Validators[0].(func(string) error)
	// groupsDescDescription is the schema descriptor for Description field.
	groupsDescDescription := groupsFields[1].Descriptor()
	// groups.DefaultDescription holds the default value on creation for the Description field.
	groups.DefaultDescription = groupsDescDescription.Default.(string)
	// groupsDescCreatedOn is the schema descriptor for created_on field.
	groupsDescCreatedOn := groupsFields[2].Descriptor()
	// groups.DefaultCreatedOn holds the default value on creation for the created_on field.
	groups.DefaultCreatedOn = groupsDescCreatedOn.Default.(func() time.Time)
	// groupsDescUpdatedOn is the schema descriptor for updated_on field.
	groupsDescUpdatedOn := groupsFields[3].Descriptor()
	// groups.DefaultUpdatedOn holds the default value on creation for the updated_on field.
	groups.DefaultUpdatedOn = groupsDescUpdatedOn.Default.(func() time.Time)
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescDescription is the schema descriptor for Description field.
	organizationDescDescription := organizationFields[1].Descriptor()
	// organization.DefaultDescription holds the default value on creation for the Description field.
	organization.DefaultDescription = organizationDescDescription.Default.(string)
	// organizationDescCreatedOn is the schema descriptor for created_on field.
	organizationDescCreatedOn := organizationFields[2].Descriptor()
	// organization.DefaultCreatedOn holds the default value on creation for the created_on field.
	organization.DefaultCreatedOn = organizationDescCreatedOn.Default.(func() time.Time)
	// organizationDescUpdatedOn is the schema descriptor for updated_on field.
	organizationDescUpdatedOn := organizationFields[3].Descriptor()
	// organization.DefaultUpdatedOn holds the default value on creation for the updated_on field.
	organization.DefaultUpdatedOn = organizationDescUpdatedOn.Default.(func() time.Time)
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescName is the schema descriptor for Name field.
	permissionDescName := permissionFields[0].Descriptor()
	// permission.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	permission.NameValidator = permissionDescName.Validators[0].(func(string) error)
	// permissionDescCreatedOn is the schema descriptor for created_on field.
	permissionDescCreatedOn := permissionFields[1].Descriptor()
	// permission.DefaultCreatedOn holds the default value on creation for the created_on field.
	permission.DefaultCreatedOn = permissionDescCreatedOn.Default.(func() time.Time)
	// permissionDescUpdatedOn is the schema descriptor for updated_on field.
	permissionDescUpdatedOn := permissionFields[2].Descriptor()
	// permission.DefaultUpdatedOn holds the default value on creation for the updated_on field.
	permission.DefaultUpdatedOn = permissionDescUpdatedOn.Default.(time.Time)
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescName is the schema descriptor for Name field.
	projectDescName := projectFields[0].Descriptor()
	// project.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	project.NameValidator = projectDescName.Validators[0].(func(string) error)
	// projectDescDescription is the schema descriptor for Description field.
	projectDescDescription := projectFields[1].Descriptor()
	// project.DefaultDescription holds the default value on creation for the Description field.
	project.DefaultDescription = projectDescDescription.Default.(string)
	// projectDescCreatedOn is the schema descriptor for created_on field.
	projectDescCreatedOn := projectFields[2].Descriptor()
	// project.DefaultCreatedOn holds the default value on creation for the created_on field.
	project.DefaultCreatedOn = projectDescCreatedOn.Default.(func() time.Time)
	// projectDescUpdatedOn is the schema descriptor for updated_on field.
	projectDescUpdatedOn := projectFields[3].Descriptor()
	// project.DefaultUpdatedOn holds the default value on creation for the updated_on field.
	project.DefaultUpdatedOn = projectDescUpdatedOn.Default.(func() time.Time)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescName is the schema descriptor for Name field.
	roleDescName := roleFields[0].Descriptor()
	// role.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescCreatedOn is the schema descriptor for created_on field.
	roleDescCreatedOn := roleFields[1].Descriptor()
	// role.DefaultCreatedOn holds the default value on creation for the created_on field.
	role.DefaultCreatedOn = roleDescCreatedOn.Default.(func() time.Time)
	// roleDescUpdatedOn is the schema descriptor for updated_on field.
	roleDescUpdatedOn := roleFields[2].Descriptor()
	// role.DefaultUpdatedOn holds the default value on creation for the updated_on field.
	role.DefaultUpdatedOn = roleDescUpdatedOn.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserName is the schema descriptor for UserName field.
	userDescUserName := userFields[1].Descriptor()
	// user.UserNameValidator is a validator for the "UserName" field. It is called by the builders before save.
	user.UserNameValidator = userDescUserName.Validators[0].(func(string) error)
	// userDescDescription is the schema descriptor for Description field.
	userDescDescription := userFields[2].Descriptor()
	// user.DefaultDescription holds the default value on creation for the Description field.
	user.DefaultDescription = userDescDescription.Default.(string)
	// userDescCreatedOn is the schema descriptor for created_on field.
	userDescCreatedOn := userFields[3].Descriptor()
	// user.DefaultCreatedOn holds the default value on creation for the created_on field.
	user.DefaultCreatedOn = userDescCreatedOn.Default.(func() time.Time)
	// userDescUpdatedOn is the schema descriptor for updated_on field.
	userDescUpdatedOn := userFields[4].Descriptor()
	// user.DefaultUpdatedOn holds the default value on creation for the updated_on field.
	user.DefaultUpdatedOn = userDescUpdatedOn.Default.(time.Time)
}
