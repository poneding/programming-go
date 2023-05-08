package handler

import (
	"api-gateway/pkg/util"
	"errors"
)

func PanicIfUserError(err error) {
	if err != nil {
		panic(errors.New("User service error:" + err.Error()))
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("taskService--" + err.Error())
		util.LogrusObject.Info(err)
		panic(err)
	}
}
