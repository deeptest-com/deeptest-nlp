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
            <template v-if="advanced">
            </template>
            <a-col :md="!advanced && 8 || 24" :sm="24">
              <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
                <a-button type="primary" @click="$refs.table.refresh(true)">{{ $t('form.search') }}</a-button>
                <a-button style="margin-left: 8px" @click="() => this.queryParam = {}">{{ $t('form.reset') }}</a-button>
                <a @click="toggleAdvanced" style="margin-left: 8px">
                  {{ advanced ? $t('form.collapse') : $t('form.expand') }}
                  <a-icon :type="advanced ? 'up' : 'down'"/>
                </a>
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

        <span slot="status" slot-scope="text">
          <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
        </span>

        <span slot="default" slot-scope="text, record">
          <template>={{record.isDefault}}=
            <a-checkbox :checked="record.isDefault==true" @click="setDefault(record)">
            </a-checkbox>
          </template>
        </span>

        <span slot="action" slot-scope="text, record">
          <template>
            <a @click="edit(record)">{{ $t('form.edit') }}</a>
            <a-divider type="vertical" />
            <a v-if="!record.disabledAt" @click="disable(record)">{{ $t('form.disable') }}</a>
            <a v-if="record.disabledAt" @click="disable(record)">{{ $t('form.enable') }}</a>
            <a-divider type="vertical" />
            <a @click="remove(record)">{{ $t('form.remove') }}</a>
          </template>
        </span>
      </s-table>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { STable, Ellipsis } from '@/components'
import { listProject, setDefaultProject, disableProject, removeProject } from '@/api/manage'

import StepByStepModal from '../../list/modules/StepByStepModal'
import CreateForm from '../../list/modules/CreateForm'

const columns = [
  {
    title: '#',
    scopedSlots: { customRender: 'serial' }
  },
  {
    title: '名称',
    dataIndex: 'name'
  },
  {
    title: '状态',
    dataIndex: 'status',
    scopedSlots: { customRender: 'status' }
  },
  {
    title: '是否默认',
    dataIndex: 'default',
    scopedSlots: { customRender: 'default' }
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '150px',
    scopedSlots: { customRender: 'action' }
  }
]

const statusMap = {
  0: {
    status: 'default',
    text: '关闭'
  },
  1: {
    status: 'processing',
    text: '运行中'
  },
  2: {
    status: 'success',
    text: '已上线'
  },
  3: {
    status: 'error',
    text: '异常'
  }
}

export default {
  name: 'ProjectList',
  components: {
    STable,
    Ellipsis,
    CreateForm,
    StepByStepModal
  },
  data () {
    this.columns = columns
    return {
      // create model
      visible: false,
      confirmLoading: false,
      mdl: null,
      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: {},
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        const requestParameters = Object.assign({}, parameter, this.queryParam)
        // console.log('loadData request parameters:', requestParameters)
        return listProject(requestParameters)
          .then(res => {
            return res
          })
      },
      selectedRowKeys: [],
      selectedRows: []
    }
  },
  filters: {
    statusFilter (type) {
      return statusMap[1].text
    },
    statusTypeFilter (type) {
      return statusMap[1].status
    }
  },
  created () {
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

      this.$router.push('/platform/project/0/edit')
    },
    edit (record) {
      this.visible = true
      this.mdl = { ...record }

      this.$router.push('/platform/project/' + record.id + '/edit')
    },
    setDefault (record) {
      console.log('===', record)
      this.visible = true
      setDefaultProject(record).then(json => {
        console.log('setDefaultProject', json)
        this.loadData()
      })
    },
    disable (record) {
      this.visible = true
      disableProject(record).then(json => {
        console.log('disableProject', json)
        this.loadData()
      })
    },
    remove (record) {
      this.visible = true
      removeProject(record).then(json => {
        console.log('removeProject', json)
        this.loadData()
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
