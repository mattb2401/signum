package signum

import (
	"testing"
)

func TestMakeHash(t *testing.T) {
	h, err := MakeHash("s1gn@m", 14)
	if err != nil {
		t.Errorf("Make Hash failed error %v", err)
	}
	if !Attempt("s1gn@m", h) {
		t.Error("Incorrect hash created.")
	}
}

func TestAttempt(t *testing.T) {
	hashpass := "$2y$14$442tbg/BISPOGK81Eg337.Vcit9UzB/VZR9LItoOgDfJGPRIkZ6P."
	pass := "s1gn@m"
	if !Attempt(pass, hashpass) {
		t.Error("Hash not verified")
	}
}

func TestGetRounds(t *testing.T) {
	hashpass := "$2y$14$442tbg/BISPOGK81Eg337.Vcit9UzB/VZR9LItoOgDfJGPRIkZ6P."
	cost, err := GetRounds(hashpass)
	if err != nil {
		t.Errorf("Make Hash failed error %v", err)
	}
	if cost != 14 {
		t.Error("Incorrect cost of the hash")
	}
}
