package perr

import (
	"encoding/json"
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
		// output ErrDict
		Map() *ErrDict
		// ErrDict >> json
		Json() []byte
		// get cause
		Unwrap() error
		// whether Perror is caused by target
		Is(target error) bool
	}

	Err struct {
		cause        error
		As           error
		Level        ErrLevel
		OccuredAt    time.Time
		msgForClient string
		traces       StackTraces
	}

	ErrDict struct {
		Error           error
		TreatedAs       error
		MsgForDeveloper string
		MsgForClient    string
		Level           string
		Traces          StackTraces
		OccuredAt       time.Time
	}

	errDictJson struct {
		Error           string      `json:"error"`
		TreatedAs       string      `json:"treated_as"`
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
	return e.cause.Error()
}

// get message for client
func (e Err) Output() error {
	if e.msgForClient != "" {
		return errors.New(e.msgForClient)
	}

	return e.As
}

// get stacktrace of Perror
func (e Err) Traces() StackTraces {
	return e.traces
}

// Convert Perror To ErrDict pointer
func (e Err) Map() *ErrDict {
	return &ErrDict{
		Error:        e.Unwrap(),
		TreatedAs:    e.As,
		MsgForClient: e.Output().Error(),
		Level:        string(e.Level),
		Traces:       e.traces,
		OccuredAt:    e.OccuredAt,
	}
}

// Convert Perror TO Json
func (e Err) Json() []byte {
	m := e.Map()
	j := errDictJson{
		Error:           m.Error.Error(),
		TreatedAs:       m.TreatedAs.Error(),
		MsgForDeveloper: m.MsgForDeveloper,
		MsgForClient:    m.MsgForClient,
		Level:           m.Level,
		Traces:          m.Traces,
		OccuredAt:       m.OccuredAt,
	}
	json_, err := json.Marshal(j)
	if err != nil {
		return nil
	}
	return json_
}

// get cause of Perror
func (e Err) Unwrap() error {
	return e.cause
}

// whether Perror is caused by target
func (e Err) Is(target error) bool { return errors.Is(e.Unwrap(), target) }

/* initialize perr */

// initialize Perror
func New(errString string, as error, msgForClient ...string) *Err {
	var out string
	if len(msgForClient) > 0 {
		out = strings.Join(msgForClient, "\n")
	}

	var cause error
	if errString != "" {
		cause = errors.New(errString)
	} else {
		cause = as
	}

	if as == nil {
		as = cause
	}

	return &Err{
		cause:        cause,
		As:           as,
		Level:        getErrLevel(as),
		msgForClient: out,
		OccuredAt:    time.Now(),
		traces:       newTrace(callers()),
	}
}

// wrap error and initialize Perror
func Wrap(cause error, as error, msgForClient ...string) *Err {
	if cause == nil {
		return nil
	}

	var out string
	if len(msgForClient) > 0 {
		out = strings.Join(msgForClient, "\n")
	}

	if as == nil {
		as = cause
	}

	return &Err{
		cause:        cause,
		Level:        getErrLevel(as),
		As:           as,
		msgForClient: out,
		OccuredAt:    time.Now(),
		traces:       newTrace(callers()),
	}
}

// initialize Perror with level
func NewWithLevel(errString string, as error, level ErrLevel, msgForClient ...string) *Err {
	p := New(errString, as, msgForClient...)
	p.Level = level
	return p
}

// wrap error and initialize Perror with Level
func WrapWithLevel(cause error, as error, level ErrLevel, msgForClient ...string) *Err {
	p := Wrap(cause, as, msgForClient...)
	p.Level = level
	return p
}
