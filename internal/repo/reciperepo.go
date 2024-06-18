package repo

import (
	"github.com/koray.oezdemir/go-back-end/internal/model"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	Create(recipe *model.Recipe) error
	Read(id uint) (*model.Recipe, error)
	Update(recipe *model.Recipe) error
	Delete(id uint) error
}

type recipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) RecipeRepository {
	return &recipeRepository{db: db}
}

func (r *recipeRepository) Create(recipe *model.Recipe) error {
	return r.db.Create(recipe).Error
}

func (r *recipeRepository) Read(id uint) (*model.Recipe, error) {
	var recipe model.Recipe
	err := r.db.First(&recipe, id).Error
	return &recipe, err
}

func (r *recipeRepository) Update(recipe *model.Recipe) error {
	return r.db.Save(recipe).Error
}

func (r *recipeRepository) Delete(id uint) error {
	return r.db.Delete(&model.Recipe{}, id).Error
}
