<template>
  <page-header-wrapper
    :title="model.name"
  >
    <template v-slot:content>
      <a-row class="agents">
        <a-col class="label">{{$t('form.exec.selenium.on.agent')}}</a-col>
        <a-col class="content">
          <a-select v-model="agentId" class="select">
            <a-select-option v-for="(item, index) in agents" :value="item.id" :key="index">
              {{ item.ip }}
            </a-select-option>
          </a-select>
        </a-col>
      </a-row>
    </template>

    <template v-slot:extra>
      <a-button @click="back()" type="primary">{{ $t('common.back') }}</a-button>
    </template>

    <a-card
      style="margin-top: 24px"
      :bordered="false">
      <div class="testing">
        <div class="row1">
          <div class="request" :style="chatHeight" ref="leftContainer">
            <div v-for="(item, index) in messages" :value="item.id" :key="index">
              <ChatMessage :msg="item" @view="view" />
            </div>
          </div>
          <div class="response" :style="chatHeight" ref="rightContainer" v-if="viewMode!==''">
            <div class="result" v-if="viewMode==='result'">
              <pre class="white-space: pre;">{{ detailResult }}</pre>
            </div>
            <div class="json" v-if="viewMode==='json'">
              <pre>{{ jsonResult }}</pre>
            </div>
          </div>
        </div>

        <div class="form">
          <div class="left">
            <a-input
              v-model="question"
              ref="question"
              @keydown.down="down"
              @keydown.up="up"
              @keyup.enter="send"
              :placeholder="$t('form.input.sent')"/>
          </div>
          <div class="right">
            <a-button type="primary" @click="send">{{ $t('form.send') }}</a-button>
          </div>
        </div>
        <div class="tips">{{ $t('form.nav.history') }}</div>

      </div>
    </a-card>

  </page-header-wrapper>
</template>

<script>

import { baseMixin } from '@/store/app-mixin'
import ChatMessage from './component/Message'
import { testProject, nluRequest } from '@/api/manage'

import storage from 'store'
import { TEST_HISTORIES } from '@/utils/const'
import { setSelectionRange } from '@/utils/domUtil'

