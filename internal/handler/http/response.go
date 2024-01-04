// Package http provides
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tingchima/gogolook/internal/domain/common"
)

// List .
type List struct {
	Data      any   `json:"data"`
	TotalSize int64 `json:"total_size"`
	NextPage  int   `json:"next_page"`
}

// ErrResponse .
type ErrResponse struct {
	Name    string `json:"name"`              // 錯誤名稱
	Message string `json:"message"`           // 錯誤訊息
	Details []any  `json:"details,omitempty"` // 錯誤細節
}

// responseToList .
func responseToList(data any, page, perPage int, totalSize int64) List {
	return List{
		Data:      data,
		TotalSize: totalSize,
		NextPage:  nextPage(page, perPage, totalSize),
	}
}

// nextPage .
func nextPage(page, perPage int, totalSize int64) int {
	if int64(page*perPage) < totalSize {
		return page + 1
	}
	return 0
}

// responseWithJSON .
func responseWithJSON(c *gin.Context, code int, resp any) {

	c.JSON(code, gin.H{"result": resp})
}

// responseWithNoContent .
func responseWithNoContent(c *gin.Context, code int) {
	c.Status(code)
}

// responseWithError .
func responseWithError(c *gin.Context, err error) {
	_ = c.Error(err)
	c.AbortWithStatusJSON(responseError(err))
}

func responseError(err error) (statusCode int, errResp ErrResponse) {

	var appErr *common.Error
	_ = common.AsErr(err, &appErr)

	if appErr == nil {
		appErr = common.NewError(common.ErrCodeInternalProcess, err, common.WithMsg(err.Error()))
	}

	causeErr := appErr.CauseErr()

	switch causeErr.(type) {
	case validator.ValidationErrors:
		fieldsErrs := causeErr.(validator.ValidationErrors)

		errDetails := make([]any, len(fieldsErrs))
		for i := range fieldsErrs {
			errDetails[i] = fieldsErrs[i].Error()
		}

		// append error detail messages
		appErr.WithDetails(errDetails...)
	}

	statusCode = appErr.HTTPStatus()

	errResp = ErrResponse{
		Name:    appErr.Name(),
		Message: appErr.ClientMsg(),
		Details: appErr.DetailMsg(),
	}

	return
}
