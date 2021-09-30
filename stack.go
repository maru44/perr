package perr

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

type (
	stackTraces []stackTrace

	stackTrace struct {
		File           string  `json:"file"`
		Line           int     `json:"line"`
		Name           string  `json:"name"`
		ProgramCounter uintptr `json:"program_counter"`
		// cause will be 0
		// Every time you stack perr.Wrap Layer will be increased by 1.
		Layer int `json:"layer"`
	}
)

/* traces method */

// output stacktace for string
func (ss stackTraces) String() string {
	var buf bytes.Buffer
	for _, s := range ss {
		fmt.Fprintf(&buf, "%s%s:%d ===> %v\n", strings.Repeat("\t", s.Layer), s.File, s.Line, s.Name)
	}
	return buf.String()
}

// getMaxLayer
func (ss stackTraces) maxLayer() (max int) {
	for _, s := range ss {
		if s.Layer > max {
			max = s.Layer
		}
	}
	return max
}

/* init trace */

// ref: https://github.com/pkg/errors/blob/816c9085562cd7ee03e7f8188a1cfd942858cded/stack.go#L133
func callers() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[0 : n-2]
}

func newTrace(pcs []uintptr) stackTraces {
	traces := make([]stackTrace, len(pcs))

	for i, pc := range pcs {
		trace := stackTrace{ProgramCounter: pc}
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			return traces
		}
		trace.Name = funcname(fn.Name())
		trace.File, trace.Line = fn.FileLine(pc - 1)
		traces[i] = trace
	}
	return traces
}

func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}
