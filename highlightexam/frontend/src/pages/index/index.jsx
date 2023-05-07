import { Component } from 'react'
import { View, Text } from '@tarojs/components'
import './index.less'
import { AtButton } from 'taro-ui'
import { knowledgeList, subjectList, studyDoing } from '../../api/api'
import Taro from '@tarojs/taro'
import store from '../../utils/store'
import mEnum from '../../utils/enum'
import num from '../../utils/num'
import Studying from '../../components/studying'

export default class Index extends Component {
  state = {
    isShowTodayCompleted: false,
    studyKnowledgeIndex: 0,
    studying: {
      id: 0,
      name: '',
      description: '',
    },
    willStudyKnowledge: {},
    knowledge: {
      id: 0,
      name: mEnum.DefaultKnowledgeName,
      description: mEnum.DefaultKnowledgeDescription,
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
        console.log(err)
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
            this.setState({
              studyKnowledgeIndex: 0,
              willStudyKnowledge: res.data.data,
            })

            const list = res.data.data.list
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
                  description: mEnum.FinishKnowledgeDescription,
                  other: {}
                }
              })
            }
          }
        })
        .catch((err) => {
          console.log("getKnowledgeList", err)
          Taro.showToast({
            title: err.message,
            icon: 'error',
            duration: 500,
          })
        })
    }
  }

  forgetOnClick = () => {
    const token = store.Get(store.TokenKey)
    if (!token) {
      Taro.navigateTo({
        url: '/pages/login/index'
      })

      return
    }

    if (this.state.knowledge.id === 0) {
      Taro.switchTab({
        url: '/pages/subject/index'
      })
    } else if (this.state.knowledge.id === mEnum.FinishKnowledgeID) {
      this.studyFinish()
      return
    } else {
      let idx = this.state.studyKnowledgeIndex
      let kList = this.state.willStudyKnowledge.list
      // last index
      if (idx + 1 == kList.length) {
        return
      }

      let forgetK = kList[idx]
      let swapIdx = num.Random(idx + 1, kList.length - 1)
      let tailK = kList[swapIdx]
      kList[idx] = tailK
      kList[swapIdx] = forgetK

      this.setState({
        knowledge: tailK
      })
    }
  }
  knowOnClick = () => {
    const token = store.Get(store.TokenKey)
    if (!token) {
      Taro.navigateTo({
        url: '/pages/login/index'
      })

      return
    }

    if (this.state.knowledge.id === 0) {
      Taro.switchTab({
        url: '/pages/subject/index'
      })

      return
    }

    if (this.state.knowledge.id === mEnum.FinishKnowledgeID) {
      this.studyFinish()
      return
    }

    studyDoing({
      subject_id: this.state.studying.id,
      knowledge_id: this.state.knowledge.id
    })
      .then((res) => {
        if (res.data.code === 0) {
          // show completed today study num
          if (!this.state.isShowTodayCompleted && res.data.data.is_completed_today) {
            Taro.showToast({
              title: "今日学习已完成",
              icon: 'success',
              duration: 2000,
            })

            this.setState({
              isShowTodayCompleted: true,
            })
          }

          const willK = this.state.willStudyKnowledge
          let index = this.state.studyKnowledgeIndex
          if (index + 1 < willK.list.length) {
            const k = willK.list[index + 1]
            this.setState({
              studyKnowledgeIndex: index + 1,
              knowledge: {
                id: k.id,
                name: k.name,
                description: k.description,
                other: k.other
              }
            })
          } else {
            this.setState(
              {
                knowledge: {
                  id: mEnum.FinishKnowledgeID,
                  name: mEnum.FinishKnowledgeName,
                  description: mEnum.FinishKnowledgeDescription,
                  other: {}
                }
              },
              this.studyFinish
            )
          }
        }
      })
      .catch((err) => {
        console.log("study doing error: ", err)
        Taro.showToast({
          title: err.message,
          icon: 'error',
          duration: 500,
        })
      })
  }

  studyFinish = () => {
    if (this.state.knowledge.id !== mEnum.FinishKnowledgeID) {
      return
    }

    setTimeout(() => {
      Taro.switchTab({
        url: '/pages/subject/index'
      })
    }, 2000);

    Taro.showToast({
      title: '恭喜您学习完成',
      icon: 'success',
      duration: 2000,
    })
  }

  render() {
    const { studying, knowledge } = this.state

    return (
      <View className='index' >
        <Studying
          studying={studying}
        />

        <View className='learning'>
          <Text className='pinyin'>{knowledge.other.pinyin ? "[" + knowledge.other.pinyin + "]" : ""}</Text>
          <Text className='knowledgeName'>{knowledge.name}</Text>
          <Text className='knowledgeDescription'>{knowledge.description}</Text>

          <AtButton className='buttonKnown' type='primary' circle onClick={this.knowOnClick}>认识</AtButton>
          <AtButton className='buttonForget' type='secondary' circle onClick={this.forgetOnClick}>忘记</AtButton>
        </View>
      </View >
    )
  }
}
