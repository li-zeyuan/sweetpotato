import { Component } from 'react'
import { View, Text } from '@tarojs/components'
import './index.less'
import { AtButton, AtList, AtListItem } from 'taro-ui'
import { knowledgeList, subjectList, studyDoing } from '../../api/api'
import Taro from '@tarojs/taro'
import store from '../../utils/store'
import mEnum from '../../utils/enum'

export default class Index extends Component {
  state = {
    studying: {
      id: 0,
      name: '',
      description: '',
    },
    willStudyKnowledge: {
      index: 0,
      list: [],
      hasMore: false,
      hasStudied: 0
    },
    knowledge: {
      id: 0,
      name: 'Highlight事考',
      description: '致力于未来市长、书记、厅长您，顺利上岸！',
      other: {}
    }
  }

  componentWillMount() { }

  componentDidMount() { }

  componentWillUnmount() { }

  componentDidShow() {
    this.getSubjectList()
  }

  componentDidHide() { }

  getSubjectList = () => {
    subjectList({
      only_mine: true
    })
      .then((res) => {
        const studying = res.data.data.studying
        if (res.data.code === 0) {
          this.setState({
            studying: {
              id: studying.id,
              name: studying.name,
              description: studying.description
            }
          })

          if (studying.id > 0) {
            this.getKnowledgeList(res.data.data.studying.id)
          }
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

  getKnowledgeList = (sudjectId) => {
    const token = store.Get(store.TokenKey)
    if (token) {
      knowledgeList({
        subject_id: sudjectId
      })
        .then((res) => {
          if (res.data.code === 0) {
            const list = res.data.data.list
            this.setState({
              willStudyKnowledge: {
                list: list,
                hasStudied: res.data.data.has_studied,
                hasMore: res.data.data.has_more
              }
            })

            if (list.length) {
              this.setState({
                knowledge: {
                  id: list[0].id,
                  name: list[0].name,
                  description: list[0].description,
                  other: list[0].other
                }
              })
            } else {
              this.setState({
                knowledge: {
                  id: mEnum.FinishKnowledgeID,
                  name: mEnum.FinishKnowledgeName,
                  description: '',
                  other: ''
                }
              })
            }
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

  forgetOnClick = () => {
    if (this.state.knowledge.id === 0) {
      Taro.navigateTo({
        url: '/pages/login/index'
      })
    }

    if (this.state.knowledge.id === mEnum.FinishKnowledgeID) {
      // todo show toast
      return
    }
  }
  knowOnClick = () => {
    if (this.state.knowledge.id === 0) {
      Taro.navigateTo({
        url: '/pages/login/index'
      })

      return
    }

    if (this.state.knowledge.id === mEnum.FinishKnowledgeID) {
      // todo show toast
      return
    }

    studyDoing({
      subject_id: this.state.studying.id,
      knowledge_id: this.state.knowledge.id
    })
      .then((res) => {
        if (res.data.code === 0) {
          const willK = this.state.willStudyKnowledge
          if (willK.index < willK.list.length) {
            const k = willK.list[willK.index++]
            this.setState({
              willStudyKnowledge: {
                index: this.state.willStudyKnowledge.index + 1
              },
              knowledge: {
                id: k.id,
                name: k.name,
                description: k.description,
                other: k.other
              }
            })
          } else {
            this.setState({
              knowledge: {
                id: mEnum.FinishKnowledgeID,
                name: mEnum.FinishKnowledgeName,
                description: '',
                other: {}
              }
            })
          }


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
    var { studying, knowledge } = this.state

    return (
      <View className='index' >
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

        <View className='learning'>
          <Text className='knowledgeName'>{knowledge.name}</Text>
          <Text className='knowledgeDescription'>{knowledge.description}</Text>

          <AtButton className='buttonKnown' type='primary' circle onClick={this.knowOnClick}>认识</AtButton>
          <AtButton className='buttonForget' type='secondary' circle onClick={this.forgetOnClick}>忘记</AtButton>
        </View>
      </View >
    )
  }
}
