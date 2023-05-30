import Taro from '@tarojs/taro'

export default {
    TokenKey: 'token',
    UserProfiltKey: 'user',

    Set(key, data) {
        Taro.setStorage({
            key: key,
            data: data
        })
    },
    Get(key) {
        let data = ''
        try {
            var value = Taro.getStorageSync(key)
            if (value) {
                data = value
            }
        } catch (e) {
            console.log('store get error: ', e)
        }

        return data
    },
    Del(key) {
        Taro.removeStorage({
            key: key,
            success: function (res) {
                console.log(res)
            }
        })
    }
}
