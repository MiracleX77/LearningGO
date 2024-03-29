package double

import "testing"

type StubSearcher struct {
	Phone string
}

func (ss StubSearcher) Search(people []*Person, firstName string, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.Phone,
	}
}

func TestFindReturnsPerson(t *testing.T) {
	fakePhone := "1234567890"
	phonebook := &Phonebook{}
	ss := StubSearcher{Phone: fakePhone}

	phone, _ := phonebook.Find(ss, "John", "Doe")

	if phone != fakePhone {
		t.Errorf("Find() = %v; want %v", phone, fakePhone)
	}
}
