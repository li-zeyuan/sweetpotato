import Taro from '@tarojs/taro'
import { Component } from 'react'
import { AtButton, AtIcon, AtAvatar } from 'taro-ui'
import { wxLogin } from '../../api/api'
import store from '../../utils/store'
import { View, Text, Image } from '@tarojs/components'
import * as images from '../../assets/images/index';
import './index.less'

export default class Index extends Component {
    componentWillMount() { }

    componentDidMount() { }

    componentWillUnmount() { }

    componentDidShow() { }

    componentDidHide() { }

    onLoginCancelClick = () => {
        Taro.navigateBack({
            delta: 1,
        })
    }

    onLoginClick = () => {
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
                        .catch((err) => {
                            Taro.showToast({
                                title: err.errMsg,
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
            <View className='index'>
                <Image className='logo' src={images.logo} />
                <View className='container'>
                    <AtButton className='button' type='secondary' size='normal' circle onClick={this.onLoginCancelClick}>
                            <Text className='cancel-text'>取消</Text>
                    </AtButton>
                    <AtButton className='button' type='primary' size='normal' circle onClick={this.onLoginClick}>
                            <Text className='loginbtn-text'>微信登录</Text>
                    </AtButton>
                </View>
            </View>
        )
    }
}