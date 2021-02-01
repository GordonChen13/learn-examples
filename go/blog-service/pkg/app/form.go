package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key   string
	Value string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Value
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Value)
	}
	return errs
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, nil
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:   key,
				Value: value,
			})
		}
		return false, errs
	}
	return true, nil
}
