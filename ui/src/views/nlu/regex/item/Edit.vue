<template>
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
</template>

<script>
import { labelCol, wrapperCol, wrapperFull } from '@/utils/const'
import { requestSuccess, getRegexItem, saveRegexItem, validDictCode } from '@/api/manage'

export default {
  name: 'RegexItemEdit',
  props: {
    id: {
      type: Number,
      default: 0
    },
    parentId: {
      type: Number,
      default: 0
    },
    afterSave: {
      type: Function,
      default: null
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
        validDictCode(value, this.model.id, 'regexItem').then(json => {
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
      if (this.id > 0) {
        this.getModel().then(json => {
          this.model = json.data
        })
      } else {
        this.reset()
      }
    },
    getModel () {
      return getRegexItem(this.id)
    },
    save (e) {
      this.model.regexId = this.parentId
      console.log(this.model)

      this.$refs.form.validate(valid => {
        if (!valid) {
          console.log('validate fail', valid)
          return false
        }

        saveRegexItem(this.model).then(json => {
          console.log('saveRegexItem', json)
          if (requestSuccess(json.code)) {
            if (this.afterSave) {
              this.afterSave(json)
            }
          }
        })
      })
    },
    reset () {
      this.model = {}
      this.$refs.form.resetFields()
    }
  }
}
</script>
