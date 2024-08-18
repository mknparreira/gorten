package errors

import "errors"

var ErrSomethingWentWrong = errors.New("something went wrong")
var ErrInternalServerError = errors.New("Internal Server Error")
var ErrContextDeadlineExceeded = errors.New("context deadline exceeded")
var ErrClientDisconnected = errors.New("client is disconnected")
var ErrUserNotFound = errors.New("user not found")
