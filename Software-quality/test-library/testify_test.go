package testlibrary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

func TestSomething(t *testing.T) {
	t.Run("should return 0 for age is -1", func(t *testing.T) {
		pp := &Person{FirstName: "asd"}

		if assert.NotNil(t, pp) {
			assert.Equal(t, "asd", pp.FirstName)
		}
	})
}
