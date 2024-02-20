package dto

import (
	"net/http"
)

type HttpError struct {
	Status   int                `json:"status"`
	Type     string             `json:"type"`
	Language string             `json:"language"`
	Errors   []errorWithMessage `json:"errors"`
}

type errorWithMessage struct {
	error
	Message string `json:"message"`
}

func errorsWithMessage(errs ...error) []errorWithMessage {
	var errorArr []errorWithMessage
	for _, err := range errs {
		if err == nil {
			continue
		}

		errorArr = append(errorArr, errorWithMessage{
			error:   err,
			Message: err.Error(),
		})
	}

	return errorArr
}

func HttpErrorFrom(status int, errs ...error) HttpError {
	return HttpError{
		Status:   status,
		Type:     http.StatusText(status),
		Language: "ko",
		Errors:   errorsWithMessage(errs...),
	}
}

func HttpUnauthorizedError(errs ...error) HttpError {
	return HttpErrorFrom(http.StatusUnauthorized, errs...)
}

func HttpUnprocessableEntityError(errs ...error) HttpError {
	return HttpErrorFrom(http.StatusUnprocessableEntity, errs...)
}

func HttpBadRequestError(errs ...error) HttpError {
	return HttpErrorFrom(http.StatusBadRequest, errs...)
}

func HttpNotFoundError(errs ...error) HttpError {
	return HttpErrorFrom(http.StatusNotFound, errs...)
}
