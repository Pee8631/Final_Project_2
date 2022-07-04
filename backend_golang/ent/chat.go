// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/chat"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Chat is the model entity for the Chat schema.
type Chat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ChatRoomName holds the value of the "Chat_room_name" field.
	ChatRoomName string `json:"Chat_room_name,omitempty"`
	// IsLockChat holds the value of the "IsLockChat" field.
	IsLockChat bool `json:"IsLockChat,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChatQuery when eager-loading is set.
	Edges ChatEdges `json:"edges"`
}

// ChatEdges holds the relations/edges for other nodes in the graph.
type ChatEdges struct {
	// ChatUser holds the value of the chat_user edge.
	ChatUser []*User `json:"chat_user,omitempty"`
	// ChatMessage holds the value of the chat_message edge.
	ChatMessage []*Message `json:"chat_message,omitempty"`
	// ChatAppointment holds the value of the chat_appointment edge.
	ChatAppointment []*Appointment `json:"chat_appointment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ChatUserOrErr returns the ChatUser value or an error if the edge
// was not loaded in eager-loading.
func (e ChatEdges) ChatUserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.ChatUser, nil
	}
	return nil, &NotLoadedError{edge: "chat_user"}
}

// ChatMessageOrErr returns the ChatMessage value or an error if the edge
// was not loaded in eager-loading.
func (e ChatEdges) ChatMessageOrErr() ([]*Message, error) {
	if e.loadedTypes[1] {
		return e.ChatMessage, nil
	}
	return nil, &NotLoadedError{edge: "chat_message"}
}

// ChatAppointmentOrErr returns the ChatAppointment value or an error if the edge
// was not loaded in eager-loading.
func (e ChatEdges) ChatAppointmentOrErr() ([]*Appointment, error) {
	if e.loadedTypes[2] {
		return e.ChatAppointment, nil
	}
	return nil, &NotLoadedError{edge: "chat_appointment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Chat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case chat.FieldIsLockChat:
			values[i] = new(sql.NullBool)
		case chat.FieldID:
			values[i] = new(sql.NullInt64)
		case chat.FieldChatRoomName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Chat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Chat fields.
func (c *Chat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case chat.FieldChatRoomName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Chat_room_name", values[i])
			} else if value.Valid {
				c.ChatRoomName = value.String
			}
		case chat.FieldIsLockChat:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IsLockChat", values[i])
			} else if value.Valid {
				c.IsLockChat = value.Bool
			}
		}
	}
	return nil
}

// QueryChatUser queries the "chat_user" edge of the Chat entity.
func (c *Chat) QueryChatUser() *UserQuery {
	return (&ChatClient{config: c.config}).QueryChatUser(c)
}

// QueryChatMessage queries the "chat_message" edge of the Chat entity.
func (c *Chat) QueryChatMessage() *MessageQuery {
	return (&ChatClient{config: c.config}).QueryChatMessage(c)
}

// QueryChatAppointment queries the "chat_appointment" edge of the Chat entity.
func (c *Chat) QueryChatAppointment() *AppointmentQuery {
	return (&ChatClient{config: c.config}).QueryChatAppointment(c)
}

// Update returns a builder for updating this Chat.
// Note that you need to call Chat.Unwrap() before calling this method if this Chat
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Chat) Update() *ChatUpdateOne {
	return (&ChatClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Chat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Chat) Unwrap() *Chat {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Chat is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Chat) String() string {
	var builder strings.Builder
	builder.WriteString("Chat(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", Chat_room_name=")
	builder.WriteString(c.ChatRoomName)
	builder.WriteString(", IsLockChat=")
	builder.WriteString(fmt.Sprintf("%v", c.IsLockChat))
	builder.WriteByte(')')
	return builder.String()
}

// Chats is a parsable slice of Chat.
type Chats []*Chat

func (c Chats) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}