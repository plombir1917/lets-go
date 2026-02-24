package add

import "testing"

func TestAddPositive(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}

	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddZero(t *testing.T) {
	_, err := Add(0, 2)
	if err == nil {
		t.Error("first arg is zero  - expected error not be nil")
	}

	_, err = Add(1, 0)
	if err == nil {
		t.Error("second arg is zero  - expected error not be nil")
	}

	_, err = Add(0, 0)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestAddNegative(t *testing.T) {
	_, err := Add(-1, 1)
	if err != nil {
		t.Error("first are negative - expected error to be nil")
	}

	_, err = Add(1, -1)
	if err != nil {
		t.Error("second are negative - expected error to be nil")
	}

	_, err = Add(-1, -1)
	if err != nil {
		t.Error("both are negative - expected errors to be nil")
	}
}
