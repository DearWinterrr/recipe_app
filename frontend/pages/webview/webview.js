// webview.js
Page({
  data: {
    url: '',
    loading: true,
    error: false
  },
  
  onLoad(options) {
    if (options.url) {
      this.setData({
        url: decodeURIComponent(options.url)
      });
    } else {
      this.setData({
        error: true
      });
    }
  },

  // 网页加载完成
  onWebviewLoad() {
    this.setData({
      loading: false
    });
  },

  // 网页加载失败
  onWebviewError(e) {
    console.error('Webview error:', e.detail);
    this.setData({
      loading: false,
      error: true
    });
  },

  // 返回上一页
  goBack() {
    wx.navigateBack();
  }
}) 