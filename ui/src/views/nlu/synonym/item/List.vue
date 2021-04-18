<template>
  <div>
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
              <a-col :md="!advanced && 8 || 24" :sm="24">
                <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
                  <a-button type="primary" @click="$refs.table.refresh(true)">{{ $t('form.search') }}</a-button>
                  <a-button style="margin-left: 8px" @click="() => this.queryParam = {}">{{ $t('form.reset') }}</a-button>
                </span>
              </a-col>
            </a-row>
          </a-form>
        </div>

        <div class="table-operator">
          <a-button type="primary" icon="plus" @click="create">{{ $t('form.create') }}</a-button>
          <a-dropdown v-if="selectedRowKeys.length > 0">
            <a-menu slot="overlay" @click="multiOpt">
              <a-menu-item key="remove"><a-icon type="delete" />{{ $t('form.remove') }}</a-menu-item>
            </a-menu>
            <a-button style="margin-left: 8px">
              批量操作 <a-icon type="down" />
            </a-button>
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
    <a-modal
      :visible="editVisible"
      :title="modelId > 0 ? $t('menu.synonym.edit.item') : $t('menu.synonym.create.item')"
      :footer="false"
      :centered="true"
      :width="700"
      @cancel="cancelEdit"
    >
      <SynonymItemEdit
        :v-if="editVisible"
        :id="modelId"
        :afterSave="saveModel"
      />
    </a-modal>
  </div>
</template>

<script>
import moment from 'moment'
import { STable, Ellipsis } from '@/components'
import { listSynonymItem, disableSynonymItem, removeSynonymItem, batchRemoveSynonymItem } from '@/api/manage'

import StepByStepModal from '../../../list/modules/StepByStepModal'
import CreateForm from '../../../list/modules/CreateForm'
import SynonymItemEdit from './Edit'

export default {
  name: 'SynonymItems',
  components: {
    STable,
    Ellipsis,
    CreateForm,
    StepByStepModal,
    SynonymItemEdit
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
        return listSynonymItem(requestParameters)
          .then(res => {
            return res
          })
      },
      selectedRowKeys: [],
      selectedRows: [],
      editVisible: false,
      modelId: -1
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
      this.modelId = new Date().getTime() * -1
      this.editVisible = true
    },
    edit (record) {
      this.modelId = record.id
      this.editVisible = true
    },
    cancelEdit () {
      this.editVisible = false
    },
    saveModel () {
      this.editVisible = false
      this.$refs.table.refresh(false)
    },
    disable (record) {
      disableSynonymItem(record).then(json => {
        console.log('disableSynonymItem', json)
        this.$refs.table.refresh(false)
      })
    },
    confirmRemove (record) {
      removeSynonymItem(record).then(json => {
        console.log('removeSynonymItem', json)
        this.$refs.table.refresh(false)
      })
    },
    cancelRemove (e) {
      console.log(e)
    },
    multiOpt ({ key }) {
      console.log(`${key}`)
      batchRemoveSynonymItem(this.selectedRowKeys).then(json => {
        console.log('batchRemoveSynonymItem', json)
        this.$refs.table.refresh(false)
      })
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
