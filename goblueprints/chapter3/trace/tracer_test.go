package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer

	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should nil")
	}

	tracer.Trace("hello")
	if buf.String() != "hello\n" {
		t.Errorf("Tracer should write: %s", buf.String())
	}
}

func TestOff(t *testing.T) {
	tracer := Off()
	tracer.Trace("something")
}
