-- 确保使用正确的数据库
USE recipe_app;

-- 清空现有表数据（如果存在）
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS banners;
DROP TABLE IF EXISTS recipe_nutrition;
DROP TABLE IF EXISTS recipe_steps;
DROP TABLE IF EXISTS recipe_ingredients;
DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS ingredients;
SET FOREIGN_KEY_CHECKS = 1;

-- 创建表
CREATE TABLE ingredients (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) UNIQUE,
  description TEXT,
  image_url VARCHAR(255)
);

CREATE TABLE recipes (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  description TEXT,
  image_url VARCHAR(255),
  video_url VARCHAR(255),
  webpage_url VARCHAR(255),
  cook_time INT
);

CREATE TABLE recipe_ingredients (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  recipe_id BIGINT,
  ingredient_id BIGINT,
  quantity FLOAT,
  unit VARCHAR(50),
  FOREIGN KEY (recipe_id) REFERENCES recipes(id),
  FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
);

CREATE TABLE recipe_steps (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  recipe_id BIGINT,
  step_number INT,
  description TEXT,
  image_url VARCHAR(255),
  FOREIGN KEY (recipe_id) REFERENCES recipes(id)
);

CREATE TABLE recipe_nutrition (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  recipe_id BIGINT,
  calories FLOAT,
  protein FLOAT,
  fat FLOAT,
  carbohydrates FLOAT,
  FOREIGN KEY (recipe_id) REFERENCES recipes(id)
);

CREATE TABLE banners (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255),
  image_url VARCHAR(255),
  recipe_id BIGINT,
  FOREIGN KEY (recipe_id) REFERENCES recipes(id)
);

-- 插入食材数据
INSERT INTO ingredients (name, description, image_url) VALUES
('胡萝卜', '富含胡萝卜素，对眼睛有益', 'https://example.com/carrot.jpg'),
('牛肉', '优质蛋白质来源，富含铁质', 'https://example.com/beef.jpg'),
('鸡蛋', '营养丰富，含有优质蛋白质', 'https://example.com/egg.jpg'),
('西红柿', '富含维生素C和抗氧化物质', 'https://example.com/tomato.jpg'),
('土豆', '碳水化合物的良好来源', 'https://example.com/potato.jpg'),
('洋葱', '含有抗氧化物质，增添菜肴风味', 'https://example.com/onion.jpg'),
('大蒜', '具有抗菌作用，增添菜肴风味', 'https://example.com/garlic.jpg'),
('青椒', '富含维生素C，增添菜肴色彩', 'https://example.com/pepper.jpg'),
('猪肉', '常见肉类，富含蛋白质', 'https://example.com/pork.jpg'),
('鸡肉', '低脂肪高蛋白质的肉类', 'https://example.com/chicken.jpg');

-- 插入菜谱数据
INSERT INTO recipes (name, description, image_url, video_url, webpage_url, cook_time) VALUES
('番茄炒蛋', '家常美味，简单易做', 'https://example.com/tomato-egg.jpg', 'https://example.com/tomato-egg-video.mp4', 'https://example.com/tomato-egg-recipe.html', 15),
('土豆炖牛肉', '营养丰富，味道浓郁', 'https://example.com/beef-potato.jpg', 'https://example.com/beef-potato-video.mp4', 'https://example.com/beef-potato-recipe.html', 60),
('胡萝卜炒鸡蛋', '色彩鲜艳，营养均衡', 'https://example.com/carrot-egg.jpg', 'https://example.com/carrot-egg-video.mp4', 'https://example.com/carrot-egg-recipe.html', 15),
('青椒土豆丝', '爽口开胃，简单易做', 'https://example.com/pepper-potato.jpg', 'https://example.com/pepper-potato-video.mp4', 'https://example.com/pepper-potato-recipe.html', 20),
('洋葱炒牛肉', '经典搭配，鲜香可口', 'https://example.com/beef-onion.jpg', 'https://example.com/beef-onion-video.mp4', 'https://example.com/beef-onion-recipe.html', 25);

-- 插入菜谱食材关联数据
-- 番茄炒蛋
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity, unit) VALUES
(1, 3, 2, '个'),  -- 鸡蛋 2个
(1, 4, 1, '个');  -- 西红柿 1个

-- 土豆炖牛肉
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity, unit) VALUES
(2, 2, 300, '克'),  -- 牛肉 300克
(2, 5, 2, '个'),    -- 土豆 2个
(2, 6, 1, '个');    -- 洋葱 1个

-- 胡萝卜炒鸡蛋
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity, unit) VALUES
(3, 1, 1, '根'),  -- 胡萝卜 1根
(3, 3, 2, '个');  -- 鸡蛋 2个

-- 青椒土豆丝
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity, unit) VALUES
(4, 5, 1, '个'),  -- 土豆 1个
(4, 8, 1, '个');  -- 青椒 1个

