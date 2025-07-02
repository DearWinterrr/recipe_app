package utils

import (
	"sort"

	"recipe-app/backend/models"
)

// SortRecipesByMatchRate 根据食材匹配率对菜谱进行排序
func SortRecipesByMatchRate(recipes []models.Recipe, selectedIngredients []int64) []models.Recipe {
	type RecipeMatch struct {
		Recipe    models.Recipe
		MatchRate float64
	}

	var matchedRecipes []RecipeMatch

	// 遍历菜谱，计算匹配度
	for _, recipe := range recipes {
		matchCount := 0
		for _, ri := range recipe.Ingredients {
			for _, id := range selectedIngredients {
				if ri.IngredientID == id {
					matchCount++
					break
				}
			}
		}

		// 计算匹配率
		var matchRate float64
		if len(recipe.Ingredients) > 0 {
			matchRate = float64(matchCount) / float64(len(recipe.Ingredients))
		}

		matchedRecipes = append(matchedRecipes, RecipeMatch{
			Recipe:    recipe,
			MatchRate: matchRate,
		})
	}

	// 根据匹配率排序
	sort.Slice(matchedRecipes, func(i, j int) bool {
		return matchedRecipes[i].MatchRate > matchedRecipes[j].MatchRate
	})

	// 提取排序后的菜谱
	var sortedRecipes []models.Recipe
	for _, rm := range matchedRecipes {
		sortedRecipes = append(sortedRecipes, rm.Recipe)
	}

	return sortedRecipes
}
