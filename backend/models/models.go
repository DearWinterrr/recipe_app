package models

// Ingredient 食材模型
type Ingredient struct {
	ID          int64  `json:"ID" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

// Recipe 菜谱模型
type Recipe struct {
	ID          int64              `json:"ID" gorm:"primaryKey"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	ImageURL    string             `json:"image_url"`
	VideoURL    string             `json:"video_url"`
	WebpageURL  string             `json:"webpage_url"`
	CookTime    int                `json:"cook_time"` // 烹饪时间（分钟）
	Ingredients []RecipeIngredient `json:"ingredients,omitempty"`
	Steps       []RecipeStep       `json:"steps,omitempty"`
	Nutrition   RecipeNutrition    `json:"nutrition,omitempty"`
}

// RecipeIngredient 菜谱食材关联模型
type RecipeIngredient struct {
	ID           int64      `json:"ID" gorm:"primaryKey"`
	RecipeID     int64      `json:"recipe_id"`
	IngredientID int64      `json:"ingredient_id"`
	Ingredient   Ingredient `json:"ingredient"`
	Quantity     float64    `json:"quantity"` // 用量（克）
	Unit         string     `json:"unit"`     // 单位
}

// RecipeStep 菜谱制作步骤模型
type RecipeStep struct {
	ID          int64  `json:"ID" gorm:"primaryKey"`
	RecipeID    int64  `json:"recipe_id"`
	StepNumber  int    `json:"step_number"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

// RecipeNutrition 菜谱营养含量模型
type RecipeNutrition struct {
	ID            int64   `json:"ID" gorm:"primaryKey"`
	RecipeID      int64   `json:"recipe_id"`
	Calories      float64 `json:"calories"`      // 热量（千卡）
	Protein       float64 `json:"protein"`       // 蛋白质（克）
	Fat           float64 `json:"fat"`           // 脂肪（克）
	Carbohydrates float64 `json:"carbohydrates"` // 碳水化合物（克）
}

// Banner 轮播图模型
type Banner struct {
	ID       int64  `json:"ID" gorm:"primaryKey"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	RecipeID int64  `json:"recipe_id"`
}

// SearchResult 搜索结果
type SearchResult struct {
	Recipes     []Recipe     `json:"recipes"`
	Ingredients []Ingredient `json:"ingredients"`
}

// RecipeRecommendRequest 菜谱推荐请求
type RecipeRecommendRequest struct {
	Ingredients []int64 `json:"ingredients"`
}
