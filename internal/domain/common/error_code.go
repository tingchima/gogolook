// Package common provides
package common

import "net/http"

// ErrCode .
type ErrCode struct {
	Name       string
	StatusCode int
}

/*
	400
*/

// ErrCodeInvalidParameter .
var ErrCodeInvalidParameter = ErrCode{
	Name:       "INVALID_PARAMETER",
	StatusCode: http.StatusBadRequest,
}

/*
	401
*/

// ErrCodeUnauthorized .
var ErrCodeUnauthorized = ErrCode{
	Name:       "UNAUTHORIZED",
	StatusCode: http.StatusUnauthorized,
}

/*
	403
*/

// ErrCodeAccessNotAllowed .
var ErrCodeAccessNotAllowed = ErrCode{
	Name:       "ACCESS_NOT_ALLOWED",
	StatusCode: http.StatusForbidden,
}

/*
	404
*/

// ErrCodeResourceNotFound .
var ErrCodeResourceNotFound = ErrCode{
	Name:       "RESOURCE_NOT_FOUND",
	StatusCode: http.StatusNotFound,
}

/*
	409
*/

// ErrCodeResourceAlreadyExisted .
var ErrCodeResourceAlreadyExisted = ErrCode{
	Name:       "RESOURCE_ALREADY_EXISTED",
	StatusCode: http.StatusConflict,
}

/*
	500
*/

// ErrCodeInternalProcess .
var ErrCodeInternalProcess = ErrCode{
	Name:       "INTERNAL_PROCESS",
	StatusCode: http.StatusInternalServerError,
}
