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

type TransactionController struct {
	TransactionService services.TransactionServiceContract
	UserService        services.UserServiceContract
	WalletService      services.WalletServiceContract
}

func (controller *TransactionController) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	apiKey := helpers.GetApiKey(r.Header)
	user := controller.UserService.FindByApiKey(r.Context(), apiKey)

	transactionResponse := controller.TransactionService.GetAll(r.Context(), user.ID)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *TransactionController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	apiKey := helpers.GetApiKey(r.Header)
	user := controller.UserService.FindByApiKey(r.Context(), apiKey)

	transactionCreateRequest := requests.TransactionCreateRequest{}
	helpers.ReadFromRequestBody(r, &transactionCreateRequest)

	controller.WalletService.FindById(r.Context(), transactionCreateRequest.WalletID, user.ID)

	transactionResponse := controller.TransactionService.Create(r.Context(), transactionCreateRequest)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *TransactionController) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	apiKey := helpers.GetApiKey(r.Header)
	user := controller.UserService.FindByApiKey(r.Context(), apiKey)

	transactionId, err := strconv.Atoi(params.ByName("transactionId"))
	helpers.PanicIfError(err)
	transactionResponse := controller.TransactionService.FindById(r.Context(), int32(transactionId), user.ID)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *TransactionController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	transactionUpdateRequest := requests.TransactionUpdateRequest{}
	helpers.ReadFromRequestBody(r, &transactionUpdateRequest)

	transactionId, err := strconv.Atoi(params.ByName("transactionId"))
	transactionUpdateRequest.ID = int32(transactionId)
	helpers.PanicIfError(err)

	apiKey := helpers.GetApiKey(r.Header)
	user := controller.UserService.FindByApiKey(r.Context(), apiKey)

	transactionUpdateRequest.UserID = user.ID

	transactionResponse := controller.TransactionService.Update(r.Context(), transactionUpdateRequest)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *TransactionController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	transactionId, err := strconv.Atoi(params.ByName("transactionId"))
	helpers.PanicIfError(err)

	apiKey := helpers.GetApiKey(r.Header)
	user := controller.UserService.FindByApiKey(r.Context(), apiKey)

	controller.TransactionService.Delete(r.Context(), int32(transactionId), user.ID)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "",
	}

	helpers.WriteToResponseBody(w, webResponse)
}
