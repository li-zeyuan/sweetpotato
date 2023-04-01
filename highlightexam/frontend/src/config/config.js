export default {
    apiHost: "http://127.0.0.1:17000",

    notNeedTokenUrl: [
        '/api/login/wechat',
        '/api/config/config',
        '/api/subject/list',
        '/api/subject/detail'
    ],
    skipToLoginCode: [
        101000,
        101001
    ]
}