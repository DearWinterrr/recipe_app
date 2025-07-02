package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"recipe-app/backend/models"
	"recipe-app/backend/utils"
)

// GetIngredients 获取所有食材
func (h *Handler) GetIngredients(c *gin.Context) {
	var ingredients []models.Ingredient

	if err := h.DB.Find(&ingredients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取食材失败"})
		return
	}

	c.JSON(http.StatusOK, ingredients)
}

// Search 搜索食材和菜谱
func (h *Handler) Search(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "搜索关键词不能为空"})
		return
	}

	// 搜索结果
	result := models.SearchResult{
		Recipes:     []models.Recipe{},
		Ingredients: []models.Ingredient{},
	}

	// 搜索食材
	if err := h.DB.Where("name LIKE ?", "%"+keyword+"%").Find(&result.Ingredients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "搜索食材失败"})
		return
	}

	// 搜索菜谱
	if err := h.DB.Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&result.Recipes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "搜索菜谱失败"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// RecommendRecipes 推荐菜谱
func (h *Handler) RecommendRecipes(c *gin.Context) {
	var req models.RecipeRecommendRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	if len(req.Ingredients) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请选择至少一种食材"})
		return
	}

	// 查询包含所选食材的菜谱
	var recipes []models.Recipe

	// 使用子查询获取包含所选食材的菜谱
	subQuery := h.DB.Table("recipe_ingredients").
		Select("recipe_id").
		Where("ingredient_id IN ?", req.Ingredients).
		Group("recipe_id")

	// 获取菜谱及其食材
	h.DB.Table("recipes").
		Joins("JOIN (?) AS matched_recipes ON recipes.id = matched_recipes.recipe_id", subQuery).
		Preload("Ingredients.Ingredient").
		Find(&recipes)

	// 使用工具函数对菜谱进行排序
	sortedRecipes := utils.SortRecipesByMatchRate(recipes, req.Ingredients)

	c.JSON(http.StatusOK, sortedRecipes)
}

// GetRecipeDetail 获取菜谱详情
func (h *Handler) GetRecipeDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的菜谱ID"})
		return
	}

	var recipe models.Recipe

	// 加载菜谱及其关联数据
	if err := h.DB.Preload("Ingredients.Ingredient").
		Preload("Steps").
		Preload("Nutrition").
		First(&recipe, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "菜谱不存在"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

// GetBanners 获取轮播图
func (h *Handler) GetBanners(c *gin.Context) {
	var banners []models.Banner

	if err := h.DB.Find(&banners).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取轮播图失败"})
		return
	}

	c.JSON(http.StatusOK, banners)
}
