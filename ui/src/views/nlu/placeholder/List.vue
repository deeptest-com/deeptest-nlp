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
                <a-select v-model="queryParam.status">
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
        :pageSize="40"
        :alert="true"
        showPagination="auto"
        class="sort-table"
      >
        <span slot="serial" slot-scope="text, record, index">
          {{ index + 1 }}
        </span>

        <!--        <span slot="code" slot-scope="text">
                  {{ text }}
                </span>-->
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
            <a @click="detail(record)">{{ $t('form.maintain') }}</a>

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
import { listPlaceholder, disablePlaceholder, removePlaceholder, resortPlaceholder } from '@/api/manage'
import Sortable from 'sortablejs'

export default {
  name: 'PlaceholderList',
  components: {
    STable,
    Ellipsis
  },
  columns: [],
  statusMap: {},
  data () {
    return {
      visible: false,
      confirmLoading: false,
      mdl: null,
      advanced: false,
      queryParam: { status: '' },
      models: [],
      loadData: parameter => {
        const requestParameters = Object.assign({}, parameter, this.queryParam)
        console.log(requestParameters)
        return listPlaceholder(requestParameters)
          .then(res => {
            this.models = res.data
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
  mounted () {
    this.initSortable()
  },
  computed: {
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
      resortPlaceholder(src.id, target.id).then(json => {
        this.$refs.table.refresh(false)
      })
    },
    create () {
      this.mdl = null
      this.visible = true

      this.$router.push('/nlu/placeholder/0/edit')
    },
    edit (record) {
      this.visible = true
      this.mdl = { ...record }

      this.$router.push('/nlu/placeholder/' + record.id + '/edit')
    },
    detail (record) {
      this.visible = true
      this.mdl = { ...record }

      this.$router.push('/nlu/placeholder/' + record.id + '/items')
    },
    disable (record) {
      disablePlaceholder(record).then(json => {
        console.log('disablePlaceholder', json)
        this.$refs.table.refresh(false)
      })
    },
    confirmRemove (record) {
      removePlaceholder(record).then(json => {
        console.log('removePlaceholder', json)
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
