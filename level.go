package perr

type ErrLevel string

const (
	/* error level */

	// internal with urgency
	ErrLevelAlert ErrLevel = "ALERT"
	// internal nwithout urgency
	ErrLevelInternal ErrLevel = "INTERNAL ERROR"
	// external (client)
	ErrLevelExternal ErrLevel = "EXTERNAL ERROR"
)
