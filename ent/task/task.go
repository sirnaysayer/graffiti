// Code generated by entc, DO NOT EDIT.

package task

import (
	"fmt"
	"graffiti/ent/schema"
	"time"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActivity holds the string denoting the activity vertex property in the database.
	FieldActivity = "activity"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"
	// FieldState holds the string denoting the state vertex property in the database.
	FieldState = "state"

	// Table holds the table name of the task in the database.
	Table = "tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldActivity,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldState,
}

var (
	fields = schema.Task{}.Fields()

	// descCreatedAt is the schema descriptor for created_at field.
	descCreatedAt = fields[1].Descriptor()
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt = descCreatedAt.Default.(func() time.Time)

	// descUpdatedAt is the schema descriptor for updated_at field.
	descUpdatedAt = fields[2].Descriptor()
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt = descUpdatedAt.Default.(func() time.Time)
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt = descUpdatedAt.UpdateDefault.(func() time.Time)
)

// State defines the type for the state enum field.
type State string

// StateUndone is the default State.
const DefaultState = StateUndone

// State values.
const (
	StateDone   State = "done"
	StateUndone State = "undone"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "s" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateDone, StateUndone:
		return nil
	default:
		return fmt.Errorf("task: invalid enum value for state field: %q", s)
	}
}