<template>
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
        <a-button @click="save()" htmlType="submit" type="primary">{{$t('form.submit')}}</a-button>
        <a-button @click="reset()" style="margin-left: 8px">{{$t('form.reset')}}</a-button>
      </a-form-item>
    </a-form-model>
  </a-card>
</template>

<script>
import { labelCol, wrapperCol, wrapperFull } from '@/utils/const'
import { requestSuccess, getLookupItem, saveLookupItem } from '@/api/manage'

export default {
  name: 'LookupItemEdit',
  props: {
    id: {
      type: Number,
      default: 0
    },
    afterSave: Function
  },
  data () {
    return {
      labelCol: labelCol,
      wrapperCol: wrapperCol,
      wrapperFull: wrapperFull,
      model: {},
      rules: {
        name: [{ required: true, message: this.$t('valid.input.name'), trigger: 'blur' }]
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
          this.model = json.data
        })
      } else {
        this.reset()
      }
    },
    getModel () {
      return getLookupItem(this.id)
    },
    save (e) {
      console.log(this.model)
      this.$refs.form.validate(valid => {
        if (!valid) {
          console.log('validate fail', valid)
          return false
        }

        saveLookupItem(this.model).then(json => {
          console.log('saveLookupItem', json)
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
