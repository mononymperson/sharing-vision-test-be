package common

import "github.com/go-playground/validator/v10"

func Validate(obj any) validator.ValidationErrors {
	v := validator.New()
	err, _ := v.Struct(obj).(validator.ValidationErrors)
	return err
}
