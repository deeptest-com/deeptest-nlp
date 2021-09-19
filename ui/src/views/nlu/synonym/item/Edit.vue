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
        <a-button @click="save()" htmlType="submit" type="primary">{{ $t('form.save') }}</a-button>
        <a-button @click="reset()" style="margin-left: 8px">{{ $t('form.reset') }}</a-button>
      </a-form-item>
    </a-form-model>
  </a-card>
</template>

<script>
import { labelCol, wrapperCol, wrapperFull } from '@/utils/const'
import { requestSuccess, getSynonymItem, saveSynonymItem } from '@/api/manage'

export default {
  name: 'SynonymItemEdit',
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
    return {
      labelCol: labelCol,
      wrapperCol: wrapperCol,
      wrapperFull: wrapperFull,
      model: {},
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
      if (this.id > 0) {
        this.getModel().then(json => {
          this.model = json.data
        })
      } else {
        this.reset()
      }
    },
    getModel () {
      return getSynonymItem(this.id)
    },
    save (e) {
      this.model.synonymId = this.parentId
      console.log(this.model)

      this.$refs.form.validate(valid => {
        if (!valid) {
          console.log('validate fail', valid)
          return false
        }

        saveSynonymItem(this.model).then(json => {
          console.log('saveSynonymItem', json)
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
