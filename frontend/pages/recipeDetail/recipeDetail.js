// recipeDetail.js
Page({
  data: {
    recipeId: null,
    recipe: null,
    loading: true
  },

  onLoad(options) {
    const { id } = options;
    this.setData({ recipeId: id });
    this.fetchRecipeDetail(id);
  },

  // 获取菜谱详情
  fetchRecipeDetail(id) {
    wx.showLoading({ title: '加载中' });
    this.setData({ loading: true });
    
    wx.request({
      url: `http://8.130.20.40:8080/api/recipes/${id}`,
      success: (res) => {
        this.setData({
          recipe: res.data,
          loading: false
        });
      },
      fail: (err) => {
        console.error('获取菜谱详情失败', err);
        wx.showToast({
          title: '获取菜谱详情失败',
          icon: 'none'
        });
        this.setData({ loading: false });
      },
      complete: () => {
        wx.hideLoading();
      }
    });
  },

  // 点击视频链接
  onVideoTap() {
    if (this.data.recipe && this.data.recipe.video_url) {
      wx.navigateTo({
        url: `/pages/webview/webview?url=${encodeURIComponent(this.data.recipe.video_url)}`
      });
    }
  },

  // 点击网页链接
  onWebpageTap() {
    if (this.data.recipe && this.data.recipe.webpage_url) {
      wx.navigateTo({
        url: `/pages/webview/webview?url=${encodeURIComponent(this.data.recipe.webpage_url)}`
      });
    }
  }
}) 