package controllers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/silverhand7/money-tracking-app/helpers"
	"github.com/silverhand7/money-tracking-app/models/web/requests"
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

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := requests.UserCreateRequest{}
	helpers.ReadFromRequestBody(r, &userCreateRequest)

	userResponse := controller.UserService.Create(r.Context(), userCreateRequest)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UserController) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId, err := strconv.Atoi(params.ByName("userId"))
	helpers.PanicIfError(err)
	userResponses := controller.UserService.FindById(r.Context(), int32(userId))

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helpers.WriteToResponseBody(w, webResponse)
}
