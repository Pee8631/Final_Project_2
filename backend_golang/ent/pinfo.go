// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/pinfo"
	"FinalProject/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// PInfo is the model entity for the PInfo schema.
type PInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Profile holds the value of the "profile" field.
	Profile string `json:"profile,omitempty"`
	// IdCardNumber holds the value of the "idCardNumber" field.
	IdCardNumber string `json:"idCardNumber,omitempty"`
	// Prefix holds the value of the "prefix" field.
	Prefix string `json:"prefix,omitempty"`
	// FirstName holds the value of the "firstName" field.
	FirstName string `json:"firstName,omitempty"`
	// LastName holds the value of the "lastName" field.
	LastName string `json:"lastName,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender int `json:"gender,omitempty"`
	// BrithDate holds the value of the "brithDate" field.
	BrithDate time.Time `json:"brithDate,omitempty"`
	// BloodGroup holds the value of the "bloodGroup" field.
	BloodGroup string `json:"bloodGroup,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// About holds the value of the "about" field.
	About string `json:"about,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PInfoQuery when eager-loading is set.
	Edges                PInfoEdges `json:"edges"`
	user_user_has_p_info *int
}

// PInfoEdges holds the relations/edges for other nodes in the graph.
type PInfoEdges struct {
	// WhoIsTheOwnerOfThisPInfo holds the value of the who_is_the_owner_of_this_PInfo edge.
	WhoIsTheOwnerOfThisPInfo *User `json:"who_is_the_owner_of_this_PInfo,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WhoIsTheOwnerOfThisPInfoOrErr returns the WhoIsTheOwnerOfThisPInfo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PInfoEdges) WhoIsTheOwnerOfThisPInfoOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.WhoIsTheOwnerOfThisPInfo == nil {
			// The edge who_is_the_owner_of_this_PInfo was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.WhoIsTheOwnerOfThisPInfo, nil
	}
	return nil, &NotLoadedError{edge: "who_is_the_owner_of_this_PInfo"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pinfo.FieldID, pinfo.FieldGender:
			values[i] = new(sql.NullInt64)
		case pinfo.FieldProfile, pinfo.FieldIdCardNumber, pinfo.FieldPrefix, pinfo.FieldFirstName, pinfo.FieldLastName, pinfo.FieldBloodGroup, pinfo.FieldAddress, pinfo.FieldAbout:
			values[i] = new(sql.NullString)
		case pinfo.FieldBrithDate:
			values[i] = new(sql.NullTime)
		case pinfo.ForeignKeys[0]: // user_user_has_p_info
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PInfo fields.
func (pi *PInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pi.ID = int(value.Int64)
		case pinfo.FieldProfile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile", values[i])
			} else if value.Valid {
				pi.Profile = value.String
			}
		case pinfo.FieldIdCardNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field idCardNumber", values[i])
			} else if value.Valid {
				pi.IdCardNumber = value.String
			}
		case pinfo.FieldPrefix:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prefix", values[i])
			} else if value.Valid {
				pi.Prefix = value.String
			}
		case pinfo.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firstName", values[i])
			} else if value.Valid {
				pi.FirstName = value.String
			}
		case pinfo.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastName", values[i])
			} else if value.Valid {
				pi.LastName = value.String
			}
		case pinfo.FieldGender:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				pi.Gender = int(value.Int64)
			}
		case pinfo.FieldBrithDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field brithDate", values[i])
			} else if value.Valid {
				pi.BrithDate = value.Time
			}
		case pinfo.FieldBloodGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bloodGroup", values[i])
			} else if value.Valid {
				pi.BloodGroup = value.String
			}
		case pinfo.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				pi.Address = value.String
			}
		case pinfo.FieldAbout:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field about", values[i])
			} else if value.Valid {
				pi.About = value.String
			}
		case pinfo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_user_has_p_info", value)
			} else if value.Valid {
				pi.user_user_has_p_info = new(int)
				*pi.user_user_has_p_info = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryWhoIsTheOwnerOfThisPInfo queries the "who_is_the_owner_of_this_PInfo" edge of the PInfo entity.
func (pi *PInfo) QueryWhoIsTheOwnerOfThisPInfo() *UserQuery {
	return (&PInfoClient{config: pi.config}).QueryWhoIsTheOwnerOfThisPInfo(pi)
}

// Update returns a builder for updating this PInfo.
// Note that you need to call PInfo.Unwrap() before calling this method if this PInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *PInfo) Update() *PInfoUpdateOne {
	return (&PInfoClient{config: pi.config}).UpdateOne(pi)
}

// Unwrap unwraps the PInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *PInfo) Unwrap() *PInfo {
	tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: PInfo is not a transactional entity")
	}
	pi.config.driver = tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *PInfo) String() string {
	var builder strings.Builder
	builder.WriteString("PInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", pi.ID))
	builder.WriteString(", profile=")
	builder.WriteString(pi.Profile)
	builder.WriteString(", idCardNumber=")
	builder.WriteString(pi.IdCardNumber)
	builder.WriteString(", prefix=")
	builder.WriteString(pi.Prefix)
	builder.WriteString(", firstName=")
	builder.WriteString(pi.FirstName)
	builder.WriteString(", lastName=")
	builder.WriteString(pi.LastName)
	builder.WriteString(", gender=")
	builder.WriteString(fmt.Sprintf("%v", pi.Gender))
	builder.WriteString(", brithDate=")
	builder.WriteString(pi.BrithDate.Format(time.ANSIC))
	builder.WriteString(", bloodGroup=")
	builder.WriteString(pi.BloodGroup)
	builder.WriteString(", address=")
	builder.WriteString(pi.Address)
	builder.WriteString(", about=")
	builder.WriteString(pi.About)
	builder.WriteByte(')')
	return builder.String()
}

// PInfos is a parsable slice of PInfo.
type PInfos []*PInfo

func (pi PInfos) config(cfg config) {
	for _i := range pi {
		pi[_i].config = cfg
	}
}
