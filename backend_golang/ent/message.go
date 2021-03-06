// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/chat"
	"FinalProject/ent/message"
	"FinalProject/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MessageText holds the value of the "message_text" field.
	MessageText string `json:"message_text,omitempty"`
	// SentDateTime holds the value of the "sent_dateTime" field.
	SentDateTime time.Time `json:"sent_dateTime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageQuery when eager-loading is set.
	Edges                  MessageEdges `json:"edges"`
	chat_chat_message      *int
	user_user_send_message *int
}

// MessageEdges holds the relations/edges for other nodes in the graph.
type MessageEdges struct {
	// WhatMessagesAreInThisChat holds the value of the What_messages_are_in_this_chat edge.
	WhatMessagesAreInThisChat *Chat `json:"What_messages_are_in_this_chat,omitempty"`
	// WhoSendMessages holds the value of the Who_send_messages edge.
	WhoSendMessages *User `json:"Who_send_messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// WhatMessagesAreInThisChatOrErr returns the WhatMessagesAreInThisChat value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) WhatMessagesAreInThisChatOrErr() (*Chat, error) {
	if e.loadedTypes[0] {
		if e.WhatMessagesAreInThisChat == nil {
			// The edge What_messages_are_in_this_chat was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: chat.Label}
		}
		return e.WhatMessagesAreInThisChat, nil
	}
	return nil, &NotLoadedError{edge: "What_messages_are_in_this_chat"}
}

// WhoSendMessagesOrErr returns the WhoSendMessages value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) WhoSendMessagesOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.WhoSendMessages == nil {
			// The edge Who_send_messages was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.WhoSendMessages, nil
	}
	return nil, &NotLoadedError{edge: "Who_send_messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			values[i] = new(sql.NullInt64)
		case message.FieldMessageText:
			values[i] = new(sql.NullString)
		case message.FieldSentDateTime:
			values[i] = new(sql.NullTime)
		case message.ForeignKeys[0]: // chat_chat_message
			values[i] = new(sql.NullInt64)
		case message.ForeignKeys[1]: // user_user_send_message
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Message", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case message.FieldMessageText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message_text", values[i])
			} else if value.Valid {
				m.MessageText = value.String
			}
		case message.FieldSentDateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sent_dateTime", values[i])
			} else if value.Valid {
				m.SentDateTime = value.Time
			}
		case message.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field chat_chat_message", value)
			} else if value.Valid {
				m.chat_chat_message = new(int)
				*m.chat_chat_message = int(value.Int64)
			}
		case message.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_user_send_message", value)
			} else if value.Valid {
				m.user_user_send_message = new(int)
				*m.user_user_send_message = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryWhatMessagesAreInThisChat queries the "What_messages_are_in_this_chat" edge of the Message entity.
func (m *Message) QueryWhatMessagesAreInThisChat() *ChatQuery {
	return (&MessageClient{config: m.config}).QueryWhatMessagesAreInThisChat(m)
}

// QueryWhoSendMessages queries the "Who_send_messages" edge of the Message entity.
func (m *Message) QueryWhoSendMessages() *UserQuery {
	return (&MessageClient{config: m.config}).QueryWhoSendMessages(m)
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return (&MessageClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", message_text=")
	builder.WriteString(m.MessageText)
	builder.WriteString(", sent_dateTime=")
	builder.WriteString(m.SentDateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message

func (m Messages) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
