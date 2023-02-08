package sketch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func willError() (string, error) {
	return "", fmt.Errorf("this is an error")
}

func wontError() (string, error) {
	return "no error", nil
}

func TestUnwrapHappyString(t *testing.T) {
	val := Unwrap(wontError())
	assert.Equal(t, "no error", val)
}

func TestUnwrapPanic(t *testing.T) {
	assert.PanicsWithValue(t, "this is an error", func() { Unwrap(willError()) })
}

func TestCustUnwrapHappyString(t *testing.T) {
	val := (CustUnwrap[string]("value")(wontError()))
	assert.Equal(t, "no error", val)
}

func TestCustUnwrapPanic(t *testing.T) {
	assert.PanicsWithValue(t, "custom panic", func() { CustUnwrap[string]("custom panic")(willError()) })
}
