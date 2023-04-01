import Taro from '@tarojs/taro'

export default {
    TokenKey: 'token',
    UserProfiltKey: 'user',

    Set(key, data) {
        console.log('store set key', key, data.length)
        Taro.setStorage({
            key: key,
            data: data
        })
    },
    Get(key) {
        console.log('store get key', key)

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
    }
}
