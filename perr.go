package perr

import (
	"strings"
	"time"
)

type (
	Perror interface {
		Output() error
		Traces() PerrStack
		Level() ErrLevel
		ToDict() *ErrDict
	}

	Err struct {
		Inner        error
		Flag         uint32
		OccuredAt    time.Time
		textInternal string
		textOutput   string
		traces       StackTraces
	}

	ErrDict struct {
		Error     string      `json:"error"`
		Output    string      `json:"output"`
		Level     string      `json:"level"`
		Traces    StackTraces `json:"traces"`
		OccuredAt time.Time   `json:"occured_at"`
	}
)

func (e Err) Error() string {
	if e.Inner != nil {
		return e.Inner.Error()
	} else if e.textInternal != "" {
		return e.textInternal
	} else {
		return IntervalServerError.Error()
	}
}

func (e Err) Traces() StackTraces {
	return e.traces
}

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

func (e Err) ToDict() *ErrDict {
	return &ErrDict{
		Error:     e.Error(),
		Output:    e.Output().Error(),
		Level:     string(e.Level()),
		Traces:    e.traces,
		OccuredAt: e.OccuredAt,
	}
}

/* initialize perr */

func New(text string, flag uint32, output ...string) *Err {
	var out string
	if len(output) > 0 {
		out = strings.Join(output, ",")
	}

	return &Err{
		Flag:         flag,
		textInternal: text,
		textOutput:   out,
		OccuredAt:    time.Now(),
		traces:       NewTrace(callers()),
	}
}

func Wrap(in error, flag uint32, output ...string) *Err {
	var out string
	if len(output) > 0 {
		out = strings.Join(output, ",")
	}

	return &Err{
		Inner:      in,
		Flag:       flag,
		textOutput: out,
		OccuredAt:  time.Now(),
		traces:     NewTrace(callers()),
	}
}
