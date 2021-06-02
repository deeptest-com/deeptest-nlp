<template>
  <page-header-wrapper
    :title="model.name"
  >
    <template v-slot:content>
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
              result
            </div>
            <div class="json" v-if="viewMode==='json'">
              json
            </div>
          </div>
        </div>

        <div class="form">
          <div class="left">
            <a-input
              v-model="question"
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
import { nluRequest } from '@/api/manage'

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
      question: '',
      messages: [],
      histories: [],
      historyIndex: 0,

      viewMode: '',
      chatHeight: {}
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
      overflow: 'auto'

    }
    this.loadData()
  },
  created () {
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
    },
    view (mode) {
      console.log('view')
      if (mode === 'result') {
        this.viewMode = 'result'
      } else if (mode === 'json') {
        this.viewMode = 'json'
      } else {
        this.viewMode = ''
      }
    },
    send () {
      console.log('send', this.question)

      this.histories.push(this.question)
      this.historyIndex = this.histories.length
      this.messages.push({ type: 'question', content: this.question })

      nluRequest(this.id, this.question).then(json => {
        console.log('nluRequest', json)
        const data = json.data
        if (data.code === 1) { // success
          this.messages.push({ type: 'answer', content: data.result })
        } else {
          this.messages.push({ type: 'pardon' })
        }

        this.question = ''
        this.scroll()
      })
    },
    up () {
      console.log('up')
      if (this.historyIndex > 0) this.historyIndex--
      this.question = this.histories[this.historyIndex]
    },
    down () {
      console.log('down')
      if (this.historyIndex < this.histories.length - 1) this.historyIndex++
      this.question = this.histories[this.historyIndex]
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

</style>
