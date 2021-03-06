// Code generated by entc, DO NOT EDIT.

package appointment

const (
	// Label holds the string label denoting the appointment type in the database.
	Label = "appointment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldReasonForAppointment holds the string denoting the reasonforappointment field in the database.
	FieldReasonForAppointment = "reason_for_appointment"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldStartTime holds the string denoting the starttime field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the endtime field in the database.
	FieldEndTime = "end_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDoctorId holds the string denoting the doctorid field in the database.
	FieldDoctorId = "doctor_id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// EdgeAppointmentSchedule holds the string denoting the appointment_schedule edge name in mutations.
	EdgeAppointmentSchedule = "appointment_schedule"
	// EdgeAppointmentChat holds the string denoting the appointment_chat edge name in mutations.
	EdgeAppointmentChat = "appointment_chat"
	// Table holds the table name of the appointment in the database.
	Table = "appointments"
	// AppointmentScheduleTable is the table that holds the appointment_schedule relation/edge.
	AppointmentScheduleTable = "appointments"
	// AppointmentScheduleInverseTable is the table name for the Schedule entity.
	// It exists in this package in order to avoid circular dependency with the "schedule" package.
	AppointmentScheduleInverseTable = "schedules"
	// AppointmentScheduleColumn is the table column denoting the appointment_schedule relation/edge.
	AppointmentScheduleColumn = "schedule_schedule_appointment"
	// AppointmentChatTable is the table that holds the appointment_chat relation/edge.
	AppointmentChatTable = "appointments"
	// AppointmentChatInverseTable is the table name for the Chat entity.
	// It exists in this package in order to avoid circular dependency with the "chat" package.
	AppointmentChatInverseTable = "chats"
	// AppointmentChatColumn is the table column denoting the appointment_chat relation/edge.
	AppointmentChatColumn = "chat_chat_appointment"
)

// Columns holds all SQL columns for appointment fields.
var Columns = []string{
	FieldID,
	FieldReasonForAppointment,
	FieldDetail,
	FieldStartTime,
	FieldEndTime,
	FieldStatus,
	FieldDoctorId,
	FieldUserId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "appointments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"chat_chat_appointment",
	"schedule_schedule_appointment",
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
