package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Args = []string{"cmd", "./songjson/opensongmigrate/json"}
}
