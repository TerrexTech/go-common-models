package model

import "github.com/TerrexTech/uuuid"

// Document can be used to transfer data between services.
type Document struct {
	// CorrelationID can be used to track Document's source or reason why it as generated.
	CorrelationID uuuid.UUID `json:"correlationID,omitempty"`

	// Data contained by Document.
	Data []byte `json:"data,omitempty"`

	// Error occurred while processing the Input.
	// Convert errors to strings, this is just an indication that
	// something went wrong, so we can signal/display-error to end-
	// user. Blank Error-string means everything was fine.
	Error string `json:"error,omitempty"`

	// ErrorCode can be used to identify type of error.
	ErrorCode int16 `json:"errorCode,omitempty"`

	// Source is the name of the service which created the command.
	Source string `json:"source,omitempty"`

	// Topic on which the document should be produced.
	Topic string `json:"topic,omitempty"`

	// UUID is the V4-UUID Document-Identifier.
	UUID uuuid.UUID `json:"uuid,omitempty"`
}
