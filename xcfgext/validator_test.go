package xcfgext

import (
	"testing"

	"github.com/xanygo/anygo/xcfg"
	"github.com/xanygo/anygo/xt"
)

func TestValidator(t *testing.T) {
	type user struct {
		Name string `validate:"required"`
		Age  int
	}

	t.Run("validator-1", func(t *testing.T) {
		var u *user
		err := xcfg.ParseBytes(".json", []byte(`{"Age":12}`), &u)
		xt.Equal(t, &user{Age: 12}, u)
		xt.Error(t, err)
	})

	t.Run("validator-2", func(t *testing.T) {
		var u *user
		err := xcfg.ParseBytes(".json", []byte(``), &u)
		xt.Nil(t, u)
		xt.Error(t, err)
	})

	t.Run("validator-3", func(t *testing.T) {
		var u *user
		err := xcfg.ParseBytes(".json", []byte(`{"Age":12,"Name":""}`), &u)
		xt.Equal(t, &user{Age: 12}, u)
		xt.Error(t, err)
	})

	t.Run("validator-4", func(t *testing.T) {
		var u *user
		err := xcfg.ParseBytes(".json", []byte(`{"Age":12,"Name":"hello"}`), &u)
		xt.Equal(t, &user{Age: 12, Name: "hello"}, u)
		xt.NoError(t, err)
	})

	t.Run("validator-5", func(t *testing.T) {
		u := &user{}
		err := xcfg.ParseBytes(".json", []byte(`{"Age":12,"Name":"hello"}`), u)
		xt.Equal(t, &user{Age: 12, Name: "hello"}, u)
		xt.NoError(t, err)
	})

	t.Run("validator-6", func(t *testing.T) {
		u := &user{}
		err := xcfg.ParseBytes(".json", []byte(`{"Age":12,"Name":""}`), u)
		xt.Equal(t, &user{Age: 12}, u)
		xt.Error(t, err)
	})
}
