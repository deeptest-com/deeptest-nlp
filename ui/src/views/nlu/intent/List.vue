<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <div class="table-page-search-wrapper">
        <a-form layout="inline">
          <a-row :gutter="48">
            <a-col :md="8" :sm="24">
              <a-form-item :label="$t('form.name')">
                <a-input v-model="queryParam.keywords" placeholder=""/>
              </a-form-item>
            </a-col>
            <a-col :md="8" :sm="24">
              <a-form-item :label="$t('form.status')">
                <a-select v-model="queryParam.status" placeholder="请选择" default-value="0">
                  <a-select-option value="">{{ $t('form.all') }}</a-select-option>
                  <a-select-option value="true">{{ $t('form.enable') }}</a-select-option>
                  <a-select-option value="false">{{ $t('form.disable') }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <!-- <template v-if="advanced">
            </template> -->
            <a-col :md="!advanced && 8 || 24" :sm="24">
              <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
                <a-button type="primary" @click="$refs.table.refresh(true)">{{ $t('form.search') }}</a-button>
                <a-button style="margin-left: 8px" @click="() => this.queryParam = {}">{{ $t('form.reset') }}</a-button>

                <!-- <a @click="toggleAdvanced" style="margin-left: 8px">
                  {{ advanced ? $t('form.collapse') : $t('form.expand') }}
                  <a-icon :type="advanced ? 'up' : 'down'"/>
                </a> -->
              </span>
            </a-col>
          </a-row>
        </a-form>
      </div>

      <div class="table-operator">
        <a-button type="primary" icon="plus" @click="create">{{ $t('form.create') }}</a-button>
      </div>

      <s-table
        ref="table"
        size="default"
        rowKey="id"
        :columns="columns"
        :data="loadData"
        :alert="true"
        showPagination="auto"
      >
        <span slot="serial" slot-scope="text, record, index">
          {{ index + 1 }}
        </span>

        <span slot="name" slot-scope="text">
          <ellipsis :length="4" tooltip>{{ text }}</ellipsis>
        </span>

        <span slot="status" slot-scope="text, record">
          <a-badge :status="!record.disabled | statusTypeFilter(statusMap)" :text="!record.disabled | statusFilter(statusMap)" />
        </span>

        <span slot="action" slot-scope="text, record">
          <template>
            <a @click="edit(record)">{{ $t('form.edit') }}</a>

            <a-divider type="vertical" />
            <a @click="maintain(record)">{{ $t('form.maintain') }}</a>

            <a-divider type="vertical" />
            <a v-if="!record.disabled" @click="disable(record)">{{ $t('form.disable') }}</a>
            <a v-if="record.disabled" @click="disable(record)">{{ $t('form.enable') }}</a>

            <a-divider type="vertical" />
            <a-popconfirm
              v-if="!record.isDefault"
              :title="$t('form.confirm.to.remove')"
              :okText="$t('form.ok')"
              :cancelText="$t('form.cancel')"
              @confirm="confirmRemove(record)"
              @cancel="cancelRemove"
            >
              <a href="#">{{ $t('form.remove') }}</a>
            </a-popconfirm>
          </template>
        </span>
      </s-table>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { STable, Ellipsis } from '@/components'
import { listIntent, disableIntent, removeIntent } from '@/api/manage'

import StepByStepModal from '../../list/modules/StepByStepModal'
import CreateForm from '../../list/modules/CreateForm'

export default {
  name: 'IntentList',
  components: {
    STable,
    Ellipsis,
    CreateForm,
    StepByStepModal
  },
  columns: [],
  statusMap: {},
  data () {
    return {
      visible: false,
      confirmLoading: false,
      mdl: null,
      advanced: false,
      queryParam: {},
      loadData: parameter => {
        const requestParameters = Object.assign({}, parameter, this.queryParam)
        return listIntent(requestParameters)
          .then(res => {
            return res
          })
      },
      selectedRowKeys: [],
      selectedRows: []
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
  created () {
    this.columns = [
      {
        title: this.$t('form.no'),
        scopedSlots: { customRender: 'serial' }
      },
      {
        title: this.$t('form.name'),
        dataIndex: 'name'
      },
      {
        title: this.$t('form.status'),
        dataIndex: 'status',
        scopedSlots: { customRender: 'status' }
      },
      {
        title: this.$t('form.opt'),
        dataIndex: 'action',
        width: '220px',
        scopedSlots: { customRender: 'action' }
      }
    ]

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
  computed: {
  },
  methods: {
    create () {
      this.mdl = null
      this.visible = true

      this.$router.push('/nlu/intent/0/edit')
    },
    edit (record) {
      this.visible = true
      this.mdl = { ...record }

      this.$router.push('/nlu/intent/' + record.id + '/edit')
    },
    maintain (record) {
      this.visible = true
      this.mdl = { ...record }

      this.$router.push('/nlu/intent/' + record.id + '/items')
    },
    disable (record) {
      disableIntent(record).then(json => {
        console.log('disableIntent', json)
        this.$refs.table.refresh(false)
      })
    },
    confirmRemove (record) {
      removeIntent(record).then(json => {
        console.log('removeIntent', json)
        this.$refs.table.refresh(false)
      })
    },
    cancelRemove (e) {
      console.log(e)
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    resetSearchForm () {
      this.queryParam = {
        date: moment(new Date())
      }
    }
  }
}
</script>
