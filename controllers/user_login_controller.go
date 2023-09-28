package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/web/requests"
	"github.com/silverhand7/money-tracking-app/models/web/responses"
	"github.com/silverhand7/money-tracking-app/services"
)

type UserLoginController struct {
	AuthService services.AuthServiceContract
}

func (controller *UserLoginController) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userLoginRequest := requests.UserLoginRequest{}
	helpers.ReadFromRequestBody(r, &userLoginRequest)

	userResponses := controller.AuthService.Login(r.Context(), userLoginRequest)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helpers.WriteToResponseBody(w, webResponse)
}
