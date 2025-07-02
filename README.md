# 菜谱推荐小程序

一个基于微信小程序和Go后端的菜谱推荐应用，用户可以选择食材，系统会自动推荐适合的菜谱。

## 项目结构

```
GoProject/
├── backend/           # Go后端
│   ├── api/           # API处理器
│   ├── config/        # 配置文件
│   ├── database/      # 数据库连接
│   ├── models/        # 数据模型
│   ├── utils/         # 工具函数
│   ├── config.json    # 配置文件
│   ├── init_data.sql  # 初始数据
│   ├── setup_db.sh    # 数据库设置脚本
│   └── main.go        # 主程序
├── frontend/          # 微信小程序前端
│   ├── pages/         # 页面
│   │   ├── index/     # 首页（食材选择和菜谱推荐）
│   │   ├── recipeDetail/ # 菜谱详情页
│   │   ├── webview/   # 外部网页查看
│   │   └── logs/      # 日志页面
│   ├── utils/         # 工具函数
│   ├── app.js         # 小程序入口
│   ├── app.json       # 小程序配置
│   └── app.wxss       # 全局样式
└── README.md          # 项目说明
```

## 技术栈

### 前端
- 微信小程序原生开发
- WXML、WXSS、JavaScript

### 后端
- Go语言
- Gin Web框架
- GORM ORM框架
- MySQL数据库

## 功能特点

- 食材选择：用户可以选择多个食材，系统会根据食材匹配度推荐菜谱
- 搜索功能：支持搜索食材和菜谱
- 轮播图推荐：展示热门菜谱、时令推荐或专题菜谱
- 菜谱详情：展示菜谱的详细信息，包括食材、制作步骤、营养成分和推荐视频/网页

## 数据库设计

### 表结构

1. **ingredients（食材表）**
   - ID（主键，自增）
   - name（食材名称，唯一）
   - description（食材描述）
   - image_url（食材图片URL）

2. **recipes（菜谱表）**
   - ID（主键，自增）
   - name（菜谱名称）
   - description（菜谱简介）
   - image_url（菜谱图片URL）
   - video_url（视频URL）
   - webpage_url（网页URL）
   - cook_time（烹饪时间，分钟）

3. **recipe_ingredients（菜谱食材关联表）**
   - ID（主键，自增）
   - recipe_id（外键，关联菜谱表）
   - ingredient_id（外键，关联食材表）
   - quantity（食材用量）
   - unit（单位）

4. **recipe_steps（菜谱步骤表）**
   - ID（主键，自增）
   - recipe_id（外键，关联菜谱表）
   - step_number（步骤编号）
   - description（步骤描述）
   - image_url（步骤图片URL）

5. **recipe_nutrition（菜谱营养成分表）**
   - ID（主键，自增）
   - recipe_id（外键，关联菜谱表）
   - calories（热量）
   - protein（蛋白质）
   - fat（脂肪）
   - carbohydrates（碳水化合物）

6. **banners（轮播图表）**
   - ID（主键，自增）
   - title（标题）
   - image_url（图片URL）
   - recipe_id（关联的菜谱ID）

## 如何运行

### 后端设置

1. 确保已安装 Go 环境（推荐 Go 1.19 或更高版本）
2. 安装 MySQL 数据库
3. 修改 `backend/config.json` 文件中的数据库配置
4. 运行数据库设置脚本：

```bash
cd backend
chmod +x setup_db.sh
./setup_db.sh
```

5. 安装依赖并启动后端服务：

```bash
cd backend
go mod tidy
go run main.go
```

### 前端设置

1. 使用微信开发者工具打开 `frontend` 目录
2. 确保后端服务已启动
3. 在微信开发者工具中点击"编译"运行小程序

## API 接口

### 1. 获取食材列表

- **URL**: `/api/ingredients`
- **方法**: `GET`
- **响应**: 食材列表

### 2. 搜索食材和菜谱

- **URL**: `/api/search?keyword=关键词`
- **方法**: `GET`
- **参数**: `keyword`（搜索关键词）
- **响应**: 匹配的食材和菜谱列表

### 3. 推荐菜谱

- **URL**: `/api/recommend`
- **方法**: `POST`
- **请求体**:
  ```json
  {
    "ingredients": [1, 2, 3]  // 食材ID列表
  }
  ```
- **响应**: 推荐的菜谱列表，按匹配度排序

### 4. 获取菜谱详情

- **URL**: `/api/recipes/:id`
- **方法**: `GET`
- **参数**: `id`（菜谱ID）
- **响应**: 菜谱详细信息，包括食材、步骤、营养成分等

### 5. 获取轮播图

- **URL**: `/api/banners`
- **方法**: `GET`
- **响应**: 轮播图列表 