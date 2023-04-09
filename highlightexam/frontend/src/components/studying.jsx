import { View, Image, Text } from '@tarojs/components'
import { Component } from 'react'
import { AtList, AtListItem } from 'taro-ui'
import * as images from '../assets/images/index'
import Taro from '@tarojs/taro'
import './studying.less'

export default class Studying extends Component {

    render() {
        let { studying } = this.props

        return (
            <View className='container'>
                <Image className='list_titel_icon' src={images.subjectStudying} />
                <Text className='list_titel'>正在学习: </Text>

                {
                    ((studying.name != undefined) &&
                    (studying.name.length) > 0) &&
                    <AtList>
                        <AtListItem
                          onClick={() => { Taro.navigateTo({ url: `/pages/knowledge/index?subject_id=${studying.id}` }) }}
                          title={studying.name}
                          arrow='right'
                          note={studying.description}
                          extraText='详情'
                        />
                    </AtList>
                }

            </View>
        )
    }
}