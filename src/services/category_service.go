package services

import (
	"context"

	"github.com/aminazadbakht1/golang-clean-web-api/api/dto"
	"github.com/aminazadbakht1/golang-clean-web-api/config"
	"github.com/aminazadbakht1/golang-clean-web-api/data/db"
	"github.com/aminazadbakht1/golang-clean-web-api/data/models"
	"github.com/aminazadbakht1/golang-clean-web-api/pkg/logging"
)

type CategoryService struct {
	base *BaseService[models.Category, dto.CreateCategoryRequest, dto.UpdateCategoryRequest, dto.CategoryResponse]
}

func NewCategoryService(cfg *config.Config) *CategoryService {
	return &CategoryService{
		base: &BaseService[models.Category, dto.CreateCategoryRequest, dto.UpdateCategoryRequest, dto.CategoryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *CategoryService) Create(ctx context.Context, req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CategoryService) Update(ctx context.Context, id int, req *dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CategoryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CategoryService) GetById(ctx context.Context, id int) (*dto.CategoryResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CategoryService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CategoryResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
