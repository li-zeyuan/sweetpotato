import { View, ScrollView, Text } from '@tarojs/components'
import { Component } from 'react'
import { AtList, AtListItem } from 'taro-ui'
import './list.less'

export default class List extends Component {
    render() {
        return (
            <ScrollView
              className='scrollview'
              scrollY
              enhanced
              
              scrollWithAnimation
              style={this.props.scrollStyle}
              lowerThreshold={50}
              onScrollToLower={this.props.onScrollToLower}
            >
                <AtList>
                    {
                        this.props.list.map((item) => {
                            return (
                                <AtListItem
                                  key={item.key}
                                  title={item.title}
                                  note={item.note}
                                />
                            )
                        })
                    }
                </AtList>
                <View id='page_tail'>{this.props.hasMore ? <Text>下拉加载更多</Text> : <Text>家人们到底了</Text>}</View>
            </ScrollView>
        )
    }
}