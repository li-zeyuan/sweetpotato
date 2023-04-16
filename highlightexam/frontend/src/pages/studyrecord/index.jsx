import React, { Component } from 'react'
import { View, ScrollView, Text } from '@tarojs/components'
import './index.less'
import VirtualList from "@tarojs/components/virtual-list";
import { AtList, AtDivider, AtProgress } from 'taro-ui'
import Taro from '@tarojs/taro'
import store from '../../utils/store'
import { studyrecord } from '../../api/api'
import myEnum from '../../utils/enum'
import PageTail from '../../components/page_tail'

export default class Index extends Component {
    state = {
        list: [],
        hasMore: true
    }

    componentWillMount() {
        if (!store.Get(store.TokenKey)) {
            Taro.navigateTo({
                url: '/pages/login/index'
            })
        }
    }

    componentDidMount() {

    }

    componentWillUnmount() { }

    componentDidShow() {
        this.loadingData(0)
    }

    componentDidHide() { }

    loadingData = (start) => {
        if (!this.state.hasMore) {
            return
        }

        studyrecord({
            start: start,
            limit: myEnum.DefaultRequestLimit
        })
            .then((res) => {
                if (res.data.code === 0) {
                    this.setState({
                        list: this.state.list.concat(res.data.data.list),
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

    render() {
        const { list } = this.state
        return (
            <View>
                <ScrollView
                  className='scrollview'
                  scrollY
                  enhanced
                  scrollWithAnimation
                  style={{height:"680px"}}
                  lowerThreshold={50}
                  onScrollToLower={() => {
                      this.loadingData(this.state.list.length)
                  }}
                >
                    <View>
                    {
                        list.map((item) => {
                            let percentNum = item.studied / item.total * 100
                            return (
                                <View className='item' key={item.subject_id}>
                                    <View className='item_text_container'>
                                        <Text className='item_text'>{item.subject_name}</Text>
                                        <Text className='item_total'>(总{item.total})</Text>
                                    </View>
                                    <AtProgress className='at-progress' percent={Math.round(percentNum)} />
                                    <AtDivider className='at-divider' height={10} />
                                </View>
                            )
                        })
                    }
                    </View>
                    <PageTail
                      hasMore={this.state.hasMore}
                    />
                </ScrollView>
                
            </View>

        )
    }
}
