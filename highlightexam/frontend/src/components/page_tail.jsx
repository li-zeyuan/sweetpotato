import { View, Text } from '@tarojs/components'
import { Component } from 'react'
import './page_tail.less'

export default class PageTail extends Component {
    render() {
        return (
            <View id='page_tail'>
                {this.props.hasMore ? <Text>下拉加载更多</Text> : <Text>家人们到底了</Text>}
            </View>
        )
    }
}