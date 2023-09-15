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

type WalletController struct {
	WalletService services.WalletServiceContract
	UserService   services.UserServiceContract
}

func (controller *WalletController) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	apiKey := helpers.GetApiKey(r.Header)
	user := controller.UserService.FindByApiKey(r.Context(), apiKey)

	walletResponse := controller.WalletService.GetAll(r.Context(), user.ID)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   walletResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *WalletController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	walletCreateRequest := requests.WalletCreateRequest{}
	helpers.ReadFromRequestBody(r, &walletCreateRequest)

	walletResponse := controller.WalletService.Create(r.Context(), walletCreateRequest)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   walletResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *WalletController) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	apiKey := helpers.GetApiKey(r.Header)
	user := controller.UserService.FindByApiKey(r.Context(), apiKey)

	walletId, err := strconv.Atoi(params.ByName("walletId"))
	helpers.PanicIfError(err)
	walletResponse := controller.WalletService.FindById(r.Context(), int32(walletId), user.ID)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   walletResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *WalletController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	walletUpdateRequest := requests.WalletUpdateRequest{}
	helpers.ReadFromRequestBody(r, &walletUpdateRequest)

	walletId, err := strconv.Atoi(params.ByName("walletId"))
	walletUpdateRequest.ID = int32(walletId)
	helpers.PanicIfError(err)

	walletResponse := controller.WalletService.Update(r.Context(), walletUpdateRequest)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   walletResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *WalletController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	walletId, err := strconv.Atoi(params.ByName("walletId"))
	helpers.PanicIfError(err)

	controller.WalletService.Delete(r.Context(), int32(walletId))

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "",
	}

	helpers.WriteToResponseBody(w, webResponse)
}
