import { Component } from 'react'
import { View } from '@tarojs/components'
import { subjectDetail } from '../../api/api'
import Taro from '@tarojs/taro'
import List from '../../components/list'
import myEnum from '../../utils/enum'

export default class Index extends Component {
  inst = Taro.getCurrentInstance()

  state = {
    list: [],
    hasMore: false,
  }

  componentWillMount() {
    this.dataLoading()
  }

  componentDidMount() { }

  componentWillUnmount() { }

  componentDidShow() { }

  componentDidHide() { }

  dataLoading = () => {
    if (!this.state.hasMore & this.state.list.length > 0) {
      return
    }

    subjectDetail({
      id: this.inst.router.params.subject_id,
      start: this.state.list.length,
      limit: myEnum.DefaultRequestLimit,
    })
      .then((res) => {
        if (res.data.code === 0) {
          let tempList = []
          res.data.data.list.map((item)=>{
            tempList.push({
              key:item.id,
              title:item.name,
              note:item.description
            })
          })

          this.setState({
            list: this.state.list.concat(tempList),
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
    return (
      <View className='index'>
        <List
          hasMore={this.state.hasMore}
          list={this.state.list}
          scrollStyle={{ height: '680px' }}
          onScrollToLower={this.dataLoading}
        />
      </View>
    )
  }
}
