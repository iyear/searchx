package validator

import (
	"github.com/go-playground/validator/v10"
	iso6391 "github.com/iyear/iso-639-1"
	"log"
)

var v *validator.Validate

func init() {
	v = validator.New()

	err := v.RegisterValidation("iso6391", func(fl validator.FieldLevel) bool {
		return iso6391.ValidCode(fl.Field().String())
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Struct(s interface{}) error {
	return v.Struct(s)
}
