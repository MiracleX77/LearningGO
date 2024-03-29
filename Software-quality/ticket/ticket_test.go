package ticket

import "testing"

type Case struct {
	name string
	age  int
	want float64
}

func TestTicketPrice(t *testing.T) {

	tests := []Case{
		{name: "should return 0 for age is -1", age: -1, want: 0.00},
		{name: "should return 0 for age is 0", age: 0, want: 0.00},
		{name: "should return 0 for age is 3", age: 3, want: 0.00},
		{name: "Ticket price $15 for age is 4", age: 4, want: 15.00},
		{name: "Ticket price $15 for age is 15", age: 15, want: 15.00},
		//{name: "Ticket price $30 for age is 16", age: 16, want: 30.00},
		//{name: "Ticket price $30 for age is 50", age: 50, want: 30.00},
		{name: "Ticket price $5 for age is 51", age: 51, want: 5.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)

			if got != tt.want {
				t.Errorf("Price(%d) = %f; want %f", tt.age, got, tt.want)
			}
		})
	}

}
