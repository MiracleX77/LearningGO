package double

import "testing"

type SpySearcher struct {
	phone           string
	searchWasCalled bool
}

func (ss *SpySearcher) Search(people []*Person, firstName string, lastName string) *Person {
	ss.searchWasCalled = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "1234567890"
	phonebook := &Phonebook{}
	spy := &SpySearcher{phone: fakePhone}

	phone, _ := phonebook.Find(spy, "John", "Doe")

	if !spy.searchWasCalled {
		t.Error("Search() was not called")
	}
	if phone != fakePhone {
		t.Errorf("Find() = %v; want %v", phone, fakePhone)
	}
}
