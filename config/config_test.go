package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 80
	t.Setenv("Port", fmt.Sprint(wantPort))

	got, err := New()

	if err != nil {
		t.Fatalf("Cannnot Create config: %v", err)
	}

	if got.Port != wantPort {
		t.Errorf("wanted port was not got")
	}

	wantEnv := "dev"

	if got.Env != wantEnv {
		t.Errorf("wanted env was not got")
	}

}
