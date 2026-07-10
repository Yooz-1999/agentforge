package apperrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "接口还没有实现")

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
