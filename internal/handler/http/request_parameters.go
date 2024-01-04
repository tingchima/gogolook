// Package http provides
package http

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tingchima/gogolook/internal/domain/common"
)

func GetPathInt(c *gin.Context, name string) (int, error) {
	strVal := c.Params.ByName(name)
	if strVal == "" {
		msg := fmt.Sprintf("the %s from path parameter value is empty or not specified", name)
		return 0, common.NewError(common.ErrCodeInvalidParameter, errors.New(msg), common.WithMsg(msg))
	}

	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		msg := fmt.Sprintf("the %s from path parameter value is invalid", name)
		return 0, common.NewError(common.ErrCodeInvalidParameter, errors.New(msg), common.WithMsg(msg))
	}

	return intVal, nil
}
