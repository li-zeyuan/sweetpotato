export default defineAppConfig({
  pages: [
    'pages/index/index',
    'pages/mine/index',
    'pages/login/index',
    'pages/subject/index',
    'pages/studyrecord/index',
    'pages/knowledge/index'
  ],
  window: {
    backgroundTextStyle: 'light',
    navigationBarBackgroundColor: '#fff',
    navigationBarTitleText: 'WeChat',
    navigationBarTextStyle: 'black'
  },
  tabBar: {
    list: [
      {
        pagePath: 'pages/index/index',
        text: '复习',
        iconPath: './assets/images/tab_trend.png',
        selectedIconPath: './assets/images/tab_trend_s.png'
      },
      {
        pagePath: 'pages/subject/index',
        text: '题库',
        iconPath: './assets/images/tab_news.png',
        selectedIconPath: './assets/images/tab_news_s.png'
      },
      {
        pagePath: 'pages/mine/index',
        text: '我的',
        iconPath: './assets/images/tab_me.png',
        selectedIconPath: './assets/images/tab_me_s.png'
      }
    ],
    color: '#8a8a8a',
    selectedColor: '#2d8cf0',
    backgroundColor: '#ffffff',
    borderStyle: 'white',
  }
})
