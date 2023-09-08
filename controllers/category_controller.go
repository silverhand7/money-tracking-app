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

type CategoryController struct {
	CategoryService services.CategoryServiceContract
}

func (controller *CategoryController) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponse := controller.CategoryService.GetAll(r.Context())

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := requests.CategoryCreateRequest{}
	helpers.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryController) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helpers.PanicIfError(err)
	categoryResponse := controller.CategoryService.FindById(r.Context(), int32(categoryId))

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateRequest := requests.CategoryUpdateRequest{}
	helpers.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	categoryUpdateRequest.ID = int32(categoryId)
	helpers.PanicIfError(err)

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helpers.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), int32(categoryId))

	webResponse := responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "",
	}

	helpers.WriteToResponseBody(w, webResponse)
}
