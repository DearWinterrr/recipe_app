# 菜谱推荐小程序后端

## 技术栈

- Go 语言
- Gin Web 框架
- GORM ORM 框架
- MySQL 数据库

## 项目结构

```
backend/
├── api/            # API 处理器
├── config/         # 配置文件
├── database/       # 数据库连接
├── models/         # 数据模型
├── utils/          # 工具函数
├── go.mod          # Go 模块文件
├── go.sum          # Go 依赖版本锁定文件
├── config.json     # 应用程序配置文件
└── main.go         # 主程序入口
```

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

## 如何运行

1. 确保已安装 Go 环境（推荐 Go 1.19 或更高版本）
2. 安装 MySQL 数据库并创建名为 `recipe_app` 的数据库
3. 修改 `config.json` 文件中的数据库配置
4. 运行以下命令安装依赖并启动服务：

```bash
cd backend
go mod tidy
go run main.go
```

服务将在配置的端口上启动（默认为 8080）

## 数据库初始化

首次运行时，系统会自动创建所需的数据表。如需添加初始数据，可以使用 MySQL 客户端执行 SQL 脚本。 


如何运行项目
后端：
安装Go和MySQL
配置数据库连接（修改config.json）
运行setup_db.sh初始化数据库
执行go mod tidy安装依赖
运行go run main.go启动服务器
前端：
使用微信开发者工具打开frontend目录
确保后端服务已启动
点击编译运行