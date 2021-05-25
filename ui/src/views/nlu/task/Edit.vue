<template>
  <page-header-wrapper content="">
    <div class="toolbar-edit">
      <div class="left"></div>
      <div class="right">
        <a-button @click="back()" type="primary">{{ $t('common.back') }}</a-button>
      </div>
    </div>

    <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
      <a-form-model ref="form" :model="model" :rules="rules">
        <a-form-model-item
          :label="$t('form.name')"
          prop="name"
          :labelCol="labelCol"
          :wrapperCol="wrapperCol">
          <a-input v-model="model.name" />
        </a-form-model-item>

        <a-form-model-item
          :label="$t('menu.project')"
          prop="projectId"
          :labelCol="labelCol"
          :wrapperCol="wrapperCol">
          <a-select v-model="model.projectId">
            <a-select-option v-for="(item, index) in projects" :value="item.id" :key="index">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-model-item>

        <a-form-model-item
          :label="$t('form.desc')"
          prop="desc"
          :labelCol="labelCol"
          :wrapperCol="wrapperCol">
          <a-input v-model="model.desc" />
        </a-form-model-item>
        <a-form-item
          :wrapperCol="wrapperFull"
          style="text-align: center"
        >
          <a-button @click="save()" htmlType="submit" type="primary">{{ $t('form.save') }}</a-button>
          <a-button @click="reset()" style="margin-left: 8px">{{ $t('form.reset') }}</a-button>
        </a-form-item>
      </a-form-model>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import { labelCol, wrapperCol, wrapperFull } from '@/utils/const'
import { requestSuccess, getTask, saveTask, listProject } from '@/api/manage'

export default {
  name: 'TaskEdit',
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
      labelCol: labelCol,
      wrapperCol: wrapperCol,
      wrapperFull: wrapperFull,
      model: {},
      projects: [],
      rules: {
        name: [{ required: true, message: this.$t('valid.required.name'), trigger: 'blur' }]
      }
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
  methods: {
    loadData () {
      if (this.id) {
        getTask(this.id).then(json => {
          this.model = json.data
        })
      }
      listProject().then(json => {
        this.projects = json.data
      })
    },
    save (e) {
      console.log(this.model)
      this.$refs.form.validate(valid => {
        if (!valid) {
          console.log('validate fail', valid)
          return false
        }

        saveTask(this.model).then(json => {
          console.log('saveTask', json)
          if (requestSuccess(json.code)) {
            this.$router.push('/nlu/task/list')
          }
        })
      })
    },
    reset () {
      this.model = {}
      this.$refs.form.resetFields()
    },
    back () {
      this.$router.push('/nlu/task/list')
    }
  }
}
</script>
