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
		As              error
		OccuredAt       time.Time
		msgForDeveloper string
		msgForClient    string
		traces          StackTraces
	}

	ErrDict struct {
		Error           error       `json:"error"`
		TreatedAs       error       `json:"teated_as"`
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

// get stacktrace of Perror
func (e Err) Traces() StackTraces {
	return e.traces
}

// get Level of Perror
func (e Err) Level() ErrLevel {
	var l ErrLevel
	switch e.As {
	case InternalServerErrorWithUrgency, NotExtendedWithUrgency, BadGatewayWithUrgency,
		ServiceUnavailableWithUrgency, GatewayTimeoutWithUrgency, HTTPVersionNotSupportedWithUrgency,
		VariantAlsoNegotiatesWithUrgency, InsufficientStorageWithUrgency, LoopDetectedWithUrgency,
		NotExtendedWithUrgency, NetworkAuthenticationRequiredWithUrgency:
		l = ErrLevelAlert
	case InternalServerError, NotImplemented, BadGateway, ServiceUnavailable,
		GatewayTimeout, HTTPVersionNotSupported, VariantAlsoNegotiates,
		InsufficientStorage, LoopDetected, NotExtended, NetworkAuthenticationRequired:
		l = ErrLevelInternal
	default:
		l = ErrLevelExternal
	}
	return l
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

// Convert Perror To ErrDict pointer
func (e Err) ToDict() *ErrDict {
	return &ErrDict{
		Error:           e.Unwrap(),
		TreatedAs:       e.Output(),
		MsgForDeveloper: e.Error(),
		MsgForClient:    e.Output().Error(),
		Level:           string(e.Level()),
		Traces:          e.traces,
		OccuredAt:       e.OccuredAt,
	}
}

/* initialize perr */

// initialize Perror
func New(msgForDeveloper string, as error, msgForClient ...string) *Err {
	var out string
	if len(msgForDeveloper) > 0 {
		out = strings.Join(msgForClient, "\n")
	}

	return &Err{
		As:              as,
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
		As:           as,
		msgForClient: out,
		OccuredAt:    time.Now(),
		traces:       NewTrace(callers()),
	}
}
