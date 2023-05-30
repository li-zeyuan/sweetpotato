import { Component } from 'react'
import { View, Text, Input } from '@tarojs/components'
import './index.less'
import { AtAvatar, AtIcon } from 'taro-ui'
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
        studyTotalDay: 0,
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
                            portrait: images.avatarBoy,
                            currentSubjectId: res.data.data.current_subject_id,
                            studyTotalDay: res.data.data.study_total_day,
                            studyNum: res.data.data.study_num // 设置学习数量
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

    studyRecordOnClick = () => {
        Taro.navigateTo({
            url: '/pages/studyrecord/index'
        })
    }

    loginOnClick = () => {
        Taro.navigateTo({
            url: '/pages/login/index'
        })
    }

    render() {
        const { portrait, name, uid, gender, currentSubjectId, studyTotalDay, studyNum } = this.state
        const isLoggedIn = store.Get(store.TokenKey)

        return (
            <View className='index'>
                <View className='profile'>
                    <AtAvatar className='avatar' circle image={portrait} size='normal'></AtAvatar>

                    {
                        isLoggedIn &&
                        <View className='userInfo'>
                            <Text className='name'>姓名: {name}</Text>
                            <Text className='uid'>UID: {uid}</Text>
                        </View>
                    }

                    {
                       !isLoggedIn &&
                       <View className='userInfo'>
                           <Text className='pls_login' onClick={this.loginOnClick}>请登陆 >></Text>
                       </View> 
                    }


                </View>

                <View id='container'>
                    <View className='item'>
                        <Input id='study_num_input' disabled type='number' value={studyNum} />
                        <Text className='item_text'>今日数量</Text>
                    </View>
                    <View className='item'>
                        <Input id='study_day_input' type='number' disabled value={studyTotalDay} />
                        <Text className='item_text'>学习天数</Text>
                    </View>
                    <View className='item' onClick={this.studyRecordOnClick}>
                        <AtIcon id='study_record_img' value='list' size='50' color='black'></AtIcon>
                        <Text className='item_text' id='study_record_text'>学习记录</Text>
                    </View>
                </View>
            </View>
        )
    }
}
