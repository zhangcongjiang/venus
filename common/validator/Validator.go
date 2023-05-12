package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	match, _ := regexp.MatchString(`^1[3456789]\d-\d{4}-\d{4}$`, phone)
	return match
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("phone", ValidatePhone)
		if err != nil {
			return
		}
	}

}
