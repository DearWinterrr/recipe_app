package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"recipe-app/backend/database"
)

// RegisterRoutes 注册API路由
func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	// 获取GORM数据库实例
	gormDB := database.DB

	// 创建API处理器
	handler := NewHandler(gormDB)

	// API路由组
	api := r.Group("/api")
	{
		// 食材相关接口
		api.GET("/ingredients", handler.GetIngredients)

		// 搜索接口
		api.GET("/search", handler.Search)

		// 菜谱推荐接口
		api.POST("/recommend", handler.RecommendRecipes)

		// 菜谱详情接口
		api.GET("/recipes/:id", handler.GetRecipeDetail)

		// 轮播图接口
		api.GET("/banners", handler.GetBanners)
	}
}

// Handler API处理器
type Handler struct {
	DB *gorm.DB
}

// NewHandler 创建新的API处理器
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}
