package double

import "testing"

type DummySearcher struct{}

func (ds DummySearcher) Search(people []*Person, firstName string, lastName string) *Person {
	return &Person{}
}

func TestFindItShouldReturnsErrorWhenFirstNameOrLastEmpty(t *testing.T) {
	phonebook := &Phonebook{}
	want := ErrMissingArgs

	dd := DummySearcher{}
	_, got := phonebook.Find(dd, "", "")

	if got != want {
		t.Errorf("Find() = %v; want %v", got, want)
	}

}
