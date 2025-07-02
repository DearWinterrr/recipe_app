// index.js
const defaultAvatarUrl = 'https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0'

Page({
  data: {
    motto: 'Hello World',
    userInfo: {
      avatarUrl: defaultAvatarUrl,
      nickName: '',
    },
    hasUserInfo: false,
    canIUseGetUserProfile: wx.canIUse('getUserProfile'),
    canIUseNicknameComp: wx.canIUse('input.type.nickname'),
    banners: [
      { id: 1, image: 'https://example.com/banner1.jpg', title: '夏季清爽菜' },
      { id: 2, image: 'https://example.com/banner2.jpg', title: '减脂餐' },
      { id: 3, image: 'https://example.com/banner3.jpg', title: '热门菜谱' }
    ],
    ingredients: [],
    selectedIngredients: [],
    recommendedRecipes: [],
    searchValue: '',
    searchResults: []
  },
  bindViewTap() {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },
  onChooseAvatar(e) {
    const { avatarUrl } = e.detail
    const { nickName } = this.data.userInfo
    this.setData({
      "userInfo.avatarUrl": avatarUrl,
      hasUserInfo: nickName && avatarUrl && avatarUrl !== defaultAvatarUrl,
    })
  },
  onInputChange(e) {
    const nickName = e.detail.value
    const { avatarUrl } = this.data.userInfo
    this.setData({
      "userInfo.nickName": nickName,
      hasUserInfo: nickName && avatarUrl && avatarUrl !== defaultAvatarUrl,
    })
  },
  getUserProfile(e) {
    // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认，开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    wx.getUserProfile({
      desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
        console.log(res)
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }
    })
  },
  onLoad() {
    this.loadIngredients();
  },
  // 加载所有食材
  loadIngredients() {
    wx.showLoading({ title: '加载中' });
    wx.request({
      url: 'http://localhost:8080/api/ingredients',
      success: (res) => {
        this.setData({
          ingredients: res.data
        });
      },
      fail: (err) => {
        console.error('加载食材失败', err);
        wx.showToast({
          title: '加载食材失败',
          icon: 'none'
        });
      },
      complete: () => {
        wx.hideLoading();
      }
    });
  },
  // 选择食材
  selectIngredient(e) {
    const id = e.currentTarget.dataset.id;
    const ingredient = this.data.ingredients.find(item => item.ID === id);
    
    if (!ingredient) return;
    
    // 检查是否已选择
    const exists = this.data.selectedIngredients.some(item => item.ID === id);
    
    if (!exists) {
      this.setData({
        selectedIngredients: [...this.data.selectedIngredients, ingredient]
      });
      this.getRecipeRecommendations();
    }
  },
  // 移除已选择的食材
  removeIngredient(e) {
    const id = e.currentTarget.dataset.id;
    this.setData({
      selectedIngredients: this.data.selectedIngredients.filter(item => item.ID !== id)
    });
    this.getRecipeRecommendations();
  },
  // 获取菜谱推荐
  getRecipeRecommendations() {
    if (this.data.selectedIngredients.length === 0) {
      this.setData({ recommendedRecipes: [] });
      return;
    }

    wx.showLoading({ title: '推荐中' });
    wx.request({
      url: 'http://localhost:8080/api/recommend',
      method: 'POST',
      data: {
        ingredients: this.data.selectedIngredients.map(item => item.ID)
      },
      success: (res) => {
        this.setData({
          recommendedRecipes: res.data
        });
      },
      fail: (err) => {
        console.error('获取推荐失败', err);
        wx.showToast({
          title: '获取推荐失败',
          icon: 'none'
        });
      },
      complete: () => {
        wx.hideLoading();
      }
    });
  },
  // 搜索功能
  onSearch(e) {
    const value = e.detail.value;
    this.setData({ searchValue: value });

    if (!value) {
      this.setData({ searchResults: [] });
      return;
    }

    wx.showLoading({ title: '搜索中' });
    wx.request({
      url: `http://localhost:8080/api/search?keyword=${encodeURIComponent(value)}`,
      success: (res) => {
        // 合并搜索结果，优先展示菜谱
        this.setData({
          searchResults: [...(res.data.recipes || []), ...(res.data.ingredients || [])]
        });
      },
      fail: (err) => {
        console.error('搜索失败', err);
      },
      complete: () => {
        wx.hideLoading();
      }
    });
  },
  // 清除搜索
  clearSearch() {
    this.setData({
      searchValue: '',
      searchResults: []
    });
  },
  // 点击食材搜索结果
  onSearchIngredientTap(e) {
    const id = e.currentTarget.dataset.id;
    const ingredient = this.data.ingredients.find(item => item.ID === id);
    
    if (ingredient) {
      const exists = this.data.selectedIngredients.some(item => item.ID === id);
      
      if (!exists) {
        this.setData({
          selectedIngredients: [...this.data.selectedIngredients, ingredient],
          searchValue: '',
          searchResults: []
        });
        this.getRecipeRecommendations();
      }
    }
  },
  // 点击菜谱搜索结果或推荐结果
  onRecipeTap(e) {
    const id = e.currentTarget.dataset.id;
    wx.navigateTo({
      url: `../recipeDetail/recipeDetail?id=${id}`
    });
  },
  // 点击轮播图
  onBannerTap(e) {
    const id = e.currentTarget.dataset.id;
    // 这里可以根据banner ID跳转到不同的专题页面
    // 例如：夏季清爽菜专题、减脂餐专题等
    // 简化处理，这里直接跳转到详情页
    wx.navigateTo({
      url: `../recipeDetail/recipeDetail?id=${id}`
    });
  }
})
