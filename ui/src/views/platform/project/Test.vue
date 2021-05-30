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
      <div class="testing" :style="chatHeight">
        <div class="row1">
          <div class="request">
            <div v-for="(item, index) in messages" :value="item.id" :key="index">
              <ChatMessage :msg="item" @view="view" />
            </div>
          </div>
          <div class="response">
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
            <a-input v-model="question.text" placeholder=""/>
          </div>
          <div class="right">
            <a-button type="primary" @click="send">{{ $t('form.send') }}</a-button>
          </div>
        </div>

      </div>
    </a-card>

  </page-header-wrapper>
</template>

<script>

import { baseMixin } from '@/store/app-mixin'
import ChatMessage from './component/Message'
import { nluRequest, requestSuccess } from '@/api/manage'

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
      question: {},
      messages: [],

      viewMode: '',
      chatHeight: 0
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
      height: `${200}px`
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
      console.log('send', this.question.text)

      nluRequest(this.id, this.question.text).then(json => {
        console.log('nluRequest', json)
        if (requestSuccess(json.code)) {
        }
      })
    },
    back () {
      this.$router.push('/platform/project/list')
    }
  }
}
</script>

<style lang="less" scoped>

.detail-layout {
  margin-left: 44px;
}

.testing {

  .row1 {
    display: flex;
    height: calc(100% - 30px);
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
    width: 90%;

    display: flex;
    margin-left: 10px;

    .left {
      flex: 1;
      margin-right: 10px;
    }

    .right {
      width: 60px;
    }
  }
}

</style>
