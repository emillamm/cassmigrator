package main

import (
	"testing"
	"os"
)

func TestCassandra(t *testing.T) {

	// Set up test session and defer closing the session

	host := os.Getenv("CASSANDRA_HOST")
	if host == "" {
		host = "localhost:9042"
	}

	t.Run("test NewCassandra returns authentication error if incorrect password is used", func(t *testing.T) {
		_, err := NewCassandra("cassandra", "invalid_password", host)
		if _, ok := err.(*InvalidCredentialsError); ok == false {
			t.Errorf("%#v was not the expected error. host %s", err, host)
		}
	})
	//t.Run("aa", func(t *testing.T) {
	//	if aaa(1) != "a" {
	//		t.Error("ab")
	//	}
	//})
}

