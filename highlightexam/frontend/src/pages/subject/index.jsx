import { Component } from 'react'
import { View, Text, Image } from '@tarojs/components'
import { study, subjectList } from '../../api/api'
import Taro from '@tarojs/taro'
import { AtButton, AtDivider } from 'taro-ui'
import * as images from '../../assets/images/index'
import Studying from '../../components/studying'
import './index.less'

export default class Index extends Component {
  state = {
    studying: {},
    others: []
  }

  componentWillMount() { }

  componentDidMount() { }

  componentWillUnmount() { }

  componentDidShow() {
    this.loadSubjectList()
  }

  componentDidHide() { }

  loadSubjectList = () => {
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

  handleStudyClick = (id) => {
    study({
      id: id
    })
      .then((res) => {
        if (res.data.code === 0) {
          Taro.switchTab({
            url: '/pages/index/index'
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
    const { studying, others } = this.state
    return (
      <View className='index'>
        <Studying
          studying={studying}
        />

        <View className='otherSubject'>
          <Image className='list_titel_icon' src={images.subjectOther} />
          <Text className='list_titel'>其他题库: </Text>
          <View className='other_subject_list'>
            <AtDivider height={10} />
            {
              others.map((item) => {
                return (
                  <View className='other_subject_list_item' key={item.id}>
                    <View className='other_subject_list_item_text'>
                      <Text className='other_subject_list_item_name'>{item.name}</Text>
                      <Text className='other_subject_list_item_description'>{item.description}</Text>
                    </View>
                    <AtButton
                      className='other_subject_list_item_study_btn'
                      circle
                      type='primary'
                      size='small'
                      onClick={() => {
                        this.handleStudyClick(item.id)
                      }}
                    >学习</AtButton>
                    <AtButton
                      className='other_subject_list_item_detail_btn'
                      circle
                      type='secondary'
                      size='small'
                      onClick={() => {
                        Taro.navigateTo({
                          url: `/pages/knowledge/index?subject_id=${item.id}`
                        })
                      }}
                    >详情</AtButton>
                  </View>
                )
              })
            }
          </View>


          {/* <AtList>
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
          </AtList> */}
        </View>
      </View>
    )
  }
}
