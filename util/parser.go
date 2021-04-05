package util

import (
	"context"
	"encoding/json"
	"html"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/asaskevich/govalidator"
)

func ParseBodyData(ctx context.Context, r *http.Request, data interface{}) error {
	bBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bBody, data)
	if err != nil {
		return err
	}

	value := reflect.ValueOf(data).Elem()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		if field.Type() != reflect.TypeOf("") {
			continue
		}
		str := field.Interface().(string)

		//Escaping value to prevent SQL Injection
		field.SetString(html.EscapeString(str))

	}

	valid, err := govalidator.ValidateStruct(data)
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	return nil
}
