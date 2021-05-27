<template>
  <page-header-wrapper
    :title="model.name"
  >
    <template v-slot:content>
      <a-descriptions size="small" :column="isMobile ? 1 : 2">
        <a-descriptions-item label="创建人">{{ model.createdBy }}</a-descriptions-item>
        <a-descriptions-item label="所在目录">{{ model.path }}</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ model.createdAt | moment }}</a-descriptions-item>
        <a-descriptions-item label="更新时间">{{ model.updtedAt | moment }}</a-descriptions-item>

        <a-descriptions-item label="描述"></a-descriptions-item>
      </a-descriptions>
    </template>

    <!-- actions -->
    <template v-slot:extra>
      <a-button-group style="margin-right: 4px;">
        <a-button @click="compile()">编译</a-button>
        <a-button @click="training()">训练</a-button>
        <a-button @click="start()">启动</a-button>
      </a-button-group>
      <a-button @click="back()" type="primary">返回</a-button>
    </template>

    <template v-slot:extraContent>
      <a-row class="status-list">
        <a-col :xs="9" :sm="9"></a-col>
        <a-col :xs="12" :sm="12">
          <div class="text">状态</div>
          <div class="heading">
            <a-badge :status="!model.disabled | statusTypeFilter(statusMap)" :text="!model.disabled | statusFilter(statusMap)" />
          </div>
        </a-col>
        <a-col :xs="3" :sm="3"></a-col>
      </a-row>
    </template>

    <!-- 操作 -->
    <a-card style="margin-top: 24px" :bordered="false" title="操作日志">
      <a-table
        rowKey="id"
        :columns="historyColumns"
        :dataSource="model.histories"
        :pagination="false"
      >
        <span slot="serial" slot-scope="text, record, index">
            {{ index + 1 }}
        </span>
        <span slot="action" slot-scope="text, record">
            {{ $t('common.' + record.action) }}
        </span>
        <span slot="createdTime" slot-scope="text, record">
          {{ record.createdTime | moment }}
        </span>
      </a-table>
    </a-card>

  </page-header-wrapper>
</template>

<script>
import { getProject, compileProject, trainingProject, startService } from '@/api/manage'

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

      historyColumns: [
        {
          title: this.$t('form.no'),
          scopedSlots: { customRender: 'serial' }
        },
        {
          title: '操作类型',
          scopedSlots: { customRender: 'action' }
        },
        {
          title: '操作人',
          dataIndex: 'userName',
          key: 'userName'
        },
        {
          title: '操作时间',
          scopedSlots: { customRender: 'createdTime' }
        }
      ]
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
  filters: {
    statusFilter (status, statusMap) {
      return statusMap[status].text
    },
    statusTypeFilter (status, statusMap) {
      return statusMap[status].type
    }
  },
  methods: {
    loadData () {
      if (!this.id) {
        return
      }
      if (this.id) {
        this.getModel().then(json => {
          this.model = json.data
        })
      } else {
        this.reset()
      }
    },
    getModel () {
      return getProject(this.id)
    },
    compile () {
      console.log('compile')
      compileProject(this.model).then(json => {
        console.log('compile', json)

        if (json.code === 200) {
          const that = this
          this.$notification['success']({
            message: that.$t('common.tips'),
            description: that.$t('msg.compile.success'),
            placement: 'bottomRight',
            duration: 8
          })

          this.loadData()
        }
      })
    },
    training () {
      console.log('training')

      trainingProject(this.model).then(json => {
        console.log('training', json)
        if (json.code === 200) {
          const that = this
          this.$notification['success']({
            message: that.$t('common.tips'),
            description: that.$t('msg.training.start'),
            placement: 'bottomRight',
            duration: 8
          })

          this.loadData()
        }
      })
    },
    start () {
      console.log('start')
      startService(this.model).then(json => {
        console.log('training', json)
        if (json.code === 200) {
          this.$notification['success']({
            message: this.$root.$t('common.tips'),
            description: this.$root.$t('msg.service.start'),
            placement: 'bottomRight',
            duration: 8
          })

          this.loadData()
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
.text {
  color: rgba(0, 0, 0, .45);
}

.heading {
  color: rgba(0, 0, 0, .85);
  font-size: 20px;
}

.no-data {
  color: rgba(0, 0, 0, .25);
  text-align: center;
  line-height: 64px;
  font-size: 16px;

  i {
    font-size: 24px;
    margin-right: 16px;
    position: relative;
    top: 3px;
  }
}

.mobile {
  .detail-layout {
    margin-left: unset;
  }
  .text {

  }
  .status-list {
    text-align: left;
  }
}
</style>