-- 洋葱炒牛肉
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity, unit) VALUES
(5, 2, 200, '克'),  -- 牛肉 200克
(5, 6, 1, '个');    -- 洋葱 1个

-- 插入菜谱步骤数据
-- 番茄炒蛋
INSERT INTO recipe_steps (recipe_id, step_number, description, image_url) VALUES
(1, 1, '将鸡蛋打散，加入少许盐搅拌均匀', 'https://example.com/tomato-egg-step1.jpg'),
(1, 2, '西红柿切块', 'https://example.com/tomato-egg-step2.jpg'),
(1, 3, '热锅倒油，倒入鸡蛋液炒至凝固，盛出备用', 'https://example.com/tomato-egg-step3.jpg'),
(1, 4, '锅中加油，放入西红柿块翻炒至出汁', 'https://example.com/tomato-egg-step4.jpg'),
(1, 5, '倒入炒好的鸡蛋，加盐调味，翻炒均匀即可', 'https://example.com/tomato-egg-step5.jpg');

-- 土豆炖牛肉
INSERT INTO recipe_steps (recipe_id, step_number, description, image_url) VALUES
(2, 1, '牛肉切块，焯水去血水', 'https://example.com/beef-potato-step1.jpg'),
(2, 2, '土豆、洋葱切块', 'https://example.com/beef-potato-step2.jpg'),
(2, 3, '热锅倒油，放入牛肉块煸炒至变色', 'https://example.com/beef-potato-step3.jpg'),
(2, 4, '加入适量水，大火烧开后转小火炖30分钟', 'https://example.com/beef-potato-step4.jpg'),
(2, 5, '加入土豆和洋葱块，继续炖20分钟', 'https://example.com/beef-potato-step5.jpg'),
(2, 6, '加盐、酱油等调味，收汁即可', 'https://example.com/beef-potato-step6.jpg');

-- 胡萝卜炒鸡蛋
INSERT INTO recipe_steps (recipe_id, step_number, description, image_url) VALUES
(3, 1, '胡萝卜切丝', 'https://example.com/carrot-egg-step1.jpg'),
(3, 2, '鸡蛋打散，加入少许盐搅拌均匀', 'https://example.com/carrot-egg-step2.jpg'),
(3, 3, '热锅倒油，倒入鸡蛋液炒至半凝固', 'https://example.com/carrot-egg-step3.jpg'),
(3, 4, '加入胡萝卜丝，翻炒均匀', 'https://example.com/carrot-egg-step4.jpg'),
(3, 5, '加盐调味，翻炒均匀即可', 'https://example.com/carrot-egg-step5.jpg');

-- 青椒土豆丝
INSERT INTO recipe_steps (recipe_id, step_number, description, image_url) VALUES
(4, 1, '土豆切丝，用清水浸泡去除淀粉', 'https://example.com/pepper-potato-step1.jpg'),
(4, 2, '青椒切丝', 'https://example.com/pepper-potato-step2.jpg'),
(4, 3, '热锅倒油，放入土豆丝翻炒至变软', 'https://example.com/pepper-potato-step3.jpg'),
(4, 4, '加入青椒丝，翻炒均匀', 'https://example.com/pepper-potato-step4.jpg'),
(4, 5, '加盐、醋等调味，翻炒均匀即可', 'https://example.com/pepper-potato-step5.jpg');

-- 洋葱炒牛肉
INSERT INTO recipe_steps (recipe_id, step_number, description, image_url) VALUES
(5, 1, '牛肉切片，用料酒、生抽腌制10分钟', 'https://example.com/beef-onion-step1.jpg'),
(5, 2, '洋葱切片', 'https://example.com/beef-onion-step2.jpg'),
(5, 3, '热锅倒油，放入牛肉片煸炒至变色', 'https://example.com/beef-onion-step3.jpg'),
(5, 4, '加入洋葱片，翻炒均匀', 'https://example.com/beef-onion-step4.jpg'),
(5, 5, '加盐、黑胡椒等调味，翻炒均匀即可', 'https://example.com/beef-onion-step5.jpg');

-- 插入菜谱营养成分数据
INSERT INTO recipe_nutrition (recipe_id, calories, protein, fat, carbohydrates) VALUES
(1, 180, 12, 10, 8),   -- 番茄炒蛋
(2, 450, 30, 20, 35),  -- 土豆炖牛肉
(3, 160, 10, 8, 12),   -- 胡萝卜炒鸡蛋
(4, 120, 3, 5, 20),    -- 青椒土豆丝
(5, 280, 25, 15, 10);  -- 洋葱炒牛肉

-- 插入轮播图数据
INSERT INTO banners (title, image_url, recipe_id) VALUES
('夏季清爽菜', 'https://example.com/banner1.jpg', 1),  -- 番茄炒蛋
('减脂餐', 'https://example.com/banner2.jpg', 3),      -- 胡萝卜炒鸡蛋
('热门菜谱', 'https://example.com/banner3.jpg', 2);    -- 土豆炖牛肉 