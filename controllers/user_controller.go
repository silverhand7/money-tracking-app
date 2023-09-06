package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
	"github.com/silverhand7/money-tracking-app/services"
)

type UserController struct {
	UserService services.UserServiceContract
}

func (controller *UserController) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userResponses := controller.UserService.GetAll(r.Context())

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helpers.WriteToResponseBody(w, webResponse)
}
