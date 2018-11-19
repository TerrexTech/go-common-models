package model

// LogEntry describes "what's currently happening" for logging purposes.
type LogEntry struct {
	// Action describes "what was happening" when this entry was created.
	Action string `json:"action,omitempty"`

	// Description of what happened, can also be an error description.
	Description string `json:"description,omitempty"`

	// ErrorCode is just to inform the kind or classification of error.
	// The actual error should be part of description.
	ErrorCode int `json:"errorCode,omitempty"`

	// Level is the severity-level, that is, info, warning, or error.
	Level string `json:"level,omitempty"`

	// ServiceName is the service associated with the log.
	ServiceName string `json:"serviceName,omitempty"`
}
