package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/secretnote/backend/util"
)

var validTypeOfLogin validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if typeOfLogin, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedLogin(typeOfLogin)
	}
	return false
}
