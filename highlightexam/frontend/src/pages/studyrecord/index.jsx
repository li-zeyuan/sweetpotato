import { Component } from 'react'
import { View } from '@tarojs/components'
import './index.less'
import { AtList, AtListItem } from 'taro-ui'
import Taro from '@tarojs/taro'
import store from '../../utils/store'
import { studyrecord } from '../../api/api'


export default class Index extends Component {
    state = {
        list: [],
        hasMore: false
    }

    componentWillMount() {
        if (!store.Get(store.TokenKey)) {
            Taro.navigateTo({
                url: '/pages/login/index'
            })
        }
    }

    componentDidMount() { }

    componentWillUnmount() { }

    componentDidShow() {
        var token = store.Get(store.TokenKey)
        if (token) {
            studyrecord({
                start: 0,
                limit: 20
            })
                .then((res) => {
                    if (res.data.code === 0) {
                        this.setState({
                            list: res.data.data.list,
                            hasMore: res.data.data.has_more
                        })
                    }
                })
                .catch((err) => {
                    Taro.showToast({
                        title: err.message,
                        icon: 'error',
                        duration: 500,
                    })
                })
        }
    }

    componentDidHide() { }

    render() {
        return (
            <View className='index'>
                <AtList>
                    {
                        this.state.list.map((item) => {
                            return (
                                <AtListItem key={item.subject_id} title={item.subject_name} />
                            )
                        })
                    }
                </AtList>
            </View>
        )
    }
}
