package jsonAssert_test

import (
	"testing"

	jsonAssert "github.com/aarongodin/json-assert"
)

var ex1 = `
{
	"test": 100
}
`

func TestExample(t *testing.T) {
	jsonAssert.AssertBytes(t, []byte(ex1), jsonAssert.Object("test", 123))
}
