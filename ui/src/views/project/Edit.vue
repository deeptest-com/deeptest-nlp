<template>
  <div>
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
            :label="$t('form.path')"
            prop="path"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol">
            <a-input v-model="model.path" />
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
  </div>
</template>

<script>
import { labelCol, wrapperCol, wrapperFull } from '@/utils/const'
import { requestSuccess, getProject, saveProject, validProjectPath } from '@/api/manage'

export default {
  name: 'ProjectEdit',
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
    const checkPath = (rule, value, callback) => {
      clearTimeout(checkPending)
      checkPending = setTimeout(() => {
        validProjectPath(value).then(json => {
          console.log('validProjectPath', json)
          if (requestSuccess(json.code) && json.data.pass) {
            callback()
          } else {
            callback(new Error(this.$t('valid.project.path')))
          }
        })
      }, 500)
    }

    return {
      labelCol: labelCol,
      wrapperCol: wrapperCol,
      wrapperFull: wrapperFull,
      model: {},
      rules: {
        name: [{ required: true, message: this.$t('valid.required.name'), trigger: 'blur' }],
        path: [{ required: true, message: this.$t('valid.required.path'), trigger: 'blur' },
               { validator: checkPath, trigger: 'change' }]
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
      if (!this.id) {
        return
      }
      if (this.id) {
        this.getModel().then(json => {
          this.model = json.data.model
        })
      } else {
        this.reset()
      }
    },
    getModel () {
      return getProject(this.id)
    },
    save (e) {
      console.log(this.model)
      this.$refs.form.validate(valid => {
        if (!valid) {
          console.log('validate fail', valid)
          return false
        }

        saveProject(this.model).then(json => {
          console.log('saveProject', json)
          if (requestSuccess(json.code)) {
            this.$router.push('/project/list')
          }
        })
      })
    },
    reset () {
      this.model = {}
      this.$refs.form.resetFields()
    },
    back () {
      this.$router.push('/project/list')
    }
  }
}
</script>
