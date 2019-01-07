package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestOptions struct {
	MyOptions       string
	MyOptionsInt    int
	MyOptionsInt64  int64
	MyOptionsNested struct {
		Name string
	}
}

func TestGetValuesWithoutEmptyInt(t *testing.T) {

	values, err := getValuesWithoutEmpty(TestOptions{
		MyOptions:    "dddd",
		MyOptionsInt: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "MyOptions=dddd", values.Encode())
}

func TestGetValuesWithoutEmptyString(t *testing.T) {

	values, err := getValuesWithoutEmpty(TestOptions{
		MyOptions:    "",
		MyOptionsInt: 3,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "MyOptionsInt=3", values.Encode())
}
