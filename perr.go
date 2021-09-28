package perr

import (
	"errors"
	"strings"
	"time"
)

type (
	Perror interface {
		// error for client or response
		Output() error
		// get stacktrace
		Traces() PerrStack
		// get Perror level
		Level() ErrLevel
		// output ErrDict
		ToDict() *ErrDict
		// get cause
		Unwrap() error
		// whether Perror is caused by target
		Is(target error) bool
	}

	Err struct {
		cause           error
		Flag            uint32
		OccuredAt       time.Time
		msgForDeveloper string
		msgForClient    string
		traces          StackTraces
	}

	ErrDict struct {
		Error           error       `json:"error"`
		MsgForDeveloper string      `json:"msg_for_developer"`
		MsgForClient    string      `json:"msg_for_client"`
		Level           string      `json:"level"`
		Traces          StackTraces `json:"traces"`
		OccuredAt       time.Time   `json:"occured_at"`
	}
)

// get message for developer
// this method is for debug
func (e Err) Error() string {
	if e.cause != nil {
		return e.cause.Error()
	} else if e.msgForDeveloper != "" {
		return e.msgForDeveloper
	} else {
		return IntervalServerError.Error()
	}
}

// get stacktrace of Perror
func (e Err) Traces() StackTraces {
	return e.traces
}

// get Level of Perror
func (e Err) Level() ErrLevel {
	l := ErrLevelExternal
	flag := e.Flag
	if FlagInternalServerError <= flag && flag < FlagInternalServerErrorWithUrgency {
		l = ErrLevelInternal
	} else if flag == FlagInternalServerErrorWithUrgency {
		l = ErrLevelAlert
	}
	return l
}

// get cause of Perror
func (e Err) Unwrap() error {
	if e.cause != nil {
		return e.cause
	} else {
		return e.error()
	}
}

// whether Perror is caused by target
func (e Err) Is(target error) bool { return errors.Is(e.Unwrap(), target) }

// Convert Perror To ErrDict pointer
func (e Err) ToDict() *ErrDict {
	return &ErrDict{
		Error:           e.Unwrap(),
		MsgForDeveloper: e.Error(),
		MsgForClient:    e.Output().Error(),
		Level:           string(e.Level()),
		Traces:          e.traces,
		OccuredAt:       e.OccuredAt,
	}
}

/* initialize perr */

// initialize Perror
func New(msgForDeveloper string, flag uint32, msgForClient ...string) *Err {
	var out string
	if len(msgForDeveloper) > 0 {
		out = strings.Join(msgForClient, "\n")
	}

	return &Err{
		Flag:            flag,
		msgForDeveloper: msgForDeveloper,
		msgForClient:    out,
		OccuredAt:       time.Now(),
		traces:          NewTrace(callers()),
	}
}

// wrap error and initialize Perror
func Wrap(cause error, flag uint32, msgForClient ...string) *Err {
	if cause == nil {
		return nil
	}

	var out string
	if len(msgForClient) > 0 {
		out = strings.Join(msgForClient, ".")
	}

	return &Err{
		cause:        cause,
		Flag:         flag,
		msgForClient: out,
		OccuredAt:    time.Now(),
		traces:       NewTrace(callers()),
	}
}
