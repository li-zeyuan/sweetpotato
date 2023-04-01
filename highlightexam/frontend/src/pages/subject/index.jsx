import { Component } from 'react'
import { View, Text } from '@tarojs/components'
import { subjectList } from '../../api/api'
import Taro from '@tarojs/taro'
import { AtList, AtListItem } from 'taro-ui'
import './index.less'

export default class Index extends Component {
  state = {
    studying: {},
    others: []
  }

  componentWillMount() {
    subjectList({})
      .then((res) => {
        if (res.data.code === 0) {
          this.setState({
            studying: res.data.data.studying,
            others: res.data.data.others
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

  componentDidMount() { }

  componentWillUnmount() { }

  componentDidShow() { }

  componentDidHide() { }

  atListItemOnClick(a) {
    console.log(a)
  }

  render() {
    const { studying, others } = this.state
    return (
      <View className='index'>
        <View>
          <Text>正在学习: </Text>
          <AtList>
            <AtListItem
              onClick={() => { Taro.navigateTo({ url: `/pages/knowledge/index?subject_id=${studying.id}` }) }}
              title={studying.name}
              arrow='right'
              note={studying.description}
              extraText='详情'
            />
          </AtList>
        </View>

        <View className='otherSubject'>
          <Text>其他题库: </Text>
          <AtList>
            {
              others.map((item) => {
                return (
                  <AtListItem
                    key={item.id}
                    onClick={() => { Taro.navigateTo({ url: `/pages/knowledge/index?subject_id=${item.id}` }) }}
                    title={item.name}
                    arrow='right'
                    note={item.description}
                    extraText='详情'
                  />
                )
              })
            }
          </AtList>
        </View>
      </View>
    )
  }
}
