import Taro from '@tarojs/taro'
import { Component } from 'react'
import { AtButton, AtIcon } from 'taro-ui'
import { wxLogin } from '../../api/api'
import store from '../../utils/store'
import { View, Text } from '@tarojs/components'
import './index.less'

export default class Index extends Component {
    componentWillMount() { }

    componentDidMount() { }

    componentWillUnmount() { }

    componentDidShow() { }

    componentDidHide() { }

    onMpLoginClick() {
        Taro.showLoading({
            title: "登陆中",
        })
        setTimeout(() => {
            Taro.hideLoading()
        }, 5000)

        Taro.login({
            success: function (res) {
                if (res.code) {
                    wxLogin({ code: res.code })
                        .then((res) => {
                            if (res.data.code === 0) {
                                store.Set(store.TokenKey, res.data.data.token)
                                Taro.navigateBack({
                                    delta: 1,
                                })
                            } else {
                                Taro.showToast({
                                    title: res.errMsg,
                                    icon: 'error',
                                    duration: 500,
                                })
                            }
                        })
                        .catch((res) => {
                            Taro.showToast({
                                title: res.errMsg,
                                icon: 'error',
                                duration: 500,
                            })
                        })
                }
            },
            fail: function (failRes) {
                Taro.showToast({
                    title: failRes.errMsg,
                    icon: 'error',
                    duration: 500,
                })
            }
        })
    }

    render() {
        return (
            <View className='container'>
                <AtButton className='button' type='primary' size='normal' circle onClick={this.onMpLoginClick}>
                    <View className='loginbtn-inner'>
                        <AtIcon value='star' size='20' color='#fff'></AtIcon>
                        <Text className='loginbtn-text'>微信授权登录</Text>
                    </View>
                </AtButton>
            </View>
        )
    }
}