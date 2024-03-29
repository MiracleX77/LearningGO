package teardown

import "testing"

func setup(t *testing.T) func() {
	t.Log("Before all tests")
	return func() {
		t.Log("After all tests")
	}
}

func TestTeardown(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	t.Run("Test 1", func(t *testing.T) {
		t.Log("Test 1")
	})

	t.Run("Test 2", func(t *testing.T) {
		t.Log("Test 2")
	})
}
