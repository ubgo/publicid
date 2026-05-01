package publicid

import (
	"strings"
	"testing"
)

func TestNew_DefaultLength(t *testing.T) {
	id, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if len(id) != DefaultLength {
		t.Errorf("len: got %d, want %d", len(id), DefaultLength)
	}
	if strings.Trim(id, Alphabet) != "" {
		t.Errorf("invalid chars in id: %q", id)
	}
}

func TestMust_NeverPanicsForValidLength(t *testing.T) {
	id := Must()
	if len(id) != DefaultLength {
		t.Errorf("len: got %d, want %d", len(id), DefaultLength)
	}
}

func TestNewN_CustomLength(t *testing.T) {
	for _, n := range []int{1, 7, 12, 14, 24, 64} {
		id, err := NewN(n)
		if err != nil {
			t.Errorf("NewN(%d): %v", n, err)
			continue
		}
		if len(id) != n {
			t.Errorf("NewN(%d): got len %d", n, len(id))
		}
	}
}

func TestNewN_InvalidLength(t *testing.T) {
	for _, n := range []int{0, -1, -100} {
		if _, err := NewN(n); err == nil {
			t.Errorf("NewN(%d): expected error", n)
		}
	}
}

func TestMustN_PanicsOnInvalidLength(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("MustN(0): expected panic")
		}
	}()
	_ = MustN(0)
}

func TestNew_Uniqueness(t *testing.T) {
	const trials = 1000
	seen := make(map[string]struct{}, trials)
	for i := 0; i < trials; i++ {
		id := Must()
		if _, dup := seen[id]; dup {
			t.Fatalf("duplicate id at iteration %d: %q", i, id)
		}
		seen[id] = struct{}{}
	}
}

func TestValidate_OK(t *testing.T) {
	id := Must()
	if err := Validate("id", id); err != nil {
		t.Errorf("Validate: %v", err)
	}
}

func TestValidate_Empty(t *testing.T) {
	err := Validate("id", "")
	if err == nil || !strings.Contains(err.Error(), "blank") {
		t.Errorf("got %v, want blank error", err)
	}
}

func TestValidate_WrongLength(t *testing.T) {
	err := Validate("id", "abc")
	if err == nil || !strings.Contains(err.Error(), "characters long") {
		t.Errorf("got %v, want length error", err)
	}
}

func TestValidate_InvalidChars(t *testing.T) {
	id := strings.Repeat("Z", DefaultLength) // uppercase not in alphabet
	err := Validate("id", id)
	if err == nil || !strings.Contains(err.Error(), "invalid characters") {
		t.Errorf("got %v, want invalid-chars error", err)
	}
}

func TestValidateN_CustomLength(t *testing.T) {
	id, _ := NewN(7)
	if err := ValidateN("id", id, 7); err != nil {
		t.Errorf("ValidateN(7): %v", err)
	}
	if err := ValidateN("id", id, 24); err == nil {
		t.Errorf("ValidateN(24) on 7-char id should fail")
	}
}
