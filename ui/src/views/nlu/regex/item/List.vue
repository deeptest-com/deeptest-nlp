<template>
  <div>
    <page-header-wrapper>
      <div class="toolbar-edit">
        <div class="left"></div>
        <div class="right">
          <a-button @click="back()" type="primary">{{ $t('common.back') }}</a-button>
        </div>
      </div>

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
          :pageSize="40"
          :alert="true"
          :rowSelection="rowSelection"
          showPagination="auto"
          class="sort-table"
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
      :title="modelId > 0 ? $t('menu.regex.edit.item') : $t('menu.regex.create.item')"
      :footer="false"
      :centered="true"
      :width="700"
      @cancel="cancelEdit"
    >
      <RegexItemEdit
        :v-if="editVisible"
        :id="modelId"
        :parentId="regexId"
        :afterSave="saveModel"
      />
    </a-modal>
  </div>
</template>

<script>
import moment from 'moment'
import { STable, Ellipsis } from '@/components'
import {
  listRegexItem,
  disableRegexItem,
  removeRegexItem,
  batchRemoveRegexItem,
  resortRegexItem
} from '@/api/manage'

import RegexItemEdit from './Edit'
import Sortable from 'sortablejs'

export default {
  name: 'RegexItems',
  components: {
    STable,
    Ellipsis,
    RegexItemEdit
  },
  props: {
    regexId: {
      type: Number,
      default: function () {
        return parseInt(this.$route.params.regexId)
      }
    }
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
      models: [],
      loadData: parameter => {
        const requestParameters = Object.assign({ regexId: this.regexId }, parameter, this.queryParam)
        return listRegexItem(requestParameters)
          .then(res => {
            this.models = res.data
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
        title: this.$t('form.code'),
        dataIndex: 'code'
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
  mounted () {
    this.initSortable()
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
    initSortable () {
      const that = this
      const el = this.$el.querySelector('.sort-table table tbody')
      Sortable.create(el, {
        handle: '.ant-table-row',
        animation: 150,
        onUpdate: function (evt) {
          console.log(evt)
          that.resort(evt)
        }
      })
    },
    resort (evt) {
      if (evt.oldIndex === evt.newIndex) {
        return
      }

      const src = this.models[evt.oldIndex]
      const target = this.models[evt.newIndex]
      resortRegexItem(src.id, target.id, this.regexId).then(json => {
        this.$refs.table.refresh(false)
      })
    },
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
      disableRegexItem(record).then(json => {
        console.log('disableRegexItem', json)
        this.$refs.table.refresh(false)
      })
    },
    confirmRemove (record) {
      removeRegexItem(record).then(json => {
        console.log('removeRegexItem', json)
        this.$refs.table.refresh(false)
      })
    },
    cancelRemove (e) {
      console.log(e)
    },
    multiOpt ({ key }) {
      console.log(`${key}`)
      batchRemoveRegexItem(this.selectedRowKeys).then(json => {
        console.log('batchRemoveRegexItem', json)
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
    },
    back () {
      this.$router.push('/nlu/regex/list')
    }
  }
}
</script>
