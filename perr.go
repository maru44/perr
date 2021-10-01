package perr

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"time"
)

type (
	Perror interface {
		// error for client or response
		Output() error
		// get level
		Level() ErrLevel
		// get stacktrace
		Traces() stackTraces
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
		as           error
		level        ErrLevel
		msgForClient string
		traces       stackTraces
		occurredAt   time.Time
	}

	ErrDict struct {
		Error        error
		TreatedAs    error
		MsgForClient string
		Level        string
		Traces       stackTraces
		OccurredAt   time.Time
	}

	errDictJson struct {
		Error        string      `json:"error"`
		TreatedAs    string      `json:"treated_as"`
		MsgForClient string      `json:"msg_for_client"`
		Level        string      `json:"level"`
		Traces       stackTraces `json:"traces"`
		OccurredAt   time.Time   `json:"occurred_at"`
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

	return e.as
}

// Get level
func (e Err) Level() ErrLevel {
	return e.level
}

// get stacktrace of Perror
func (e Err) Traces() stackTraces {
	return e.traces
}

// Convert Perror To ErrDict pointer
func (e Err) Map() *ErrDict {
	return &ErrDict{
		Error:        e.Unwrap(),
		TreatedAs:    e.as,
		MsgForClient: e.Output().Error(),
		Level:        string(e.level),
		Traces:       e.traces,
		OccurredAt:   e.occurredAt,
	}
}

// Convert Perror TO Json
func (e Err) Json() []byte {
	m := e.Map()
	j := errDictJson{
		Error:        m.Error.Error(),
		TreatedAs:    m.TreatedAs.Error(),
		MsgForClient: m.MsgForClient,
		Level:        m.Level,
		Traces:       m.Traces,
		OccurredAt:   m.OccurredAt,
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
		// if not blank
		cause = errors.New(errString)

		if as == nil {
			as = cause
		}
	} else {
		cause = as
	}

	if cause == nil {
		return nil
	}

	return &Err{
		cause:        cause,
		as:           as,
		level:        getErrLevel(as),
		msgForClient: out,
		occurredAt:   time.Now(),
		traces:       newTrace(callers()),
	}
}

// "Wrap" function wrap error and initialize *perr.Err.
// If `cause` is nil return nil.
// if as is nil, error for client will be `cause`.
func Wrap(cause error, as error, msgForClient ...string) *Err {
	if cause == nil || reflect.ValueOf(cause).IsNil() {
		return nil
	}

	var out string
	var traces stackTraces
	if perror, ok := cause.(Perror); ok {
		as = perror.Output()
		traces = perror.Traces()
		max := traces.maxLayer()
		// for _, t := range newTrace(callers()) {
		// 	t.Layer = max
		// 	traces = append(traces, t)
		// }

		t := newTrace(callers())[0]
		t.Layer = max + 1
		traces = append(traces, t)
	} else {
		if as == nil {
			as = cause
		}
		traces = newTrace(callers())
	}

	// overwrite msgForClient
	if len(msgForClient) > 0 {
		out = strings.Join(msgForClient, "\n")
	}

	return &Err{
		cause:        cause,
		level:        getErrLevel(as),
		as:           as,
		msgForClient: out,
		occurredAt:   time.Now(),
		traces:       traces,
	}
}

// initialize Perror with level
func NewWithLevel(errString string, as error, level ErrLevel, msgForClient ...string) *Err {
	p := New(errString, as, msgForClient...)
	p.level = level
	return p
}

// wrap error and initialize Perror with Level
func WrapWithLevel(cause error, as error, level ErrLevel, msgForClient ...string) *Err {
	p := Wrap(cause, as, msgForClient...)
	p.level = level
	return p
}
