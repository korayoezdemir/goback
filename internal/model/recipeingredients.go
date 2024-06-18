package model

type RecipeIngredient struct {
	ID           uint `gorm:"primaryKey"`         
	RecipeID     uint                             
	IngredientID uint                             
	Quantity     string
	Recipe       Recipe      `gorm:"foreignKey:RecipeID"`
	Ingredient   Ingredient  `gorm:"foreignKey:IngredientID"`
}
