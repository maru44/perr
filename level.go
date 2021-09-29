package perr

// "ErrLevel" show error level
type ErrLevel string

const (
	// internal with urgency
	ErrLevelAlert ErrLevel = "ALERT"
	// internal nwithout urgency
	ErrLevelInternal ErrLevel = "INTERNAL ERROR"
	// external (client)
	ErrLevelExternal ErrLevel = "EXTERNAL ERROR"
)
