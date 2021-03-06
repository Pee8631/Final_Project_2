// Code generated by entc, DO NOT EDIT.

package treatment

const (
	// Label holds the string label denoting the treatment type in the database.
	Label = "treatment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTreatmentRecord holds the string denoting the treatmentrecord field in the database.
	FieldTreatmentRecord = "treatment_record"
	// FieldDateTime holds the string denoting the datetime field in the database.
	FieldDateTime = "date_time"
	// FieldTakeTime holds the string denoting the taketime field in the database.
	FieldTakeTime = "take_time"
	// EdgeTreatmentWasRecordedByDoctor holds the string denoting the treatment_was_recorded_by_doctor edge name in mutations.
	EdgeTreatmentWasRecordedByDoctor = "treatment_was_recorded_by_doctor"
	// EdgeUserIsTheTreatmentOfRecord holds the string denoting the user_is_the_treatment_of_record edge name in mutations.
	EdgeUserIsTheTreatmentOfRecord = "user_is_the_treatment_of_record"
	// Table holds the table name of the treatment in the database.
	Table = "treatments"
	// TreatmentWasRecordedByDoctorTable is the table that holds the treatment_was_recorded_by_doctor relation/edge.
	TreatmentWasRecordedByDoctorTable = "treatments"
	// TreatmentWasRecordedByDoctorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	TreatmentWasRecordedByDoctorInverseTable = "users"
	// TreatmentWasRecordedByDoctorColumn is the table column denoting the treatment_was_recorded_by_doctor relation/edge.
	TreatmentWasRecordedByDoctorColumn = "user_doctor_record_treatment"
	// UserIsTheTreatmentOfRecordTable is the table that holds the user_is_the_treatment_of_record relation/edge.
	UserIsTheTreatmentOfRecordTable = "treatments"
	// UserIsTheTreatmentOfRecordInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserIsTheTreatmentOfRecordInverseTable = "users"
	// UserIsTheTreatmentOfRecordColumn is the table column denoting the user_is_the_treatment_of_record relation/edge.
	UserIsTheTreatmentOfRecordColumn = "user_user_have_treatment"
)

// Columns holds all SQL columns for treatment fields.
var Columns = []string{
	FieldID,
	FieldTreatmentRecord,
	FieldDateTime,
	FieldTakeTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "treatments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_doctor_record_treatment",
	"user_user_have_treatment",
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
