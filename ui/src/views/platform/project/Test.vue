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
        <div class="request">
          <a-comment name="chat">
            <a-avatar slot="avatar" icon="android" class="my-avatar-icon"/>
            <span slot="author">
              <span>{{ $t('msg.testing.can.not.understand') }}</span>
            </span>
            <span slot="content">
              <a @click="view('result')">{{ $t('common.view.result') }}</a>
              &nbsp;&nbsp;&nbsp;  | &nbsp;&nbsp;&nbsp;
              <a @click="view('json')">{{ $t('common.view.json') }}</a>
              &nbsp;&nbsp;&nbsp; | &nbsp;&nbsp;&nbsp;
              <a @click="view('nothing')">{{ $t('common.view.nothing') }}</a>
            </span>
          </a-comment>

          <div class="form">
            <div class="left">
              <a-input v-model="data.text" placeholder=""/>
            </div>
            <div class="right">
              <a-button type="primary" @click="send()">{{ $t('form.send') }}</a-button>
            </div>
          </div>
        </div>
        <div class="result" v-if="viewMode==='result'">
          result
        </div>
        <div class="json" v-if="viewMode==='json'">
          json
        </div>

      </div>
    </a-card>

  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { getProject } from '@/api/manage'

import { baseMixin } from '@/store/app-mixin'

export default {
  name: 'ProjectEdit',
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
      data: {},
      viewMode: '',
      moment
    }
  },
  watch: {
    id: function () {
      console.log('watch id', this.id)
      this.loadData()
    }
  },
  mounted () {
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
      if (!this.id) {
        return
      }
      if (this.id) {
        getProject(this.id).then(json => {
          this.model = json.data
        })
      } else {
        this.reset()
      }
    },
    view (mode) {
      console.log('view')
      if (mode === 'result') {
        this.viewMode = 'result'
      } else if (mode === 'json') {
        this.viewMode = 'json'
      } else if (mode === 'nothing') {
        this.viewMode = ''
      }
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
  display: flex;

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

    .form {
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

  .result {
    flex: 1;
    padding: 10px 20px;
  }
  .json {
    flex: 1;
    padding: 10px 20px;
  }
}

</style>
