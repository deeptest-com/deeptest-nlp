<template>
  <div>
    <page-header-wrapper>
      <a-card :bordered="false">
        <div class="table-page-search-wrapper">
          <a-form layout="inline">
            <a-row :gutter="48">
              <a-col :md="6" :sm="24">
                <a-form-item :label="$t('menu.project')">
                  <a-select v-model="queryParam.projectId">
                    <a-select-option value="0">{{ $t('form.all') }}</a-select-option>
                    <a-select-option v-for="(item, index) in projects" :value="item.id" :key="index">
                      {{item.name}}
                    </a-select-option>
                  </a-select>
                </a-form-item>
              </a-col>
              <a-col :md="6" :sm="24">
                <a-form-item :label="$t('form.name')">
                  <a-input v-model="queryParam.keywords" placeholder=""/>
                </a-form-item>
              </a-col>
              <a-col :md="6" :sm="24">
                <a-form-item :label="$t('form.status')">
                  <a-select v-model="queryParam.status">
                    <a-select-option value="">{{ $t('form.all') }}</a-select-option>
                    <a-select-option value="true">{{ $t('form.enable') }}</a-select-option>
                    <a-select-option value="false">{{ $t('form.disable') }}</a-select-option>
                  </a-select>
                </a-form-item>
              </a-col>
              <a-col :md="!advanced && 6 || 24" :sm="24">
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

          <span slot="project" slot-scope="text">
            <span>{{ text }}</span>
          </span>

          <span slot="action" slot-scope="text, record">
            <template>
              <a @click="edit(record)">{{ $t('form.edit') }}</a>
              <a-divider type="vertical" />
              <a @click="design(record)">{{ $t('form.design') }}</a>

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
                @cancel="cancelRemove">
                <a href="#">{{ $t('form.remove') }}</a>
              </a-popconfirm>
            </template>
          </span>
        </s-table>
      </a-card>
    </page-header-wrapper>

    <div class="full-screen-modal">
      <task-design
        ref="designPage"
        :visible="designVisible"
        :modelProp="designModel"
        :time="time"
        @ok="handleDesignOk"
        @cancel="handleDesignCancel" >
      </task-design>
    </div>
  </div>
</template>

<script>
import moment from 'moment'
import { STable, Ellipsis } from '@/components'
import { listTask, disableTask, removeTask, listForSelect } from '@/api/manage'

import StepByStepModal from '../../list/modules/StepByStepModal'
import CreateForm from '../../list/modules/CreateForm'
import TaskDesign from './Design'

export default {
  name: 'TaskList',
  components: {
    STable,
    Ellipsis,
    CreateForm,
    StepByStepModal,
    TaskDesign
  },
  columns: [],
  statusMap: {},
  data () {
    return {
      designVisible: false,
      designModel: {},
      confirmLoading: false,
      time: 0,
      advanced: false,
      queryParam: { projectId: '0', status: '' },
      projects: [],

      isInit: true, // first time, use projectId in session to query
      loadData: parameter => {
        const requestParameters = Object.assign({ isInit: this.isInit }, parameter, this.queryParam)
        return listTask(requestParameters)
          .then(res => {
            this.isInit = false
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
    listForSelect().then(json => {
      this.projects = json.data.projects
      this.queryParam['projectId'] = json.data.projectId ? json.data.projectId : '0'
    })

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
        title: this.$t('menu.project'),
        dataIndex: 'projectName'
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
      this.$router.push('/nlu/task/0/edit')
    },
    edit (record) {
      this.$router.push('/nlu/task/' + record.id + '/edit')
    },
    design (record) {
      this.time = Date.now() // trigger data refresh
      this.designModel = record
      this.designVisible = true
    },
    handleDesignOk () {
      this.designVisible = false
    },
    handleDesignCancel () {
      this.designVisible = false
      this.designModel = {}
    },
    disable (record) {
      disableTask(record).then(json => {
        console.log('disableTask', json)
        this.$refs.table.refresh(false)
      })
    },
    confirmRemove (record) {
      removeTask(record).then(json => {
        console.log('removeTask', json)
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
