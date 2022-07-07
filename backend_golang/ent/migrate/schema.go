// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppointmentsColumns holds the columns for the "appointments" table.
	AppointmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "reason_for_appointment", Type: field.TypeString},
		{Name: "detail", Type: field.TypeString, Size: 2147483647},
		{Name: "start_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "end_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "status", Type: field.TypeString},
		{Name: "doctor_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "chat_chat_appointment", Type: field.TypeInt, Nullable: true},
		{Name: "schedule_schedule_appointment", Type: field.TypeInt, Nullable: true},
	}
	// AppointmentsTable holds the schema information for the "appointments" table.
	AppointmentsTable = &schema.Table{
		Name:       "appointments",
		Columns:    AppointmentsColumns,
		PrimaryKey: []*schema.Column{AppointmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "appointments_chats_chat_appointment",
				Columns:    []*schema.Column{AppointmentsColumns[8]},
				RefColumns: []*schema.Column{ChatsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "appointments_schedules_schedule_appointment",
				Columns:    []*schema.Column{AppointmentsColumns[9]},
				RefColumns: []*schema.Column{SchedulesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CertificationsColumns holds the columns for the "certifications" table.
	CertificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString},
		{Name: "diloma", Type: field.TypeString},
		{Name: "date_of_issuing", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "date_of_exp", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "issuer", Type: field.TypeString},
		{Name: "user_doctor_has_certification", Type: field.TypeInt, Nullable: true},
	}
	// CertificationsTable holds the schema information for the "certifications" table.
	CertificationsTable = &schema.Table{
		Name:       "certifications",
		Columns:    CertificationsColumns,
		PrimaryKey: []*schema.Column{CertificationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "certifications_users_doctor_has_certification",
				Columns:    []*schema.Column{CertificationsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChatsColumns holds the columns for the "chats" table.
	ChatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "chat_room_name", Type: field.TypeString},
		{Name: "is_lock_chat", Type: field.TypeBool, Default: false},
	}
	// ChatsTable holds the schema information for the "chats" table.
	ChatsTable = &schema.Table{
		Name:       "chats",
		Columns:    ChatsColumns,
		PrimaryKey: []*schema.Column{ChatsColumns[0]},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:       "departments",
		Columns:    DepartmentsColumns,
		PrimaryKey: []*schema.Column{DepartmentsColumns[0]},
	}
	// DiseasesColumns holds the columns for the "diseases" table.
	DiseasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "symtoms", Type: field.TypeString, Size: 2147483647},
	}
	// DiseasesTable holds the schema information for the "diseases" table.
	DiseasesTable = &schema.Table{
		Name:       "diseases",
		Columns:    DiseasesColumns,
		PrimaryKey: []*schema.Column{DiseasesColumns[0]},
	}
	// HospitalsColumns holds the columns for the "hospitals" table.
	HospitalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// HospitalsTable holds the schema information for the "hospitals" table.
	HospitalsTable = &schema.Table{
		Name:       "hospitals",
		Columns:    HospitalsColumns,
		PrimaryKey: []*schema.Column{HospitalsColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "message_text", Type: field.TypeString},
		{Name: "sent_date_time", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "chat_chat_message", Type: field.TypeInt, Nullable: true},
		{Name: "user_user_send_message", Type: field.TypeInt, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_chats_chat_message",
				Columns:    []*schema.Column{MessagesColumns[3]},
				RefColumns: []*schema.Column{ChatsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "messages_users_user_send_message",
				Columns:    []*schema.Column{MessagesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "message", Type: field.TypeString},
		{Name: "created_date", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "recipient_id", Type: field.TypeInt},
		{Name: "sender_id", Type: field.TypeInt},
		{Name: "appointment_id", Type: field.TypeInt},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
	}
	// PinfosColumns holds the columns for the "pinfos" table.
	PinfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "profile", Type: field.TypeString},
		{Name: "id_card_number", Type: field.TypeString, Unique: true},
		{Name: "prefix", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "gender", Type: field.TypeInt},
		{Name: "brith_date", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "blood_group", Type: field.TypeString},
		{Name: "address", Type: field.TypeString, Size: 2147483647},
		{Name: "about", Type: field.TypeString, Size: 2147483647},
		{Name: "user_user_has_p_info", Type: field.TypeInt, Nullable: true},
	}
	// PinfosTable holds the schema information for the "pinfos" table.
	PinfosTable = &schema.Table{
		Name:       "pinfos",
		Columns:    PinfosColumns,
		PrimaryKey: []*schema.Column{PinfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pinfos_users_user_has_PInfo",
				Columns:    []*schema.Column{PinfosColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// SchedulesColumns holds the columns for the "schedules" table.
	SchedulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "activity", Type: field.TypeString},
		{Name: "detail", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "end_time", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "user_doctor_schedule", Type: field.TypeInt, Nullable: true},
	}
	// SchedulesTable holds the schema information for the "schedules" table.
	SchedulesTable = &schema.Table{
		Name:       "schedules",
		Columns:    SchedulesColumns,
		PrimaryKey: []*schema.Column{SchedulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "schedules_users_doctor_schedule",
				Columns:    []*schema.Column{SchedulesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TelecomsColumns holds the columns for the "telecoms" table.
	TelecomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "telephone", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString},
		{Name: "platform", Type: field.TypeString},
		{Name: "user_user_have_telecoms", Type: field.TypeInt, Nullable: true},
	}
	// TelecomsTable holds the schema information for the "telecoms" table.
	TelecomsTable = &schema.Table{
		Name:       "telecoms",
		Columns:    TelecomsColumns,
		PrimaryKey: []*schema.Column{TelecomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "telecoms_users_user_have_telecoms",
				Columns:    []*schema.Column{TelecomsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "auth_token", Type: field.TypeString},
		{Name: "generated_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "user_user_have_token", Type: field.TypeInt, Nullable: true},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tokens_users_user_have_token",
				Columns:    []*schema.Column{TokensColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TreatmentsColumns holds the columns for the "treatments" table.
	TreatmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "treatment_record", Type: field.TypeString, Size: 2147483647},
		{Name: "date_time", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "take_time", Type: field.TypeFloat64},
		{Name: "user_doctor_record_treatment", Type: field.TypeInt, Nullable: true},
		{Name: "user_user_have_treatment", Type: field.TypeInt, Nullable: true},
	}
	// TreatmentsTable holds the schema information for the "treatments" table.
	TreatmentsTable = &schema.Table{
		Name:       "treatments",
		Columns:    TreatmentsColumns,
		PrimaryKey: []*schema.Column{TreatmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "treatments_users_doctor_record_treatment",
				Columns:    []*schema.Column{TreatmentsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "treatments_users_user_have_treatment",
				Columns:    []*schema.Column{TreatmentsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "department_department_has_doctor", Type: field.TypeInt, Nullable: true},
		{Name: "hospital_hospital_has_doctor", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_departments_department_has_doctor",
				Columns:    []*schema.Column{UsersColumns[3]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_hospitals_hospital_has_doctor",
				Columns:    []*schema.Column{UsersColumns[4]},
				RefColumns: []*schema.Column{HospitalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChatChatUserColumns holds the columns for the "chat_chat_user" table.
	ChatChatUserColumns = []*schema.Column{
		{Name: "chat_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ChatChatUserTable holds the schema information for the "chat_chat_user" table.
	ChatChatUserTable = &schema.Table{
		Name:       "chat_chat_user",
		Columns:    ChatChatUserColumns,
		PrimaryKey: []*schema.Column{ChatChatUserColumns[0], ChatChatUserColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "chat_chat_user_chat_id",
				Columns:    []*schema.Column{ChatChatUserColumns[0]},
				RefColumns: []*schema.Column{ChatsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "chat_chat_user_user_id",
				Columns:    []*schema.Column{ChatChatUserColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// DiseaseDiseaseUserColumns holds the columns for the "disease_disease_user" table.
	DiseaseDiseaseUserColumns = []*schema.Column{
		{Name: "disease_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// DiseaseDiseaseUserTable holds the schema information for the "disease_disease_user" table.
	DiseaseDiseaseUserTable = &schema.Table{
		Name:       "disease_disease_user",
		Columns:    DiseaseDiseaseUserColumns,
		PrimaryKey: []*schema.Column{DiseaseDiseaseUserColumns[0], DiseaseDiseaseUserColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "disease_disease_user_disease_id",
				Columns:    []*schema.Column{DiseaseDiseaseUserColumns[0]},
				RefColumns: []*schema.Column{DiseasesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "disease_disease_user_user_id",
				Columns:    []*schema.Column{DiseaseDiseaseUserColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoleRoleUserColumns holds the columns for the "role_role_user" table.
	RoleRoleUserColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// RoleRoleUserTable holds the schema information for the "role_role_user" table.
	RoleRoleUserTable = &schema.Table{
		Name:       "role_role_user",
		Columns:    RoleRoleUserColumns,
		PrimaryKey: []*schema.Column{RoleRoleUserColumns[0], RoleRoleUserColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_role_user_role_id",
				Columns:    []*schema.Column{RoleRoleUserColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_role_user_user_id",
				Columns:    []*schema.Column{RoleRoleUserColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserUserSendNotificationColumns holds the columns for the "user_user_send_notification" table.
	UserUserSendNotificationColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "notification_id", Type: field.TypeInt},
	}
	// UserUserSendNotificationTable holds the schema information for the "user_user_send_notification" table.
	UserUserSendNotificationTable = &schema.Table{
		Name:       "user_user_send_notification",
		Columns:    UserUserSendNotificationColumns,
		PrimaryKey: []*schema.Column{UserUserSendNotificationColumns[0], UserUserSendNotificationColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_user_send_notification_user_id",
				Columns:    []*schema.Column{UserUserSendNotificationColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_user_send_notification_notification_id",
				Columns:    []*schema.Column{UserUserSendNotificationColumns[1]},
				RefColumns: []*schema.Column{NotificationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppointmentsTable,
		CertificationsTable,
		ChatsTable,
		DepartmentsTable,
		DiseasesTable,
		HospitalsTable,
		MessagesTable,
		NotificationsTable,
		PinfosTable,
		RolesTable,
		SchedulesTable,
		TelecomsTable,
		TokensTable,
		TreatmentsTable,
		UsersTable,
		ChatChatUserTable,
		DiseaseDiseaseUserTable,
		RoleRoleUserTable,
		UserUserSendNotificationTable,
	}
)

func init() {
	AppointmentsTable.ForeignKeys[0].RefTable = ChatsTable
	AppointmentsTable.ForeignKeys[1].RefTable = SchedulesTable
	CertificationsTable.ForeignKeys[0].RefTable = UsersTable
	MessagesTable.ForeignKeys[0].RefTable = ChatsTable
	MessagesTable.ForeignKeys[1].RefTable = UsersTable
	PinfosTable.ForeignKeys[0].RefTable = UsersTable
	SchedulesTable.ForeignKeys[0].RefTable = UsersTable
	TelecomsTable.ForeignKeys[0].RefTable = UsersTable
	TokensTable.ForeignKeys[0].RefTable = UsersTable
	TreatmentsTable.ForeignKeys[0].RefTable = UsersTable
	TreatmentsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = DepartmentsTable
	UsersTable.ForeignKeys[1].RefTable = HospitalsTable
	ChatChatUserTable.ForeignKeys[0].RefTable = ChatsTable
	ChatChatUserTable.ForeignKeys[1].RefTable = UsersTable
	DiseaseDiseaseUserTable.ForeignKeys[0].RefTable = DiseasesTable
	DiseaseDiseaseUserTable.ForeignKeys[1].RefTable = UsersTable
	RoleRoleUserTable.ForeignKeys[0].RefTable = RolesTable
	RoleRoleUserTable.ForeignKeys[1].RefTable = UsersTable
	UserUserSendNotificationTable.ForeignKeys[0].RefTable = UsersTable
	UserUserSendNotificationTable.ForeignKeys[1].RefTable = NotificationsTable
}
