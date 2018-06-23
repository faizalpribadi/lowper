package lowper

import (
	"errors"
	"strings"
)

var (
	NullableString = errors.New("Can't be defined nullable string")
)

func ToUpperCase(param string, override bool) (string, error) {
	if !override {
		override = true
	}

	if param == "" {
		return "", NullableString
	}

	return strings.ToUpper(param), nil
}

func ToLowerCase(param string, override bool) (string, error) {
	if !override {
		override = true
	}

	if param == "" {
		return "", NullableString
	}

	return strings.ToLower(param), nil
}
