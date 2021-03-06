// Code generated by entc, DO NOT EDIT.

package role

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeRoleUser holds the string denoting the role_user edge name in mutations.
	EdgeRoleUser = "role_user"
	// Table holds the table name of the role in the database.
	Table = "roles"
	// RoleUserTable is the table that holds the role_user relation/edge. The primary key declared below.
	RoleUserTable = "role_role_user"
	// RoleUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RoleUserInverseTable = "users"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// RoleUserPrimaryKey and RoleUserColumn2 are the table columns denoting the
	// primary key for the role_user relation (M2M).
	RoleUserPrimaryKey = []string{"role_id", "user_id"}
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
