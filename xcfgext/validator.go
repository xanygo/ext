package xcfgext

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/xanygo/anygo/xcfg"
)

var _ xcfg.Validator = (*validatorV10)(nil)

func init() {
	xcfg.DefaultValidator = newValidatorV10()
}

func newValidatorV10() *validatorV10 {
	return &validatorV10{
		vv: validator.New(),
	}
}

type validatorV10 struct {
	vv *validator.Validate
}

func (v *validatorV10) Validate(val any) error {
	if v.vv == nil {
		return nil
	}
	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Struct:
		return v.vv.Struct(val)
	case reflect.Ptr:
		rvv := rv.Elem()
		switch rvv.Kind() {
		case reflect.Ptr:
			return v.vv.Struct(rvv.Interface())
		case reflect.Struct:
			return v.vv.Struct(val)
		}
	}
	return nil
}
