// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/certification"
	"FinalProject/ent/department"
	"FinalProject/ent/hospital"
	"FinalProject/ent/schedule"
	"FinalProject/ent/schema"
	"FinalProject/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	certificationFields := schema.Certification{}.Fields()
	_ = certificationFields
	// certificationDescCode is the schema descriptor for code field.
	certificationDescCode := certificationFields[0].Descriptor()
	// certification.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	certification.CodeValidator = certificationDescCode.Validators[0].(func(string) error)
	// certificationDescDiloma is the schema descriptor for diloma field.
	certificationDescDiloma := certificationFields[1].Descriptor()
	// certification.DilomaValidator is a validator for the "diloma" field. It is called by the builders before save.
	certification.DilomaValidator = certificationDescDiloma.Validators[0].(func(string) error)
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescName is the schema descriptor for name field.
	departmentDescName := departmentFields[0].Descriptor()
	// department.NameValidator is a validator for the "name" field. It is called by the builders before save.
	department.NameValidator = departmentDescName.Validators[0].(func(string) error)
	hospitalFields := schema.Hospital{}.Fields()
	_ = hospitalFields
	// hospitalDescName is the schema descriptor for name field.
	hospitalDescName := hospitalFields[0].Descriptor()
	// hospital.NameValidator is a validator for the "name" field. It is called by the builders before save.
	hospital.NameValidator = hospitalDescName.Validators[0].(func(string) error)
	scheduleFields := schema.Schedule{}.Fields()
	_ = scheduleFields
	// scheduleDescStatus is the schema descriptor for status field.
	scheduleDescStatus := scheduleFields[2].Descriptor()
	// schedule.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	schedule.StatusValidator = scheduleDescStatus.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
}
