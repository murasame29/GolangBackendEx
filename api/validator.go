package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/murasame29/GolangBackendEx/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// 通貨がサポートされているか
		return util.IsSupportedCurrency(currency)
	}
	return false
}
