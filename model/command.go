package model

import (
	"github.com/TerrexTech/uuuid"
)

// Command can be used to invoke a procedure in another service.
type Command struct {
	// Action for which the command is being produced.
	Action string `json:"action,omitempty"`

	// CorrelationID can be used to track Command's source or reason why it as generated.
	CorrelationID uuuid.UUID `json:"correlationID,omitempty"`

	// Data required for invoking the command.
	Data []byte `json:"data,omitempty"`

	// ResponseTopic is the Messaging-Topic on which the
	// response from processing the command should be sent.
	ResponseTopic string `json:"responseTopic,omitempty"`

	// Source is the name of the service which created the command.
	Source string `json:"source,omitempty"`

	// SourceTopic is the Messaging-Topic on which the Command was produced.
	SourceTopic string `json:"sourceTopic,omitempty"`

	// Timestamp is the time in Unix-milliseconds when this Command was generated.
	Timestamp int64 `json:"timestamp,omitempty"`

	// TTLSec is the Time-To-Live in seconds.
	// The Command should not be processed if current-time exceeds Timestamp+TTLSec.
	TTLSec int32 `json:"ttlSec,omitempty"`

	// UUID is the unique-identifier for Command.
	UUID uuuid.UUID `json:"uuid,omitempty"`
}
