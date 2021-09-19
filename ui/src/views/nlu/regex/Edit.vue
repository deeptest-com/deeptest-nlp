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
          :label="$t('form.code')"
          prop="code"
          :labelCol="labelCol"
          :wrapperCol="wrapperCol">
          <a-input v-model="model.code" />
        </a-form-model-item>
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
import { requestSuccess, getRegex, saveRegex, validDictCode, listProject } from '@/api/manage'

export default {
  name: 'RegexEdit',
  props: {
    id: {
      type: Number,
      default: function () {
        return parseInt(this.$route.params.id)
      }
    }
  },
  data () {
    let checkPending
    const checkCode = (rule, value, callback) => {
      if (this.model.id) {
        callback()
      }

      clearTimeout(checkPending)
      const that = this
      checkPending = setTimeout(() => {
        validDictCode(value, this.model.id, 'regex').then(json => {
          console.log('validDictCode', json)
          if (requestSuccess(json.code) && json.data.pass) {
            callback()
          } else {
            callback(new Error(that.$t('valid.dict.code.unique')))
          }
        })
      }, 500)
    }

    return {
      labelCol: labelCol,
      wrapperCol: wrapperCol,
      wrapperFull: wrapperFull,
      model: {},
      projects: [],
      rules: {
        name: [{ required: true, message: this.$t('valid.required.name'), trigger: 'blur' }],
        code: [{ required: true, message: this.$t('valid.required.code'), trigger: 'blur' },
          { pattern: /^[a-z][_a-z0-9]*$/, message: this.$t('valid.format.code'), trigger: 'blur' },
          { validator: checkCode, trigger: 'change' }]
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
        this.getModel().then(json => {
          this.model = json.data
        })
      } else {
        this.reset()
      }
      listProject().then(json => {
        this.projects = json.data
      })
    },
    getModel () {
      return getRegex(this.id)
    },
    save (e) {
      console.log(this.model)
      this.$refs.form.validate(valid => {
        if (!valid) {
          console.log('validate fail', valid)
          return false
        }

        saveRegex(this.model).then(json => {
          console.log('saveRegex', json)
          if (requestSuccess(json.code)) {
            this.$router.push('/nlu/regex/list')
          }
        })
      })
    },
    reset () {
      this.model = {}
      this.$refs.form.resetFields()
    },
    back () {
      this.$router.push('/nlu/regex/list')
    }
  }
}
</script>
