package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/techschool/simplebank/db/util"
)

var validCurrency validator.Func = func(feildLevel validator.FieldLevel) bool {
	if currency, ok := feildLevel.Field().Interface().(string); ok {
		//check if currency is supported or not

		return util.IsSupportedCurrency(currency)
	}

	return false
}
