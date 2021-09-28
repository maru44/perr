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
		cause           error
		As              error
		Level           ErrLevel
		OccuredAt       time.Time
		msgForDeveloper string
		msgForClient    string
		traces          StackTraces
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
		TreatedAs       string      `json:"teated_as"`
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
		return InternalServerError.Error()
	}
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
		Error:           e.Unwrap(),
		TreatedAs:       e.Output(),
		MsgForDeveloper: e.Error(),
		MsgForClient:    e.Output().Error(),
		Level:           string(e.Level),
		Traces:          e.traces,
		OccuredAt:       e.OccuredAt,
	}
}

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
	if e.cause != nil {
		return e.cause
	} else {
		return e.As
	}
}

// whether Perror is caused by target
func (e Err) Is(target error) bool { return errors.Is(e.Unwrap(), target) }

/* initialize perr */

// initialize Perror
func New(msgForDeveloper string, as error, msgForClient ...string) *Err {
	var out string
	if len(msgForDeveloper) > 0 {
		out = strings.Join(msgForClient, "\n")
	}

	return &Err{
		As:              as,
		Level:           getErrLevel(as),
		msgForDeveloper: msgForDeveloper,
		msgForClient:    out,
		OccuredAt:       time.Now(),
		traces:          NewTrace(callers()),
	}
}

// wrap error and initialize Perror
func Wrap(cause error, as error, msgForClient ...string) *Err {
	if cause == nil {
		return nil
	}

	var out string
	if len(msgForClient) > 0 {
		out = strings.Join(msgForClient, ".")
	}

	return &Err{
		cause:        cause,
		Level:        getErrLevel(as),
		As:           as,
		msgForClient: out,
		OccuredAt:    time.Now(),
		traces:       NewTrace(callers()),
	}
}

// initialize Perror with level
func NewWithLevel(msgForDeveloper string, as error, level ErrLevel, msgForClient ...string) *Err {
	var out string
	if len(msgForDeveloper) > 0 {
		out = strings.Join(msgForClient, "\n")
	}

	return &Err{
		As:              as,
		Level:           level,
		msgForDeveloper: msgForDeveloper,
		msgForClient:    out,
		OccuredAt:       time.Now(),
		traces:          NewTrace(callers()),
	}
}

// wrap error and initialize Perror with Level
func WrapWithLevel(cause error, as error, level ErrLevel, msgForClient ...string) *Err {
	if cause == nil {
		return nil
	}

	var out string
	if len(msgForClient) > 0 {
		out = strings.Join(msgForClient, ".")
	}

	return &Err{
		cause:        cause,
		Level:        level,
		As:           as,
		msgForClient: out,
		OccuredAt:    time.Now(),
		traces:       NewTrace(callers()),
	}
}