export default {
  name: 'ProjectEdit',
  components: {
    ChatMessage
  },
  mixins: [baseMixin],
  statusMap: {},
  props: {
    id: {
      type: Number,
      default: function () {
        return parseInt(this.$route.params.id)
      }
    }
  },
  data () {
    return {
      model: {},
      agents: [],
      agentId: 0,
      question: '',
      messages: [],
      histories: [],
      historyIndex: 0,

      viewMode: '',
      chatHeight: {},
      detailResult: '',
      jsonResult: '',
      answerMap: {}
    }
  },
  watch: {
    id: function () {
      console.log('watch id', this.id)
      this.loadData()
    }
  },
  mounted () {
    this.chatHeight = {
      height: document.body.clientHeight - 330 + 'px',
      'overflow-y': 'auto',
      'overflow-x': 'hidden'
    }
    this.loadData()
  },
  created () {
    this.histories = storage.get(TEST_HISTORIES, [])
    this.historyIndex = this.histories.length

    this.statusMap = {
      true: {
        type: 'processing',
        text: this.$t('status.enable')
      },
      false: {
        type: 'default',
        text: this.$t('status.disable')
      }
    }
  },
  methods: {
    loadData () {
      this.messages = [{ type: 'welcome' }]

      testProject(this.id).then(json => {
        this.agents = json.data.agents
      })
    },
    view (data) {
      const mode = data.mode
      const key = data.key

      console.log('view', mode, key)

      if (mode === 'result') {
        this.viewMode = 'result'
      } else if (mode === 'json') {
        this.viewMode = 'json'
      } else {
        this.viewMode = ''
      }

      this.detailResult = this.answerMap[key].detail
      this.jsonResult = this.answerMap[key].json
    },
    send () {
      console.log('send', this.question)

      this.histories.push(this.question)
      if (this.histories.length > 30) this.histories = this.histories.slice(this.histories.length - 30)
      storage.set(TEST_HISTORIES, this.histories)

      this.historyIndex = this.histories.length
      this.messages.push({ type: 'question', content: this.question })

      nluRequest(this.id, this.question, this.agentId).then(json => {
        console.log('nluRequest', json)
        const data = json.data

        const key = new Date().getTime()
        if (data.code === 1) { // success
          const msg = data.result.intent.name

          const slots = []
          data.result.entities.forEach((item, index) => {
            const name = item.entity

            let value = item.valueOrigin
            if (value === '') {
              value = data.result.text.substr(item.start, item.end - item.start)
            }

            slots.push('  ' + name + ' = ' + value)
          })

          let detailResult = this.$t('menu.intent') + ': ' + msg
          detailResult += '\n' + this.$t('common.nlu.confidence') + ': ' + (data.result.intent.confidence * 100) + '%'
          detailResult += '\n' + this.$t('common.nlu.slots') + ': ' + slots.length + '\n' + slots.join('\n')

          this.answerMap[key] = { detail: detailResult, json: JSON.stringify(data.result, null, 3) }
          this.detailResult = this.answerMap[key].detail
          this.jsonResult = this.answerMap[key].json

          this.messages.push({ key: key, type: 'answer', content: msg })
        } else {
          const msg = this.$t('msg.' + data.result.i118, { 'content': data.result.content })
          this.answerMap[key] = { detail: msg, json: data.result }
          this.detailResult = this.answerMap[key].detail
          this.jsonResult = this.answerMap[key].json

          this.messages.push({ key: key, type: 'pardon' })
        }

        this.question = ''
        this.scroll()
      })
    },
    up () {
      console.log('up')
      if (this.historyIndex > 0) this.historyIndex--

      this.question = this.histories[this.historyIndex]
      setSelectionRange(this.$refs.question.$el, this.question.length)
    },
    down () {
      console.log('down')
      if (this.historyIndex < this.histories.length - 1) this.historyIndex++

      this.question = this.histories[this.historyIndex]
      setSelectionRange(this.$refs.question.$el, this.question.length)
    },
    scroll () {
      setTimeout(() => {
        const leftContainer = this.$refs.leftContainer
        leftContainer.scrollTop = leftContainer.scrollHeight
        const rightContainer = this.$refs.rightContainer
        if (rightContainer) rightContainer.scrollTop = rightContainer.scrollHeight
      }, 300)
    },
    back () {
      this.$router.push('/project/list')
    }
  }
}
</script>

<style lang="less" scoped>
.ant-page-header {
  padding-bottom: 0 !important;
}

.detail-layout {
  margin-left: 44px;
}

.testing {
  padding-bottom: 0px;

  .row1 {
    display: flex;
    height: calc(100% - 45px);
    margin-bottom: 10px;
    .request {
      flex: 1;

      .chat {
        .ant-comment-content-author {
          margin-bottom: 0px !important;

          span {
            font-size: 14px;
          }
        }

        .ant-comment-content-detail {
          a {
            font-size: 12px;
          }
        }
      }
    }
    .response {
      flex: 1;
      padding: 10px 20px;
      border-left: 1px solid #e8e8e8;

      .result {
        height: 100%;
        pre {
          height: 100%;
        }
      }
      .json {
        height: 100%;
        pre {
          height: 100%;
        }
      }
    }
  }

  .form {
    margin-top: 10px;
    margin-left: 10px;
    width: 90%;

    display: flex;

    .left {
      flex: 1;
      margin-right: 10px;
    }

    .right {
      width: 60px;
    }
  }
  .tips {
    margin-left: 10px;
  }
}

.title {
  font-weight: bolder;
  font-size: 16px;
}
.agents {
  display: flex;
  .label {
    width: 110px;
    line-height: 28px;
  }
  .content {
    flex: 1;
    .select {
      width: 200px;
    }
  }
}

</style>
