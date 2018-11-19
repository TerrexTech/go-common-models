package model

import (
	"github.com/TerrexTech/uuuid"
)

// Event refers to any change in system, such as insertion of some data.
type Event struct {
	// Action that was performed by the Event.
	Action string `cql:"action,omitempty" json:"action,omitempty"`

	// AggregateID for Aggregate associated with the Event.
	AggregateID int8 `cql:"aggregate_id,omitempty" json:"aggregateID,omitempty"`

	// CorrelationID can be used to track Event's source or reason why it as generated.
	CorrelationID uuuid.UUID `cql:"correlation_id,omitempty" json:"correlationID,omitempty"`

	// Data is the data contained by event.
	Data []byte `cql:"data,omitempty" json:"data,omitempty"`

	// NanoTime is the time in nanoseconds since Unix-epoch to when the event was generated.
	NanoTime int64 `cql:"nano_time,omitempty" json:"nanoTime,omitempty"`

	// Source is the name of service responsible for generating Event.
	Source string `cql:"source,omitempty" json:"source,omitempty"`

	// UserUUID is the V4-UUID of the user who generated the event.
	UserUUID uuuid.UUID `cql:"user_uuid,omitempty" json:"userUUID,omitempty"`

	// UUID is the V4-UUID unique-indentifier for event.
	UUID uuuid.UUID `cql:"uuid,omitempty" json:"uuid,omitempty"`

	// Version is the version for events as processed for aggregate-projection.
	// This is incremented by the aggregate itself each time it updates its state.
	Version int64 `cql:"version,omitempty" json:"version,omitempty"`

	// Year bucket is the year in which the event was generated.
	// This is used as the partitioning key.
	YearBucket int16 `cql:"year_bucket,omitempty" json:"yearBucket,omitempty"`
}
