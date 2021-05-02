<template>
  <div>
    {{ model.name }}
  </div>
</template>

<script>

import { getIntent } from '@/api/manage'

export default {
  name: 'IntentEdit',
  props: {
    modelId: {
      type: Number,
      default: () => 0
    }
  },
  data () {
    return {
      model: {}
    }
  },
  watch: {
    modelId: function () {
      console.log('watch modelId', this.modelId)
      this.getModel()
    }
  },
  methods: {
    getModel () {
      console.log('getModel')
      getIntent(this.modelId).then(json => {
        console.log(json)
        if (json.code === 200) {
          this.model = json.data
        }
      })
    },
    save (e) {
      console.log('save')
    },
    reset () {
      this.model = {}
      // this.$refs.form.resetFields()
    },
    back () {
      this.$router.push('/nlu/task/list')
    }
  }
}
</script>
