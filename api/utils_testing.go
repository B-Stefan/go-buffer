package api

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"reflect"
)

func tryJsonEncoding(jsonStr string, model interface{}) error {

	err := json.Unmarshal([]byte(jsonStr), &model)

	if err != nil {
		return err
	}

	byteOut, err := json.Marshal(&model)

	if err != nil {
		return err
	}

	eq, err := JSONBytesEqual(byteOut, []byte(jsonStr))

	fmt.Println("a=c\t", eq, "with error", err)

	if err != nil {
		fmt.Printf("%+v\n", model)
		return err
	}

	if !eq {
		fmt.Printf("%+v\n", model)
		fmt.Printf("%+v\n", string(byteOut))
		return errors.New("Json documents are not equal")
	}
	return nil
}

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) do(req *http.Request, v interface{}) (*http.Response, error) {
	args := c.Called()

	err := copier.Copy(v, args.Get(0))
	res := &http.Response{
		StatusCode: 200,
		// Must be set to non-nil value or it panics
		Header: make(http.Header),
	}
	return res, err
}

func (c *ClientMock) newRequest(method, path string, body interface{}) (*http.Request, error) {
	c.Called(method, path, body)
	req := &http.Request{}
	return req, nil
}
