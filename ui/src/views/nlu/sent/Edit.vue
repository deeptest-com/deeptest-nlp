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
          :label="$t('form.content')"
          prop="content"
          :labelCol="labelCol"
          :wrapperCol="wrapperCol">
          <a-input v-model="model.content" />
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
import { requestSuccess, getSent, saveSent } from '@/api/manage'

export default {
  name: 'SentEdit',
  props: {
    intentId: {
      type: Number,
      default: function () {
        return parseInt(this.$route.params.intentId)
      }
    },
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
      rules: {
        content: [{ required: true, message: this.$t('valid.required.content'), trigger: 'blur' }]
      }
    }
  },
  watch: {
    intentId: function () {
      console.log('watch intentId', this.intentId)
    },
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
      return getSent(this.id)
    },
    save (e) {
      console.log(this.model)
      this.$refs.form.validate(valid => {
        if (!valid) {
          console.log('validate fail', valid)
          return false
        }

        saveSent(this.model).then(json => {
          console.log('saveSent', json)
          if (requestSuccess(json.code)) {
            this.$router.push('/nlu/intent/' + this.intentId + '/sent/list')
          }
        })
      })
    },
    reset () {
      this.model = {}
      this.$refs.form.resetFields()
    },
    back () {
      this.$router.push('/nlu/intent/' + this.intentId + '/sent/list')
    }
  }
}
</script>
