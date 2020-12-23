package utilities

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"gopkg.in/validator.v2"
	"io"
	"reflect"
	"time"
)

var decoder *schema.Decoder
var timeConverter = func(value string) reflect.Value {
	if v, err := time.Parse("2006-02-01", value); err == nil {
		return reflect.ValueOf(v)
	}
	return reflect.Value{} // this is the same as the private const invalidType
}

func required(v interface{}, param string) error {
	return nil
}
func DecodeBody(io io.Reader, v interface{}) error {
	dec := json.NewDecoder(io)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&v); err != nil {
		return err
	}
	if errs := validator.Validate(v); errs != nil {
		return errs
	}
	return nil

}
func DecodeQuery(query map[string][]string, v interface{}) error {
	if err := decoder.Decode(v, query); err != nil {
		return err
	}
	if errs := validator.Validate(v); errs != nil {

		return errs
	}
	return nil
}
func init() {
	decoder = schema.NewDecoder()
	decoder.RegisterConverter(time.Time{}, timeConverter)
	validator.SetValidationFunc("required", required)
}
