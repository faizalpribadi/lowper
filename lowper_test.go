package lowper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Params struct {
	Upper string
	Lower string
}

func New() *Params {
	return &Params{
		Upper: "string to upper case",
		Lower: "STRING TO LOWER CASE",
	}
}

func TestUpperCase(t *testing.T) {
	param := New()
	val, err := ToUpperCase(param.Upper, true)
	if err != nil {
		t.Errorf("%s", err)
	}

	assert.True(t, true, val)
}

func TestLowerCase(t *testing.T) {
	param := New()
	val, err := ToLowerCase(param.Lower, true)
	if err != nil {
		t.Errorf("%s", err)
	}

	assert.True(t, true, val)
}
