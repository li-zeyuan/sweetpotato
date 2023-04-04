import React, { Component } from 'react'
import { View } from '@tarojs/components'
import './index.less'
import VirtualList from "@tarojs/components/virtual-list";
import { AtList, AtListItem } from 'taro-ui'
import Taro from '@tarojs/taro'
import store from '../../utils/store'
import { studyrecord } from '../../api/api'
import myEnum from '../../utils/enum'

const Row = React.memo(({ id, index, data }) => {
    return (
        <View id={id} className={index % 2 ? 'ListItemOdd' : 'ListItemEven'}>
            Row {index} : {data[index]}
        </View>
    )
})

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
        this.loadingData()
     }

    componentDidHide() { }

    loadingData = () => {
        studyrecord({
            start: 0,
            limit: 20
        })
            .then((res) => {
                if (res.data.code === 0) {
                    debugger
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

    render() {
        const { list } = this.state
        const dataLen = list.length
        return (
            <VirtualList
              height={800} /* 列表的高度 */
              width='100%' /* 列表的宽度 */
              item={Row} /* 列表单项组件，这里只能传入一个组件 */
              itemData={list} /* 渲染列表的数据 */
              itemCount={dataLen} /* 渲染列表的长度 */
              itemSize={100} /* 列表单项的高度  */
            />
        )
    }
}
