// Code generated by entc, DO NOT EDIT.

package pinfo

const (
	// Label holds the string label denoting the pinfo type in the database.
	Label = "pinfo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProfile holds the string denoting the profile field in the database.
	FieldProfile = "profile"
	// FieldIdCardNumber holds the string denoting the idcardnumber field in the database.
	FieldIdCardNumber = "id_card_number"
	// FieldPrefix holds the string denoting the prefix field in the database.
	FieldPrefix = "prefix"
	// FieldFirstName holds the string denoting the firstname field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the lastname field in the database.
	FieldLastName = "last_name"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldBrithDate holds the string denoting the brithdate field in the database.
	FieldBrithDate = "brith_date"
	// FieldBloodGroup holds the string denoting the bloodgroup field in the database.
	FieldBloodGroup = "blood_group"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldAbout holds the string denoting the about field in the database.
	FieldAbout = "about"
	// EdgeWhoIsTheOwnerOfThisPInfo holds the string denoting the who_is_the_owner_of_this_pinfo edge name in mutations.
	EdgeWhoIsTheOwnerOfThisPInfo = "who_is_the_owner_of_this_PInfo"
	// Table holds the table name of the pinfo in the database.
	Table = "pinfos"
	// WhoIsTheOwnerOfThisPInfoTable is the table that holds the who_is_the_owner_of_this_PInfo relation/edge.
	WhoIsTheOwnerOfThisPInfoTable = "pinfos"
	// WhoIsTheOwnerOfThisPInfoInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	WhoIsTheOwnerOfThisPInfoInverseTable = "users"
	// WhoIsTheOwnerOfThisPInfoColumn is the table column denoting the who_is_the_owner_of_this_PInfo relation/edge.
	WhoIsTheOwnerOfThisPInfoColumn = "user_user_has_p_info"
)

// Columns holds all SQL columns for pinfo fields.
var Columns = []string{
	FieldID,
	FieldProfile,
	FieldIdCardNumber,
	FieldPrefix,
	FieldFirstName,
	FieldLastName,
	FieldGender,
	FieldBrithDate,
	FieldBloodGroup,
	FieldAddress,
	FieldAbout,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "pinfos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_user_has_p_info",
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
