package perr

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

type (
	PerrStack interface {
		String() string
	}

	StackTraces []StackTrace

	StackTrace struct {
		File           string  `json:"file"`
		Line           int     `json:"line"`
		Name           string  `json:"name"`
		ProgramCounter uintptr `json:"program_counter"`
	}
)

/* traces method */

func (ss StackTraces) String() string {
	var buf bytes.Buffer
	for _, s := range ss {
		fmt.Fprintf(&buf, "%s:%d ===> %v\n", s.File, s.Line, s.Name)
	}
	return buf.String()
}

/* init trace */

// ref: https://github.com/pkg/errors/blob/816c9085562cd7ee03e7f8188a1cfd942858cded/stack.go#L133
func callers() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[0 : n-2]
}

func NewTrace(pcs []uintptr) StackTraces {
	traces := make([]StackTrace, len(pcs))

	for i, pc := range pcs {
		trace := StackTrace{ProgramCounter: pc}
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			return traces
		}
		trace.Name = trimPkgName(fn)
		trace.File, trace.Line = fn.FileLine(pc - 1)
		traces[i] = trace
	}
	return traces
}

func trimPkgName(fn *runtime.Func) string {
	name := fn.Name()
	if ld := strings.LastIndex(name, "."); ld >= 0 {
		name = name[ld+1:]
	}
	return name
}
