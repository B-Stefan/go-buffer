package api

import (
	"errors"
	"github.com/go-playground/form"
	"log"
	"net/url"
)

func getValuesWithoutEmpty(options interface{}) (url.Values, error) {

	encoder := form.NewEncoder()
	values, err := encoder.Encode(options)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	for key, value := range values {

		if len(value) == 0 {
			log.Panic(errors.New("Map has value with length of zero."))
		}
		first := value[0]
		if first != "0" && len(first) != 0 {
			params.Add(key, first)
		}
	}
	return params, err
}
