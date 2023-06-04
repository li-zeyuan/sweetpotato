export default {
    // apiHost: "https://highlightexam.lizeyuan.ltd",
    apiHost: "http://39.108.101.229:7001",
    // apiHost: "http://127.0.0.1:7001",

    notNeedTokenUrl: [
        '/hl_api/login/wechat',
        '/hl_api/config/config',
        '/hl_api/subject/list',
        '/hl_api/subject/detail'
    ],
    skipToLoginCode: [
        101000,
        101001 // token 过期
    ]
}
