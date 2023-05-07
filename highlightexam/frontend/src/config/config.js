export default {
    // apiHost: "https://highlightexam.lizeyuan.ltd",
    apiHost: "http://127.0.0.1:17000",

    notNeedTokenUrl: [
        '/hl_api/login/wechat',
        '/hl_api/config/config',
        '/hl_api/subject/list',
        '/hl_api/subject/detail'
    ],
    skipToLoginCode: [
        101000,
        101001
    ]
}