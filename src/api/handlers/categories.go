package handlers

import (
	_ "github.com/aminazadbakht1/golang-clean-web-api/api/dto"
	_ "github.com/aminazadbakht1/golang-clean-web-api/api/helper"
	"github.com/aminazadbakht1/golang-clean-web-api/config"
	"github.com/aminazadbakht1/golang-clean-web-api/services"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(cfg *config.Config) *CategoryHandler {
	return &CategoryHandler{
		service: services.NewCategoryService(cfg),
	}
}

// CreateCategory godoc
// @Summary Create a Category
// @Description Create a Category
// @Tags Categories
// @Accept json
// @produces json
// @Param Request body dto.CreateCategoryRequest true "Create a Category"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CategoryResponse} "Category response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/categories/ [post]
// @Security AuthBearer
func (h *CategoryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// UpdateCategory godoc
// @Summary Update a Category
// @Description Update a Category
// @Tags Categories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCategoryRequest true "Update a Category"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CategoryResponse} "Category response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/categories/{id} [put]
// @Security AuthBearer
func (h *CategoryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// DeleteCategory godoc
// @Summary Delete a Category
// @Description Delete a Category
// @Tags Categories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/categories/{id} [delete]
// @Security AuthBearer
func (h *CategoryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetCategory godoc
// @Summary Get a Category
// @Description Get a Category
// @Tags Categories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CategoryResponse} "Category response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/categories/{id} [get]
// @Security AuthBearer
func (h *CategoryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCategories godoc
// @Summary Get Categories
// @Description Get Categories
// @Tags Categories
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// // @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CategoryResponse]} "Category response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/categories/get-by-filter [post]
// @Security AuthBearer
func (h *CategoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
