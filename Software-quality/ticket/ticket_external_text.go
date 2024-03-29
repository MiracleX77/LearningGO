package ticket_test

import "testing"

func TestTicket(t *testing.T) {
	t.Run("should return 0 for age is -1", func(t *testing.T) {
		age := -1
		want := 0.00

		got := Price(age)

		if got != want {
			t.Errorf("Price(%d) = %f; want %f", age, got, want)
		}
	})
}
