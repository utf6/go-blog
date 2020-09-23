package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/utf6/go-blog/pkg/logs"
)

func MarkErrors(errors []*validation.Error)  {
	for _, err := range errors {
		logs.Info(err.Key, err.Message)
	}

	return
}
