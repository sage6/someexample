package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}

func New(out io.Writer) Tracer {
	return &tracer{out: out}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprintln(t.out, a...)
	fmt.Println(a...)
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTracer{}
}
