package exceptions

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if validationErrors(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("X-Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := responses.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helpers.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("X-Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := responses.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helpers.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("X-Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := responses.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helpers.WriteToResponseBody(w, webResponse)
}
