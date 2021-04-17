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
        <a-dropdown v-action:edit v-if="selectedRowKeys.length > 0">
          <a-menu slot="overlay">
            <a-menu-item key="1"><a-icon type="delete" />{{ $t('form.remove') }}<</a-menu-item>
          </a-menu>
        </a-dropdown>
      </div>

      <s-table
        ref="table"
        size="default"
        rowKey="id"
        :columns="columns"
        :data="loadData"
        :alert="true"
        :rowSelection="rowSelection"
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
            <a v-if="!record.disabled" @click="disable(record)">{{ $t('form.disable') }}</a>
            <a v-if="record.disabled" @click="disable(record)">{{ $t('form.enable') }}</a>
            <a-divider type="vertical" />
            <a-popconfirm
              v-if="!record.isDefault"
              :title="$t('form.confirmToRemove')"
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
import { listLookup, disableLookup, removeLookup } from '@/api/manage'

import StepByStepModal from '../../list/modules/StepByStepModal'
import CreateForm from '../../list/modules/CreateForm'

export default {
  name: 'LookupList',
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
        return listLookup(requestParameters)
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
        title: this.$t('common.no'),
        scopedSlots: { customRender: 'serial' }
      },
      {
        title: this.$t('common.name'),
        dataIndex: 'name'
      },
      {
        title: this.$t('common.status'),
        dataIndex: 'status',
        scopedSlots: { customRender: 'status' }
      },
      {
        title: this.$t('common.opt'),
        dataIndex: 'action',
        width: '180px',
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
    rowSelection () {
      return {
        selectedRowKeys: this.selectedRowKeys,
        onChange: this.onSelectChange
      }
    }
  },
  methods: {
    create () {
      this.mdl = null
      this.visible = true

      this.$router.push('/nlu/lookup/0/edit')
    },
    edit (record) {
      this.visible = true
      this.mdl = { ...record }

      this.$router.push('/nlu/lookup/' + record.id + '/edit')
    },
    disable (record) {
      disableLookup(record).then(json => {
        console.log('disableLookup', json)
        this.$refs.table.refresh(false)
      })
    },
    confirmRemove (record) {
      removeLookup(record).then(json => {
        console.log('removeLookup', json)
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
