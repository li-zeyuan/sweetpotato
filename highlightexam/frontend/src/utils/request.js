import Taro from '@tarojs/taro'
import store from "./store"
import config from "../config/config"

const apiHost = config.apiHost

export default function request(method, url, data) {
    const token = store.Get(store.TokenKey)
    if (!config.notNeedTokenUrl.includes(url) && !token) {
        Taro.navigateTo({
            url: '/pages/login/index',
        })
    }

    return new Promise((resolve, reject) => {
        let apiUrl = apiHost
        switch (method) {
            case 'get': {
                apiUrl += url + '?' + Object.keys(data).map(k => {
                    const v = data[k]
                    return `${k}=${v}`
                }).join('&')
                break
            }
            case 'post': {
                apiUrl += url
                break
            }
            case 'put': {
                apiUrl += url
                break
            }
        }

        Taro.request({
            method: method,
            url: apiUrl,
            data: (method == 'post') || (method == 'put')  ? data : {},
            header: { 'Authorization': token },
            success: function (res) {
                if (res.statusCode == 403) {
                    store.Del(store.TokenKey)
                    
                    return
                }

                if (config.skipToLoginCode.includes(res.code)) {
                    Taro.reLaunch({
                        url: '/pages/login/index',
                    })
                }

                resolve(res)
            },
            fail: function (res) {
                reject(res)
            }
        })
    })
}