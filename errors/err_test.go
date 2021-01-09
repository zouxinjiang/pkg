package errors

import (
	"testing"
)

func TestErr(t *testing.T) {
	err := New("eee", "msg")
	t.Logf("%+v", err.RecursiveError())
}
