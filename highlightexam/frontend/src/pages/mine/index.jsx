import { Component } from 'react'
import { View, Text,Input } from '@tarojs/components'
import './index.less'
import { AtAvatar, AtGrid } from 'taro-ui'
import Taro from '@tarojs/taro'
import store from '../../utils/store'
import { userProfile } from '../../api/api'
import * as images from '../../assets/images/index';


export default class Index extends Component {
    state = {
        uid: 0,
        name: '',
        gender: 0,
        portrait: '',
        currentSubjectId: 0,
        studyTotal: 0,
        studyNum: 0
    };

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
            userProfile({})
                .then((res) => {
                    if (res.data.code === 0) {
                        this.setState({
                            uid: res.data.data.uid,
                            name: res.data.data.name,
                            gender: res.data.data.gender,
                            portrait: images.avatar,
                            currentSubjectId: res.data.data.current_subject_id,
                            studyTotal: res.data.data.study_total,
                            studyNum: res.data.data.study_num
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

    atGridOnClick(item, index) {
        switch (index) {
            case 1:
                Taro.navigateTo({
                    url: '/pages/studyrecord/index'
                })
                break
        }
    }

    render() {
        const { portrait, name, uid, gender, currentSubjectId, studyTotal, studyNum } = this.state
        return (
            <View className='index'>
                <View className='profile'>
                    <View className='userAvatar'>
                        <AtAvatar className='avatar' circle image={portrait} size='normal'></AtAvatar>
                    </View>
                    <View className='userInfo'>
                        <Text className='name'>姓名: {name}</Text>
                        <Text className='uid'>uid: {uid % 10000000}</Text>
                    </View>
                </View>

                <View>
                    <View>
                        <Input type='number' value={studyNum} />
                        <Text>今日数量</Text>
                    </View>
                    <View>
                        <Input type='number' value='0' />
                        <Text>学习记录</Text>
                    </View>
                    <View>
                        <Text>{studyTotal}</Text>
                        <Text>累计学习</Text>
                    </View>
                </View>
            </View>
        )
    }
}
