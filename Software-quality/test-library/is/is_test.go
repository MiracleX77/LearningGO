package is

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func ParseBinary(b string) (bool, error) {
	return true, nil
}

func TestSomething(t *testing.T) {
	is := is.New(t)

	b, err := ParseBinary("1")

	is.NoErr(err)
	is.Equal(true, b)

	got := "asd"
	is.True(strings.Contains(got, "as"))
}
