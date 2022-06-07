// Code generated by entc, DO NOT EDIT.

package admin

const (
	// Label holds the string label denoting the admin type in the database.
	Label = "admin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeHabit holds the string denoting the habit edge name in mutations.
	EdgeHabit = "habit"
	// EdgeInvitations holds the string denoting the invitations edge name in mutations.
	EdgeInvitations = "invitations"
	// Table holds the table name of the admin in the database.
	Table = "admins"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "admins"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_admins"
	// HabitTable is the table that holds the habit relation/edge.
	HabitTable = "admins"
	// HabitInverseTable is the table name for the Habit entity.
	// It exists in this package in order to avoid circular dependency with the "habit" package.
	HabitInverseTable = "habits"
	// HabitColumn is the table column denoting the habit relation/edge.
	HabitColumn = "habit_admins"
	// InvitationsTable is the table that holds the invitations relation/edge.
	InvitationsTable = "invitations"
	// InvitationsInverseTable is the table name for the Invitation entity.
	// It exists in this package in order to avoid circular dependency with the "invitation" package.
	InvitationsInverseTable = "invitations"
	// InvitationsColumn is the table column denoting the invitations relation/edge.
	InvitationsColumn = "admin_invitations"
)

// Columns holds all SQL columns for admin fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "admins"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"habit_admins",
	"user_admins",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
