package main

import (
	"testing"
	"reflect"
)

func TestGoCqlUtil(t *testing.T) {
	t.Run("test ParseErrorString InvalidCredentialsError", func(t *testing.T) {
		errStr := "gocql: unable to create session: unable to discover protocol version: Provided username test_role and/or password are incorrect"
		got := ParseErrorString(errStr)
		want := &InvalidCredentialsError{errStr}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%#v is different from %#v", got, want)
		}
	})
}

