# json-assert

A go library for complex nested assertions on JSON data types, inspired by Jest's [`objectContaining(object)`](https://jestjs.io/docs/expect#expectobjectcontainingobject) and other matchers .

## Installing

```
go get -u github.com/aarongodin/json-assert
```

## Usage

Basic usage:

```go
package main_test

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
  jsonAssert.AssertBytes(t, []byte(ex1), jsonAssert.Object(
    "test", 123
  ))
}
```

Executing the above will show a test failure:

```
--- FAIL: TestExample (0.00s)
    object.go:80: numbers must be equal; expected 13, received 123
```

