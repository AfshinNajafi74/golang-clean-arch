package validations

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IranianMobileValidator(flt validator.FieldLevel) bool {
	value, ok := flt.Field().Interface().(string)
	if !ok {
		return false
	}

	// TODO: fill regexp
	result, err := regexp.MatchString(``, value)
	if err != nil {
		log.Print(err.Error())
	}
	return result

}
