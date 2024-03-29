package sum

import "testing"

func TestSum(t *testing.T) {

	t.Run("Test sum(1, 2)", func(t *testing.T) {
		//Arrange
		want := 3

		//Act
		got := sum(1, 2)

		//Assert
		if got != want {
			t.Errorf("sum(1, 2) = %d; want %d", got, want)
		}
	})

	t.Run("Test sum(0, 0)", func(t *testing.T) {

		got := sum(0, 0)

		if got != 0 {
			t.Errorf("sum(0, 0) = %d; want 0", got)
		}
	})
	t.Run("Test sum(1, 0)", func(t *testing.T) {

		got := sum(1, 0)

		if got != 1 {
			t.Errorf("sum(1, 0) = %d; want 1", got)
		}
	})
}

func TestSumOne(t *testing.T) {
	got := sum(1, 0)

	if got != 1 {
		t.Errorf("sum(1, 0) = %d; want 1", got)
	}
}

func TestSumNegative(t *testing.T) {
	got := sum(-1, -1)

	if got != -2 {
		t.Errorf("sum(-1, -1) = %d; want -2", got)
	}
}
