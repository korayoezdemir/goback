package repo

import (
	"github.com/koray.oezdemir/go-back-end/internal/model"
	"gorm.io/gorm"
)

type IngredientRepository interface {
	Create(ingredient *model.Ingredient) error
	Read(id uint) (*model.Ingredient, error)
	Update(ingredient *model.Ingredient) error
	Delete(id uint) error
}

type ingredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) IngredientRepository {
	return &ingredientRepository{db: db}
}

func (r *ingredientRepository) Create(ingredient *model.Ingredient) error {
	return r.db.Create(ingredient).Error
}

func (r *ingredientRepository) Read(id uint) (*model.Ingredient, error) {
	var ingredient model.Ingredient
	err := r.db.First(&ingredient, id).Error
	return &ingredient, err
}

func (r *ingredientRepository) Update(ingredient *model.Ingredient) error {
	return r.db.Save(ingredient).Error
}

func (r *ingredientRepository) Delete(id uint) error {
	return r.db.Delete(&model.Ingredient{}, id).Error
}
