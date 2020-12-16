package errors

import "net/http"

func (e *appError) BadRequest() AppError {
	e.status = http.StatusBadRequest
	return e
}
func (e *appError) Unauthorized() AppError {
	e.status = http.StatusUnauthorized
	return e
}
func (e *appError) NotFound() AppError {
	e.status = http.StatusNotFound
	return e
}
func (e *appError) InternalServerError() AppError {
	e.status = http.StatusInternalServerError
	return e
}
