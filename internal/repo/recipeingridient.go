package repo

import (
	"github.com/koray.oezdemir/go-back-end/internal/model"
	"gorm.io/gorm"
)

type RecipeIngredientRepository interface {
	Create(ri *model.RecipeIngredient) error
}

type recipeIngredientRepository struct {
	db *gorm.DB
}

func NewRecipeIngredientRepository(db *gorm.DB) RecipeIngredientRepository {
	return &recipeIngredientRepository{db: db}
}

func (r *recipeIngredientRepository) Create(ri *model.RecipeIngredient) error {
	return r.db.Create(ri).Error
}
